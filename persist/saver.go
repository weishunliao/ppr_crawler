package persist

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
)

func PropertySaver() (chan interface{}, error) {
	out := make(chan interface{})
	esClient, _ := elasticsearch.NewDefaultClient()
	res, err := esClient.Info()
	defer res.Body.Close()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, err
	}
	log.Println(res)
	go func() {
		for {
			property := <- out
			save(property, esClient)
		}
	}()
	return out, nil
}

func save(property interface{}, esClient *elasticsearch.Client) {
	data := property.(map[string]string)
	jsonString, _ := json.Marshal(property)
	request := esapi.IndexRequest{
		Index:      "ppr",
		DocumentID: data["id"],
		Body:       strings.NewReader(string(jsonString)),
		Refresh:    "true",
	}
	res, err := request.Do(context.Background(), esClient)
	if err != nil {
		return 
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%s", res.Status(), data["id"])
	} else {
		log.Printf("Indexing document ID=%s done", data["id"])
	}
}
