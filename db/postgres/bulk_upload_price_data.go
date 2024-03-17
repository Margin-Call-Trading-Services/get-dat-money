package postgres

import (
	"log"

	"github.com/MCTS/get-dat-money/model"
)

func (db *Database) BulkUploadPriceData(table string, priceData []model.PriceData) error {

	if err := db.conn.Table(table).Create(&priceData).Error; err != nil {
		return err
	}

	log.Printf("Successfully uploaded %d price records to %s", len(priceData), table)
	return nil
}
