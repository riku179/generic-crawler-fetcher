package fetcher

import (
	"github.com/gocolly/colly"
)

type Crawler struct {
	id string
	entryUrl string
	operations   []Operation
	current	     *Operation
	corrector colly.Collector
}

func NewCrawler(siteMap SiteMap) Crawler {
	if len(siteMap.StartUrl) != 1 {
		panic("siteMap.StartUrl needs just one")
	}
	return Crawler{
		id: siteMap.Id,
		entryUrl: siteMap.StartUrl[0],
		operations: NewOperations(siteMap.Selectors),
		current: nil,
	}
}

func (c *Crawler) Exec() {
	for _, rootOp := range c.operations {
		switch rootOp.Type() {
		case Element:
		}
	}
}
