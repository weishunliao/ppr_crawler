package parser

import (
	"github.com/anaskhan96/soup"
	"github.com/weishunliao/crawler/engine"
	"regexp"
)


func ParsePropertyDetail(data []byte) engine.ParseResult {
	re := regexp.MustCompile(pprIdRe)
	match := re.FindSubmatch(data)
	doc := soup.HTMLParse(string(data))
	pprInfo := make([]string, 13)
	pprInfo[0] = string(match[1])
	for i, info := range doc.Find("table", "id", "AddInfo").FindAll("td") {
		pprInfo[i + 1] = info.Text()
	}
	for i, info := range doc.Find("table", "id", "SaleInfo").FindAll("td") {
		if i % 2 != 0 {
			pprInfo[(i / 2) + 7] = info.Text()
		}
	}
	//fmt.Println("pprInfo: ",pprInfo)
	//time.Sleep(2 * time.Second)
	result := engine.ParseResult{}
	result.Properties = append(result.Properties, pprInfo)
	result.Store = true
	return result
}
