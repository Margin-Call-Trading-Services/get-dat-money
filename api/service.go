package server

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/ryanlattanzi/go-hello-world/objects/db"
	"github.com/ryanlattanzi/go-hello-world/objects/fetchers"
)

type Service interface {
	GetTickerData(ctx context.Context, ticker string, startDate, endDate time.Time, interval string) ([]fetchers.PriceData, error)
}

type service struct {
	db      db.Database
	dbConn  *sql.DB
	fetcher fetchers.DataFetcher
}

func NewService(db db.Database, dbConn *sql.DB, fetcher fetchers.DataFetcher) Service {
	return &service{
		db:      db,
		dbConn:  dbConn,
		fetcher: fetcher,
	}
}

func (s *service) GetTickerData(ctx context.Context, ticker string, startDate, endDate time.Time, interval string) ([]fetchers.PriceData, error) {

	tickerExists, err := s.db.CheckTickerExists(ticker, s.dbConn)
	if err != nil {
		return nil, err
	}

	// TODO: If !tickerExists, we want to save to the DB
	if !tickerExists {
		log.Printf("%s does not exist in db...fetching from external API.", ticker)
		data, err := s.fetcher.GetTickerData(ticker, startDate, endDate, interval)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, errors.New("Did not retrieve any data.")
}
