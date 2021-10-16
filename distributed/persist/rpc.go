package persist

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/weishunliao/crawler/engine"
	"github.com/weishunliao/crawler/persist"
)

type PropertySaverService struct {
	Client *elasticsearch.Client
}

func (s PropertySaverService) Save(property engine.Property, result *string) error {
	err := persist.Save(property, s.Client)
	if err == nil {
		*result = "Ok"
	}
	return err
}