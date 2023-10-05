package model

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

var Count = 0

type Scraper struct {
	C           *colly.Collector
	ProductData []ProductData
}

func ScraperInit(allowedDomains, cacheDir string) *Scraper {
	return &Scraper{C: colly.NewCollector(
		colly.AllowedDomains(allowedDomains),
		colly.CacheDir(cacheDir),
	),
		ProductData: nil,
	}
}

func (s *Scraper) VisitedURLS() {
	s.C.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		visited, err := s.C.HasVisited(link)
		if err != nil {
			fmt.Errorf("Error in function HAsVisited: %v", err)
			e.Request.Abort()
		}
		if !visited {
			Count++
			e.Request.Visit(link)
		}
	})
}

func (s *Scraper) ExtractElementData(selector string) string {
	var data []string
	fmt.Println(selector)
	s.C.OnHTML(selector, func(e *colly.HTMLElement) {
		fmt.Println(e.Attr(selector))
		fmt.Printf("Selector: %v, Value:%s \n", selector, e.Text)
		data = append(data, e.Text)
	})
	fmt.Println(data)
	return strings.Join(data, ",")

}

func (s *Scraper) ExtractPageData(e *colly.HTMLElement) {
	// s.C.OnRequest(func(r *colly.Request) {
	// 	link := r.URL.String()
	// 	fmt.Printf("Link from ExtractPage: %s\n", link)
	// 	if slices.Contains(strings.Split(link, "/"), "product") {
	// fmt.Printf("HTML tag: %s,  ", e.Name)
	re := regexp.MustCompile(`\s+`)

	PageData := ProductData{
		// Title: e.ChildText(selectors["Title"]), - не работает
		Brand:    e.ChildText(selectors["Brand"]),
		Category: e.ChildText(selectors["Category"]),
		FullName: e.ChildText(selectors["FullName"]),
		Price:    re.ReplaceAllString(strings.TrimSpace(e.ChildText(selectors["Price"])), " "),
		// SizeList: e.ChildText(selectors["SizeList"]),
		Properties: e.ChildTexts(selectors["Properties"]),
		//Properties:  strings.TrimSpace(e.ChildTexts(selectors["Properties"])),
	}
	fmt.Println(PageData.Brand, PageData.Category, PageData.FullName, PageData.Price, PageData.Properties)
	// 	}
	// })
}

func (s *Scraper) SetupCallback() {
	s.C.OnHTML(".item-details", s.ExtractPageData)
	//s.VisitedURLS()
	// s.ExtractPageData()
}

type ProductData struct {
	Title      string   `json:"title"`
	Brand      string   `json:"brand"`
	Category   string   `json:"category"`
	FullName   string   `json:"full-name"`
	Price      string   `json:"price"`
	SizeList   string   `json:"size-list"`
	Properties []string `json:"properties"`
}

var selectors = map[string]string{
	"Title":      ".title",
	"Brand":      "div.item-brand",
	"Category":   "div.item-category",
	"FullName":   "div.item-fullname",
	"Price":      "span[itemprop=price]",
	"SizeList":   "li.item-size",
	"Properties": "div.item-prop",
}
