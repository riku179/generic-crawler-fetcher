package fetcher_test

import (
	"encoding/json"
	"fetcher"
	"github.com/k0kubun/pp"
	"io/ioutil"
	"testing"
)

func TestNewOperations(t *testing.T) {
	input, err := ioutil.ReadFile("./sitemap.json")
	check(err)

	var sitemap fetcher.SiteMap
	err = json.Unmarshal(input, &sitemap)
	check(err)

	operations := fetcher.NewOperations(sitemap.Selectors)
	_, _ = pp.Print(operations)
}
