package alphavantage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/triskaidekaphile/go-marketdata/apis/alphavantage/types"
	"github.com/triskaidekaphile/go-marketdata/httpext"
	"io"
	"net/http"
)

type (
	API interface {
		GetStockTimeSeriesDaily(ctx context.Context, symbol string) (*types.StockTimeSeriesDaily, error)
		GetCryptoTimeSeriesDaily(ctx context.Context, symbol string) (*types.CryptoTimeSeriesDaily, error)
	}

	api struct {
		baseURL string
		apiKey  string
		verbose bool
	}
)

func New(baseURL, apiKey string, verbose bool) API {
	return &api{
		baseURL: baseURL,
		apiKey:  apiKey,
		verbose: verbose,
	}
}

func (a *api) GetStockTimeSeriesDaily(ctx context.Context, symbol string) (*types.StockTimeSeriesDaily, error) {
	uri := fmt.Sprintf("%s/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", a.baseURL, symbol, a.apiKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	httpext.Dump(req, res, a.verbose)

	if res == nil || res.Body == nil {
		return nil, errors.New("unable to get time series daily: nullable response")
	}

	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable tot get time series daily: status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data *types.StockTimeSeriesDaily
	err = json.Unmarshal(body, &data)
	return data, err
}

func (a *api) GetCryptoTimeSeriesDaily(ctx context.Context, symbol string) (*types.CryptoTimeSeriesDaily, error) {
	uri := fmt.Sprintf("%s/query?function=DIGITAL_CURRENCY_DAILY&symbol=%s&market=CNY&apikey=%s", a.baseURL, symbol, a.apiKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	httpext.Dump(req, res, a.verbose)

	if res == nil || res.Body == nil {
		return nil, errors.New("unable to get time series digital currency daily: nullable response")
	}

	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable tot get time series digital currency daily: status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data *types.CryptoTimeSeriesDaily
	err = json.Unmarshal(body, &data)
	return data, err
}
