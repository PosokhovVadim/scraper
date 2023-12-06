package main

import (
	"fmt"
	"scraper/internal/model"
	"scraper/internal/storage"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var allowedDomains = "www.noone.ru"

const startedURL = "https://www.noone.ru/"
const storagePath = "mongodb://localhost:27017/"
const cacheDir = "cacheUrls"

func main() {
	storage, err := storage.ConnectStorage(storagePath)
	if err != nil {
		fmt.Printf("Error initial loading storage: %e\n", err)
	}
	defer storage.CloseStorage()
	if err != nil {
		fmt.Printf("Error closing storage: %e", err)
	}

	start := time.Now()
	scraper := model.ScraperInit(allowedDomains, cacheDir)

	scraper.SetupCallback()

	scraper.C.Visit(startedURL)
	
	end := time.Since(start)
	fmt.Println(len(scraper.ProductData))
	for _, product := range scraper.ProductData {
		data, _ := bson.Marshal(product)
		storage.InsertData(data)
		if err != nil {
			fmt.Printf("Error inserting: %e", err)
		}
	}


	fmt.Println("Time of execution:", end)
	fmt.Println("Count of page:", model.Count)

}
