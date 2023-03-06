package db

import (
	_ "github.com/lib/pq"
)

const (
	priceTableSchema = "price_data"

	colDate     = "date"
	colOpen     = "open"
	colHigh     = "high"
	colLow      = "low"
	colClose    = "close"
	colAdjClose = "adj_close"
	colVolume   = "volume"
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
	CheckTickerPriceTableExists(ticker string) (bool, error)
	CreateTickerPriceTable(ticker string) error
	BulkUploadPriceData(ticker string, priceData []PriceData) error
	GetDataBetweenDates(ticker, startDate, endDate string) ([]PriceData, error)
}
