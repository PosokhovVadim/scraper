package collector

import (
	colly "github.com/gocolly/colly/v2"
)

type Collector struct {
	C    *colly.Collector
	Data []byte
}

func CollectorInit() *Collector {
	return &Collector{
		C: colly.NewCollector(),
	}

}
