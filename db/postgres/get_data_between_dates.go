package postgres

import (
	"fmt"
	"log"

	"github.com/MCTS/get-dat-money/model"
)

func (db *Database) GetDataBetweenDates(table, startDate, endDate string) ([]model.PriceData, error) {

	var priceData []model.PriceData
	query := fmt.Sprintf("%s BETWEEN ? AND ?", model.ColDate)

	if err := db.conn.Table(table).Where(query, startDate, endDate).Find(&priceData).Error; err != nil {
		return nil, err
	}

	log.Printf("Successfully retrieved data for %s from %s to %s", table, startDate, endDate)
	return priceData, nil
}
