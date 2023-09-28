package scraper

// callback
func SaveData(selector string, data string) {

}

func extractData(selector string, data string) error {
	return nil
}

// func Scraper(c *collector.Collector, startedURL string) {
// 	fmt.Println("I am here")
// 	for _, selector := range model.InfoSelectros {
// 		c.C.OnHTML(selector, func(e *colly.HTMLElement) {
// 			//TODO: extract data and save in to Storage (now it is file) зачем мне дата ???
// 			// extractData(selector, e.Text) // ???
// 			fmt.Println("write data to storage")
// 		})
// 	}

// 	c.C.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		//c.C.Visit(e.Attr("href"))
// 		fmt.Println("I am on the page: " + e.Attr("href"))
// 	})

// }
