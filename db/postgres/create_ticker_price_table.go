package postgres

import (
	"log"

	"github.com/MCTS/get-dat-money/model"
)

func (db *Database) CreateTickerPriceTable(table string) error {

	err := db.conn.Table(table).AutoMigrate(&model.PriceData{})
	if err != nil {
		return err
	}
	log.Printf("Successfully created price table %s", table)
	return nil
}
