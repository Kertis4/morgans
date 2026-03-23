package main

import (
	"context"
	"fmt"
	"log"
	"os"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	finnhubKey := os.Getenv("FINNHUB_KEY")

	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", finnhubKey)
	finhubbClient := finnhub.NewAPIClient(cfg).DefaultApi

	generalNews, _, err := finhubbClient.MarketNews(context.Background()).Category("general").Execute()
	if err != nil {
		log.Printf("Error Fetching MarketNews %v", err)
	}
	for _, article := range generalNews {
		fmt.Println(*article.Headline)
		fmt.Println(*article.Summary)
		fmt.Println(*article.Url)
		if *article.Related != "" {
			fmt.Println(*article.Related)
		} else {
			fmt.Println("No Related Companies")
		}
		fmt.Println("----------------------------------")
	}
}
