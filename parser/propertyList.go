package parser

import (
	"fmt"
	"github.com/weishunliao/crawler/engine"
	"regexp"
)

const pprIdRe = `UNID-([0-9a-zA-Z]+)?`
const baseUrl = "https://www.propertypriceregister.ie/Website/npsra/PPR/npsra-ppr.nsf/eStampUNID/UNID-"
func ParsePropertyList(data []byte) engine.ParseResult {
	result := engine.ParseResult{}
	re1 := regexp.MustCompile(pprIdRe)
	pprIds := re1.FindAllSubmatch(data, -1)
	fmt.Printf("Total: %d\n", len(pprIds))
	for _, match := range pprIds {
		result.Properties = append(result.Properties, "id-" + string(match[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url: baseUrl + string(match[1]) + "?OpenDocument",
			ParseFunc: ParsePropertyDetail,
		})
	}
	return result
}
