package persist

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/weishunliao/crawler/config"
	"github.com/weishunliao/crawler/engine"
	"github.com/weishunliao/crawler/utils"
	"log"
	"strings"
)

func PropertySaver() (chan engine.Property, error) {
	out := make(chan engine.Property)
	esClient, err := utils.CreateESClient()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			property := <- out
			Save(property, esClient)
		}
	}()
	return out, nil
}

func Save(property engine.Property, esClient *elasticsearch.Client) error {
	data := property.(map[string]interface{})
	jsonString, _ := json.Marshal(property)
	request := esapi.IndexRequest{
		Index:      config.EsIndex,
		DocumentID: data["id"].(string),
		Body:       strings.NewReader(string(jsonString)),
		Refresh:    "true",
	}
	res, err := request.Do(context.Background(), esClient)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%s", res.Status(), data["id"])
	} else {
		log.Printf("Indexing document ID=%s on data %s done.", data["id"], data["dateOfSale"])
	}
	return nil
}
