package main

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
)
import "../../rpcsupport"
import "../../persist"

func main() {
	log.Fatalln(serverRpc(":1234", "dating_profile"))
}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = rpcsupport.ServerRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
	return err
}
