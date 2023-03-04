package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/ryanlattanzi/go-hello-world/objects/fetchers"
	"github.com/ryanlattanzi/go-hello-world/utils"
)

type Database interface {
	CheckTickerExists(ticker string) (bool, error)
	GetDataBetweenDates(t string, startDate, endDate time.Time) ([]fetchers.PriceData, error)
}

func NewPostgresDatabase(conn *sql.DB) *PostgresDatabase {
	return &PostgresDatabase{
		dbConn: conn,
	}
}

type PostgresDatabase struct {
	dbConn *sql.DB
}

func (p *PostgresDatabase) CheckTickerExists(t string) (bool, error) {

	rows, err := p.dbConn.Query("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'public';")

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

func (p *PostgresDatabase) GetDataBetweenDates(t string, startDate, endDate time.Time) ([]fetchers.PriceData, error) {

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE %s BETWEEN %s AND %s;",
		t,
		fetchers.DateCol,
		startDate.Format(utils.DateOnly),
		endDate.Format(utils.DateOnly),
	)

	rows, err := p.dbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var priceData []fetchers.PriceData

	for rows.Next() {
		var pd fetchers.PriceData
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
