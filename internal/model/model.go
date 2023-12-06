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
	}
}

func (s *Scraper) VisitedURLS(e *colly.HTMLElement) {
	link := e.Attr("href")
	visited, err := s.C.HasVisited(link)
	if err != nil {
		fmt.Errorf("Error in function HAsVisited: %v", err)
		e.Request.Abort()
	}
	if !visited {
		Count++
		if Count%100000 == 0 {
			fmt.Println(Count)
		}
		// fmt.Println(link)
		e.Request.Visit(link)
	}
}

func (s *Scraper) PageDataIsValid(p *ProductData) bool {
	return p.Brand != "" && p.Category != "" && p.FullName != "" && p.Price != "" &&
		p.Color != "" && len(p.Properties) > 0
}
func (s *Scraper) ExtractPageData(e *colly.HTMLElement) {
	re := regexp.MustCompile(`[\s\n]+`)
	PageData := ProductData{
		Brand:      e.ChildText(selectors["Brand"]),
		Category:   e.ChildText(selectors["Category"]),
		FullName:   e.ChildText(selectors["FullName"]),
		Price:      re.ReplaceAllString(strings.TrimSpace(e.ChildText(selectors["Price"])), " "),
		Color:      e.ChildText(selectors["Color"]),
		Properties: e.ChildTexts(selectors["Properties"]),
	}
	for i := range PageData.Properties {
		PageData.Properties[i] = re.ReplaceAllString(strings.TrimSpace(PageData.Properties[i]), " ")
	}
	if s.PageDataIsValid(&PageData) {
		s.ProductData = append(s.ProductData, PageData)
	}
}

func (s *Scraper) SetupCallback() {
	s.C.OnHTML("a[href]", s.VisitedURLS)
	s.C.OnHTML(".item-details", s.ExtractPageData)
	// s.VisitedURLS()
	// s.ExtractPageData() if slices.Contains(strings.Split(link, "/"), "product") {
}

type ProductData struct {
	Brand      string   `json:"brand"`
	Category   string   `json:"category"`
	FullName   string   `json:"full-name"`
	Price      string   `json:"price"`
	Color      string   `json:"color"`
	Properties []string `json:"properties"`
}

var selectors = map[string]string{
	"Brand":      "div.item-brand",
	"Category":   "div.item-category",
	"FullName":   "div.item-fullname",
	"Price":      "span[itemprop=price]",
	"Color":      "div.item-color-name",
	"Properties": "div.item-prop",
}
