package fetcher_test

import (
	"encoding/json"
	"fetcher"
	"io/ioutil"
	"testing"

	"github.com/k0kubun/pp"
)

func TestNewOperations(t *testing.T) {
	input, err := ioutil.ReadFile("./sitemaps/simple.json")
	check(err)

	var sitemap fetcher.SiteMap
	err = json.Unmarshal(input, &sitemap)
	check(err)

	operations := fetcher.NewOperations(sitemap.Selectors)
	_, _ = pp.Print(operations)
}
