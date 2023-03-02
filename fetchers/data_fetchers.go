package fetchers

import (
	"fmt"
	"time"

	"github.com/ryanlattanzi/go-hello-world/utils"
)

type PriceData struct {
	Date     string  `json:"date"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	AdjClose float64 `json:"adj_close"`
	Volume   int64   `json:"volume"`
}

type DataFetcher interface {
	GetTickerData(ticker string, starDate, endDate time.Time, interval string) ([]PriceData, error)
}

func NewYahooFinanceFetcher() DataFetcher {
	return YahooFinanceFetcher{}
}

type YahooFinanceFetcher struct{}

func (yff YahooFinanceFetcher) GetTickerData(ticker string, starDate, endDate time.Time, interval string) ([]PriceData, error) {
	url, err := yff.buildUrl(ticker, starDate, endDate, interval)
	if err != nil {
		return nil, err
	}

	data, err := utils.ReadCSVFromUrl(url)
	if err != nil {
		return nil, err
	}

	var priceData []PriceData

	for i, d := range data {
		if i == 0 {
			continue
		}
		priceData = append(priceData, PriceData{
			Date:     d[0],
			Open:     utils.StrToFloat(d[1]),
			High:     utils.StrToFloat(d[2]),
			Low:      utils.StrToFloat(d[3]),
			Close:    utils.StrToFloat(d[4]),
			AdjClose: utils.StrToFloat(d[5]),
			Volume:   utils.StrToInt(d[6]),
		})
	}

	return priceData, nil
}

func (yff YahooFinanceFetcher) buildUrl(ticker string, starDate, endDate time.Time, interval string) (string, error) {
	startEpoch := utils.ParseTimeToEpoch(starDate)
	endEpoch := utils.ParseTimeToEpoch(endDate)

	url := fmt.Sprintf(
		"https://query1.finance.yahoo.com/v7/finance/download/%s"+
			"?period1=%d&period2=%d&interval=%s&events=history"+
			"&includeAdjustedClose=true", ticker, startEpoch, endEpoch, interval,
	)

	return url, nil
}
