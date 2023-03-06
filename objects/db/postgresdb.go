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

func (p *PostgresDatabase) CheckTickerPriceTableExists(ticker string) (bool, error) {

	query := fmt.Sprint("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'public';")
	rows, err := p.dbConn.Raw(query).Rows()

	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var existingTicker string
		err := rows.Scan(&existingTicker)
		if err != nil {
			return false, err
		}
		if existingTicker == ticker {
			log.Printf("%s price table already exists...fetching from database.", ticker)
			return true, nil
		}
	}
	log.Printf("%s does not exist in db...fetching from external API.", ticker)
	return false, nil
}

func (p *PostgresDatabase) CreateTickerPriceTable(ticker string) error {

	err := p.dbConn.Table(ticker).AutoMigrate(&PriceData{})
	p.dbConn.Migrator().CreateTable()
	if err != nil {
		return err
	}
	log.Printf("Successfully created price table %s", ticker)
	return nil
}

func (p *PostgresDatabase) BulkUploadPriceData(ticker string, priceData []PriceData) error {

	if err := p.dbConn.Table(ticker).Create(&priceData).Error; err != nil {
		return err
	}

	log.Printf("Successfully uploaded %d price records to %s", len(priceData), ticker)
	return nil
}

func (p *PostgresDatabase) GetDataBetweenDates(ticker, startDate, endDate string) ([]PriceData, error) {

	var priceData []PriceData
	query := fmt.Sprintf("%s BETWEEN ? AND ?", colDate)

	if err := p.dbConn.Table(ticker).Where(query, startDate, endDate).Find(&priceData).Error; err != nil {
		return nil, err
	}

	log.Printf("Successfully retrieved data for %s from %s to %s", ticker, startDate, endDate)
	return priceData, nil
}
