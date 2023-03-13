package db

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func NewPostgresDatabase(conn *gorm.DB) *PostgresDatabase {
	return &PostgresDatabase{
		dbConn: conn,
	}
}

type PostgresDatabase struct {
	dbConn *gorm.DB
}

func (p *PostgresDatabase) CheckTickerPriceTableExists(table string) (bool, error) {

	query := fmt.Sprint("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'public';")
	rows, err := p.dbConn.Raw(query).Rows()

	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var existingTable string
		err := rows.Scan(&existingTable)
		if err != nil {
			return false, err
		}
		if existingTable == table {
			log.Printf("%s price table already exists...fetching from database.", table)
			return true, nil
		}
	}
	log.Printf("%s does not exist in db...fetching from external API.", table)
	return false, nil
}

func (p *PostgresDatabase) CreateTickerPriceTable(table string) error {

	err := p.dbConn.Table(table).AutoMigrate(&PriceData{})
	if err != nil {
		return err
	}
	log.Printf("Successfully created price table %s", table)
	return nil
}

func (p *PostgresDatabase) BulkUploadPriceData(table string, priceData []PriceData) error {

	if err := p.dbConn.Table(table).Create(&priceData).Error; err != nil {
		return err
	}

	log.Printf("Successfully uploaded %d price records to %s", len(priceData), table)
	return nil
}

func (p *PostgresDatabase) GetDataBetweenDates(table, startDate, endDate string) ([]PriceData, error) {

	var priceData []PriceData
	query := fmt.Sprintf("%s BETWEEN ? AND ?", colDate)

	if err := p.dbConn.Table(table).Where(query, startDate, endDate).Find(&priceData).Error; err != nil {
		return nil, err
	}

	log.Printf("Successfully retrieved data for %s from %s to %s", table, startDate, endDate)
	return priceData, nil
}
