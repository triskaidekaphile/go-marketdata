package types

import "encoding/json"

type (
	CryptoTimeSeriesDaily struct {
		Metadata   *CryptoMetadata   `json:"Meta Data"`
		TimeSeries *CryptoTimeSeries `json:"Time Series (Digital Currency Daily)"`
	}

	CryptoMetadata struct {
		Information         string `json:"1. Information"`
		DigitalCurrencyCode string `json:"2. Digital Currency Code"`
		DigitalCurrencyName string `json:"3. Digital Currency Name"`
		MarketCode          string `json:"4. Market Code"`
		MarketName          string `json:"5. Market Name"`
		LastRefreshed       string `json:"6. Last Refreshed"`
		TimeZone            string `json:"7. Time Zone"`
	}

	CryptoTimeSeries map[string]*CryptoDaily

	CryptoDaily struct {
		AOpenCNY     string `json:"1a. open (CNY)"`
		BOpenUSD     string `json:"1b. open (USD)"`
		AHighCNY     string `json:"2a. high (CNY)"`
		BHighUSD     string `json:"2b. high (USD)"`
		ALowCNY      string `json:"3a. low (CNY)"`
		BLowUSD      string `json:"3b. low (USD)"`
		ACloseCNY    string `json:"4a. close (CNY)"`
		BCloseUSD    string `json:"4b. close (USD)"`
		Volume       string `json:"5. volume"`
		MarketCapUSD string `json:"6. market cap (USD)"`
	}
)

func (s *CryptoTimeSeriesDaily) JSON(pretty bool) string {
	var b []byte
	if pretty {
		b, _ = json.MarshalIndent(s, "", "  ")
	} else {
		b, _ = json.Marshal(s)
	}
	return string(b)
}
