package model

type PageData struct {
	Url         string   `json:"url"`
	RefUrls     []string `json:"ref-urls`
	ProductData ProductData
}

// Пишем вообще все что захотим в данном месте, просто в другом реализовываем нужные нам селекторы
type ProductData struct {
	Titile     string   `json:"titile"`
	Brand      string   `json:"brand"`
	Category   string   `json:"category"`
	FullName   string   `json:"full-name"`
	Price      string   `json:"price"`
	SizeList   []string `json:"size-list"`
	Properties []string `json:"properties"`
}

//SECELTORS: class names of blocks with needed information

var InfoSelectros = []string{"div.item-brand", "div.item-category", "div.item-fullname", "div.item-price"}
