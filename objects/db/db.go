package db

import (
	"time"

	_ "github.com/lib/pq"
)

type PriceData struct {
	Date     string  `json:"date" gorm:"primaryKey" gorm:"index"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	AdjClose float64 `json:"adj_close"`
	Volume   int64   `json:"volume"`
}

type Database interface {
	CheckTickerExists(ticker string) (bool, error)
	GetDataBetweenDates(t string, startDate, endDate time.Time) ([]PriceData, error)
}
