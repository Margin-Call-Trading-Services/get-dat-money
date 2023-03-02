package server

import (
	"context"
	"time"

	"github.com/ryanlattanzi/go-hello-world/db"
	"github.com/ryanlattanzi/go-hello-world/fetchers"
)

type Service interface {
	GetTickerData(ctx context.Context, ticker string, startDate, endDate time.Time, interval string) ([]fetchers.PriceData, error)
}

type service struct {
	db      db.Database
	dbConn  db.DatabaseConn
	fetcher fetchers.DataFetcher
}

func NewService(db db.Database, dbConn db.DatabaseConn, fetcher fetchers.DataFetcher) Service {
	return &service{
		db:      db,
		dbConn:  dbConn,
		fetcher: fetcher,
	}
}

func (s *service) GetTickerData(ctx context.Context, ticker string, startDate, endDate time.Time, interval string) ([]fetchers.PriceData, error) {

	// TODO: Once the DB part is done, uncomment this for the real logic flow.
	// tickerExists, err := s.db.CheckTickerExists(ticker)
	// if err != nil {
	// 	return nil, err
	// }

	// if !tickerExists {
	// 	data, err := s.fetcher.GetTickerData(ticker, startDate, endDate, interval)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return data, nil
	// }

	data, err := s.fetcher.GetTickerData(ticker, startDate, endDate, interval)
	if err != nil {
		return nil, err
	}
	return data, nil

}
