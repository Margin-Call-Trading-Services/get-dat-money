package db

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/ryanlattanzi/go-hello-world/objects/fetchers"
)

type Database interface {
	CheckTickerExists(ticker string) (bool, error)
	GetDataBetweenDates(t string, startDate, endDate time.Time) ([]fetchers.PriceData, error)
}
