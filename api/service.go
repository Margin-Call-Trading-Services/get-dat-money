package api

import (
	"context"
	"fmt"

	"github.com/MCTS/get-dat-money/objects/db"
	"github.com/MCTS/get-dat-money/objects/fetchers"
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

	table := fmt.Sprintf("%s_%s", ticker, interval)
	tickerExists, err := s.db.CheckTickerPriceTableExists(table)
	if err != nil {
		return nil, err
	}

	if !tickerExists {
		data, err := s.fetcher.GetTickerData(
			ticker,
			defaultTickerStartDate(),
			defaultTickerEndDate(),
			interval,
		)
		if err != nil {
			return nil, err
		}

		if err := s.db.CreateTickerPriceTable(table); err != nil {
			return nil, err
		}

		if err := s.db.BulkUploadPriceData(table, data); err != nil {
			return nil, err
		}
	}

	data, err := s.db.GetDataBetweenDates(table, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return data, nil
}
