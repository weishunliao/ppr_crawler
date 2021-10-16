package parser

import (
	"github.com/anaskhan96/soup"
	"github.com/weishunliao/crawler/engine"
	"regexp"
	"strings"
)


func ParsePropertyDetail(data []byte) engine.ParseResult {
	re := regexp.MustCompile(pprIdRe)
	match := re.FindSubmatch(data)
	doc := soup.HTMLParse(string(data))
	pprInfo := make([]string, 13)
	pprInfo[0] = string(match[1])
	for i, info := range doc.Find("table", "id", "AddInfo").FindAll("td") {
		pprInfo[i + 1] = strings.TrimSpace(info.Text())
	}
	for i, info := range doc.Find("table", "id", "SaleInfo").FindAll("td") {
		if i % 2 != 0 {
			pprInfo[(i / 2) + 7] = strings.TrimSpace(info.Text())
		}
	}

	m := make(map[string]interface{})

	m["id"] = pprInfo[0]
	m["address1"] = pprInfo[1]
	m["address2"] = pprInfo[2]
	m["address3"] = pprInfo[3]
	m["address4"] = pprInfo[4]
	m["address5"] = pprInfo[5]
	m["eircode"] = pprInfo[6]
	m["dateOfSale"] = pprInfo[7]
	m["price"] = pprInfo[8]
	m["notFullMarketPrice"] = pprInfo[9]
	m["VATExclusive"] = pprInfo[10]
	m["description"] = pprInfo[11]

	result := engine.ParseResult{}
	result.Properties = append(result.Properties, m)
	result.Store = true
	return result
}
