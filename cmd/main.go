package main

import (
	"fmt"
	c "scraper/internal/collector"
	"strings"
	"time"

	"slices"

	"github.com/gocolly/colly/v2"
)

const startedURL = "https://www.noone.ru/"

// TODO: REFACTOR ALL
func main() {
	// TODO: create a config
	c := c.CollectorInit()
	//Implement Logger

	start := time.Now()
	c.C.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href") // ABSOLUTE PATH OR NOT?
		visited, err := c.C.HasVisited(link)
		if err != nil {
			fmt.Errorf("Error in function HAsVisited: %v", err) // TODO: CHANGE to logger
			e.Request.Abort()
		}
		absUrl := e.Request.AbsoluteURL(link)

		if !visited {
			e.Request.Visit(absUrl)
			c.C.OnRequest(func(r *colly.Request) {
				if slices.Contains(strings.Split(absUrl, "/"), "product") {
					fmt.Println("We on page with product url:", absUrl)
					//TODO: EXTRACT DATA, SAVE IT IN JSON

				}

				fmt.Println("we on page:", r.URL)
			})

		}

	})

	c.C.Visit("https://www.noone.ru/")
	//TODO: CHECK SPEED
	end := time.Since(start)
	fmt.Println("Time of execution:", end)
}
