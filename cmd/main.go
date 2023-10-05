package main

import (
	"fmt"
	"scraper/internal/model"
	"time"
)

var allowedDomains = "www.noone.ru"

const startedURL = "https://www.noone.ru/product/botinki-1772996/"

const cacheDir = "cacheUrls"

// TODO: REFACTOR ALL
func main() {

	//c := collector.CollectorInit(allowedDomains, cacheDir)
	scraper := model.ScraperInit(allowedDomains, cacheDir)
	scraper.SetupCallback()

	scraper.C.Visit(startedURL)
	start := time.Now()

	end := time.Since(start)

	fmt.Println("Time of execution:", end)
	fmt.Println("Count of page:", model.Count)

}
