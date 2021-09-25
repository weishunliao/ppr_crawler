package main

import (
	"fmt"
	"github.com/weishunliao/crawler/engine"
	"github.com/weishunliao/crawler/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://www.propertypriceregister.ie/Website/npsra/PPR/npsra-ppr.nsf/PPR-By-Date&Start=1&Query=%5Bdt_execution_date%5D%3E=01/01/2021%20AND%20%5Bdt_execution_date%5D%3C01/2/2021%20AND%20%5Bdc_county%5D=Carlow&County=Carlow&Year=2021&StartMonth=01&EndMonth=01&Address=",
		ParseFunc: parser.ParsePropertyList,
	})
	fmt.Println("Done")
}
