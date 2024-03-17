package yahoofinance

import (
	"fmt"
	"log"

	"github.com/MCTS/get-dat-money/model"
	"github.com/MCTS/get-dat-money/utils"
)

func (yff YahooFinanceFetcher) GetTickerData(ticker, starDate, endDate, interval string) ([]model.PriceData, error) {
	url, err := yff.buildUrl(ticker, starDate, endDate, interval)
	if err != nil {
		return nil, err
	}

	data, err := utils.ReadCSVFromUrl(url)
	if err != nil {
		return nil, err
	}

	var priceData []model.PriceData

	for i, d := range data {
		if i == 0 {
			continue
		}
		priceData = append(priceData, model.PriceData{
			Date:     d[0],
			Open:     utils.StrToFloat(d[1]),
			High:     utils.StrToFloat(d[2]),
			Low:      utils.StrToFloat(d[3]),
			Close:    utils.StrToFloat(d[4]),
			AdjClose: utils.StrToFloat(d[5]),
			Volume:   utils.StrToInt(d[6]),
		})
	}

	log.Printf("Successfully retrieved full historical data for %s", ticker)
	return priceData, nil
}

func (yff YahooFinanceFetcher) buildUrl(ticker, starDate, endDate, interval string) (string, error) {

	// Skipping error handling since this was already validated in api.GetTickerDataHandler()
	startDateTime, _ := utils.ParseTimeStringDateOnly(starDate)
	endDateTime, _ := utils.ParseTimeStringDateOnly(endDate)

	startEpoch := utils.ParseTimeToEpoch(startDateTime)
	endEpoch := utils.ParseTimeToEpoch(endDateTime)

	url := fmt.Sprintf(
		"https://query1.finance.yahoo.com/v7/finance/download/%s"+
			"?period1=%d&period2=%d&interval=%s&events=history"+
			"&includeAdjustedClose=true", ticker, startEpoch, endEpoch, interval,
	)

	return url, nil
}
