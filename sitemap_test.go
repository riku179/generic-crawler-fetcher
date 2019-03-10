package fetcher_test

import (
	"encoding/json"
	"fetcher"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestSiteMap_UnmarshalJSON(t *testing.T) {
	input, err := ioutil.ReadFile("./sitemap.json")
	check(err)

	var sitemap fetcher.SiteMap
	err = json.Unmarshal(input, &sitemap)
	check(err)
	assert.Equal(t, "f1-data", sitemap.Id)
	assert.True(t, sitemap.Selectors[0].IsMultiple())
	assert.Equal(t, fetcher.SelectorId("link"), sitemap.Selectors[1].GetId())
	assert.Equal(t, fetcher.Image, sitemap.Selectors[2].GetType())
	assert.Equal(t, "datetime", sitemap.Selectors[3].(*fetcher.SelectorElementAttr).ExtractAttribute)
}
