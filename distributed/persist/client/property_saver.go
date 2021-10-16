package client

import (
	"github.com/weishunliao/crawler/config"
	"github.com/weishunliao/crawler/distributed/rpc_support"
	"github.com/weishunliao/crawler/engine"
	"log"
)

func PropertySaver(host string) (chan engine.Property, error) {
	client, err := rpc_support.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Property)
	go func() {
		for {
			property := <- out
			result := ""
			err = client.Call(config.PropertySaverRPC, property, &result)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	return out, nil
}
