package app

import (
	"context"
	"fmt"

	"github.com/MCTS/get-dat-money/model"
)

type PriceDatabase interface {
	CheckTickerPriceTableExists(table string) (bool, error)
	CreateTickerPriceTable(table string) error
	BulkUploadPriceData(table string, priceData []model.PriceData) error
	GetDataBetweenDates(table, startDate, endDate string) ([]model.PriceData, error)
}

type DataFetcher interface {
	GetTickerData(ticker, starDate, endDate, interval string) ([]model.PriceData, error)
}

type App struct {
	db      PriceDatabase
	fetcher DataFetcher
}

func New(db PriceDatabase, fetcher DataFetcher) *App {
	return &App{
		db:      db,
		fetcher: fetcher,
	}
}

func (a *App) GetTickerData(ctx context.Context, ticker, startDate, endDate, interval string) ([]model.PriceData, error) {

	table := fmt.Sprintf("%s_%s", ticker, interval)
	tickerExists, err := a.db.CheckTickerPriceTableExists(table)
	if err != nil {
		return nil, err
	}

	if !tickerExists {
		data, err := a.fetcher.GetTickerData(
			ticker,
			startDate,
			endDate,
			interval,
		)
		if err != nil {
			return nil, err
		}

		if err := a.db.CreateTickerPriceTable(table); err != nil {
			return nil, err
		}

		if err := a.db.BulkUploadPriceData(table, data); err != nil {
			return nil, err
		}
	}

	data, err := a.db.GetDataBetweenDates(table, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return data, nil
}
