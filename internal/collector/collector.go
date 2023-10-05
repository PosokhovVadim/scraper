package collector

import (
	colly "github.com/gocolly/colly/v2"
)

type Collector struct {
	C *colly.Collector
}

func CollectorInit(allowedDomains []string, cacheDir string) *Collector {
	return &Collector{
		C: colly.NewCollector(
			colly.AllowedDomains(allowedDomains...),
			colly.CacheDir(cacheDir),
		),
	}
}
