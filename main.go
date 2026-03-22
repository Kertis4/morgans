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

	quote, _, err := finhubbClient.Quote(context.Background()).Symbol("AAPL").Execute()
	if err != nil {
		log.Printf("error fetching quote %v", err)
	}
	fmt.Printf(" Current: %+v\n", *quote.C)
	fmt.Printf(" High: %+v\n", *quote.H)
	fmt.Printf(" Low: %+v\n", *quote.L)
	fmt.Printf(" Open: %+v\n", *quote.O)
}
