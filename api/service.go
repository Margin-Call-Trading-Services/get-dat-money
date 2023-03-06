package api

import (
	"context"

	"github.com/ryanlattanzi/get-dat-money/objects/db"
	"github.com/ryanlattanzi/get-dat-money/objects/fetchers"
)

type Service interface {
	GetTickerData(ctx context.Context, ticker, startDate, endDate, interval string) ([]db.PriceData, error)
}

type service struct {
	db      db.Database
	fetcher fetchers.DataFetcher
}

func NewService(db db.Database, fetcher fetchers.DataFetcher) Service {
	return &service{
		db:      db,
		fetcher: fetcher,
	}
}

func (s *service) GetTickerData(ctx context.Context, ticker, startDate, endDate, interval string) ([]db.PriceData, error) {

	tickerExists, err := s.db.CheckTickerPriceTableExists(ticker)
	if err != nil {
		return nil, err
	}

	if !tickerExists {
		data, err := s.fetcher.GetTickerData(
			ticker,
			DefaultTickerStartDate(),
			DefaultTickerEndDate(),
			interval,
		)
		if err != nil {
			return nil, err
		}

		if err := s.db.CreateTickerPriceTable(ticker); err != nil {
			return nil, err
		}

		if err := s.db.BulkUploadPriceData(ticker, data); err != nil {
			return nil, err
		}
	}

	data, err := s.db.GetDataBetweenDates(ticker, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return data, nil
}
