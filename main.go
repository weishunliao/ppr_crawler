package main

import (
	"fmt"
	"github.com/weishunliao/crawler/engine"
	"github.com/weishunliao/crawler/parser"
	"github.com/weishunliao/crawler/persist"
	"github.com/weishunliao/crawler/scheduler"
	"strconv"
	"time"
)

const BaseUrl = "https://www.propertypriceregister.ie/Website/npsra/PPR/npsra-ppr.nsf/"

func main() {
	start := time.Now().UTC()
	fmt.Println("Start at: ", start)
	propertyChan, err := persist.PropertySaver()
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler: &scheduler.QueueScheduler{},
		WorkerCount: 10,
		PropertyChan: propertyChan,
	}

	concurrentEngine.Run(engine.Request{
		Url:       getUrl("Dublin", 2020, 1, 1),
		ParseFunc: parser.ParsePropertyList,
	})
	//engine.SingleEngine{}.Run(engine.Request{
	//	Url:       getUrl("Tipperary", 2021, 1, 1),
	//	ParseFunc: parser.ParsePropertyList,
	//})
	fmt.Println("Done, duration:", time.Now().Sub(start).Seconds())
}

func getUrl(county string, year int, from int, to int) string {
	y := strconv.Itoa(year)
	f := strconv.Itoa(from)
	t := strconv.Itoa(to + 1)
	return BaseUrl + "/PPR-By-Date&Start=1&Query=%5Bdt_execution_date%5D%3E=01/" + f + "/" + y + "%20AND%20%5Bdt_execution_date%5D%3C01/" + t + "/" + y + "%20AND%20%5Bdc_county%5D=" + county
}
