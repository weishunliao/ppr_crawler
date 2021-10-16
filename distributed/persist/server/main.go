package main

import (
	"fmt"
	"github.com/weishunliao/crawler/config"
	"github.com/weishunliao/crawler/distributed/persist"
	"github.com/weishunliao/crawler/distributed/rpc_support"
	"github.com/weishunliao/crawler/utils"
)

func main() {
	client, err := utils.CreateESClient()
	if err != nil {
		panic(err)
	}
	rpc_support.ServeRPC(fmt.Sprintf(":%d", config.PropertySaverPort), persist.PropertySaverService{
		Client: client,
	})
}
