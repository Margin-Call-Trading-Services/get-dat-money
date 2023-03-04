package db

import (
	"fmt"
	"time"

	"github.com/ryanlattanzi/go-hello-world/utils"
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

func (p *PostgresDatabase) CheckTickerExists(t string) (bool, error) {

	rows, err := p.dbConn.Raw("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'public';").Rows()

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
		if existingTicker == t {
			return true, nil
		}
	}

	return false, nil
}

func (p *PostgresDatabase) GetDataBetweenDates(t string, startDate, endDate time.Time) ([]PriceData, error) {

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE %s BETWEEN %s AND %s;",
		t,
		"date",
		startDate.Format(utils.DateOnly),
		endDate.Format(utils.DateOnly),
	)

	rows, err := p.dbConn.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var priceData []PriceData

	for rows.Next() {
		var pd PriceData
		err := rows.Scan(
			&pd.Date,
			&pd.Open,
			&pd.High,
			&pd.Low,
			&pd.Close,
			&pd.AdjClose,
			&pd.Volume,
		)
		if err != nil {
			return nil, err
		}
		priceData = append(priceData, pd)
	}

	return priceData, nil
}
