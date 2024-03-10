package types

import "encoding/json"

type (
	StockTimeSeriesDaily struct {
		Metadata   *Metadata   `json:"Meta Data"`
		TimeSeries *TimeSeries `json:"Time Series (Daily)"`
	}

	Metadata struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		OutputSize    string `json:"4. Output Size"`
		TimeZone      string `json:"5. Time Zone"`
	}

	TimeSeries map[string]*Daily

	Daily struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	}
)

func (s *StockTimeSeriesDaily) JSON(pretty bool) string {
	var b []byte
	if pretty {
		b, _ = json.MarshalIndent(s, "", "  ")
	} else {
		b, _ = json.Marshal(s)
	}
	return string(b)
}
