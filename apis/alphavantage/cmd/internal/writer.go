package internal

import (
	"context"
	"fmt"
	"github.com/triskaidekaphile/go-marketdata/apis/alphavantage"
	"log"
	"os"
	"path/filepath"
)

var Stocks = [10]string{
	"TSLA",
	"AMZN",
	"AAPL",
	"KO",
	"NU",
	"GOLD",
	"AMD",
	"NKE",
	"NVDA",
	"BRK-B",
}

var Cryptos = [10]string{
	"BTC",
	"ETH",
	"BNB",
	"USDT",
	"XRP",
	"SOL",
	"USDC",
	"ADA",
	"DOGE",
	"SHIB",
}

func WriteStocks(ctx context.Context, api alphavantage.API) {
	var name string
	for _, stock := range Stocks {
		ts, err := api.GetStockTimeSeriesDaily(ctx, stock)
		if err != nil {
			log.Printf("error on get symbol [%s]: %s\n", stock, err.Error())
			continue
		}

		name = fmt.Sprintf("stocks/%s", stock)

		err = Write(name, []byte(ts.JSON(true)))
		if err != nil {
			log.Printf("error on write time series %s\n", err.Error())
		}
	}
}

func WriteCryptos(ctx context.Context, api alphavantage.API) {
	var name string
	for _, crypto := range Cryptos {
		ts, err := api.GetCryptoTimeSeriesDaily(ctx, crypto)
		if err != nil {
			log.Printf("error on get crypto [%s]: %s\n", crypto, err.Error())
			continue
		}

		name = fmt.Sprintf("cryptos/%s", crypto)

		err = Write(name, []byte(ts.JSON(true)))
		if err != nil {
			log.Printf("error on write crypto time series %s\n", err.Error())
		}
	}
}

func Write(name string, b []byte) error {
	var (
		filename = fmt.Sprintf("%s.json", name)
		fullPath = filepath.Join("./assets", filename)
	)

	writer, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error on create file %s: %w", fullPath, err)
	}

	defer func() { _ = writer.Close() }()

	_, err = writer.Write(b)
	if err != nil {
		return fmt.Errorf("error on write file %s: %w", fullPath, err)
	}

	err = writer.Sync()
	if err != nil {
		return fmt.Errorf("error on sync file %s: %w", fullPath, err)
	}

	return nil
}
