package parser

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/weishunliao/crawler/engine"
	"time"
)


func ParsePropertyDetail(data []byte) engine.ParseResult {
	doc := soup.HTMLParse(string(data))
	addInfo := make([]string, 6)
	saleInfo := make([]string, 6)
	for i, info := range doc.Find("table", "id", "AddInfo").FindAll("td") {
		addInfo[i] = info.Text()
	}
	for i, info := range doc.Find("table", "id", "SaleInfo").FindAll("td") {
		if i % 2 != 0 {
			saleInfo[i / 2] = info.Text()
		}
	}
	fmt.Println(addInfo)
	fmt.Println(saleInfo)
	time.Sleep(2 * time.Second)
	return engine.ParseResult{}
}
