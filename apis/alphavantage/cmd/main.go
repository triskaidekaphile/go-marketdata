package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/triskaidekaphile/go-marketdata/apis/alphavantage"
	"github.com/triskaidekaphile/go-marketdata/apis/alphavantage/cmd/internal"
	"log"
	"os"
)

var (
	baseURL string
	apiKey  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}

	baseURL = os.Getenv("ALPHA_VANTAGE_BASE_URL")
	apiKey = os.Getenv("ALPHA_VANTAGE_API_KEY")
}

func main() {
	var (
		ctx = context.Background()
		api = alphavantage.New(baseURL, apiKey, false)
	)

	internal.WriteStocks(ctx, api)
	internal.WriteCryptos(ctx, api)
}
