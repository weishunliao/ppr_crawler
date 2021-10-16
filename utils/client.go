package utils

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func CreateESClient() (*elasticsearch.Client, error) {
	esClient, _ := elasticsearch.NewDefaultClient()
	res, err := esClient.Info()
	defer res.Body.Close()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, err
	}
	log.Println(res)
	return esClient, nil
}
