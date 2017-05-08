package main

import (
	"log"
	"os"

	"github.com/hashicorp/vault-plugins/database/mongodb"
	"github.com/hashicorp/vault/helper/pluginutil"
)

func main() {
	apiClientMeta := &pluginutil.APIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args)

	err := mongodb.Run(apiClientMeta.GetTLSConfig())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
