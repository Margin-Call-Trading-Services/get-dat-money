package db

import (
	"time"
	"database/sql"

	_ "github.com/lib/pq"
)

type DatabaseConn *sql.DB

type Database interface {
	Connect(cs string) (DatabaseConn, error)
	CheckTickerExists(ticker string) (bool, error)
}

func NewPostgresDatabase(cfg DatabaseConfig) Database {
	return &PostgresDatabase{
		Config: cfg,
	}
}

type PostgresDatabase struct {
	Config DatabaseConfig
}

func (p *PostgresDatabase) Connect(cs string) (DatabaseConn, error) {
	connStr := p.Config.GetConnectionStr()
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db, nil
}

func (p *PostgresDatabase) CheckTickerExists(t string, db DatabaseConn) (bool, error) {

	type ticker struct {
		name string
	}

	rows, err := db.Query(
		"SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'public';"
	)

	if err != nil {
        return nil, err
    }
    defer rows.Close()

	for rows.Next() {
        var existingTicker ticker
        err := rows.Scan(&existingTicker.name)
		if err != nil {
            return nil, err
        }
		if existingTicker.name == t {
			return true, nil
		}
    }

	return false, nil
}

func (p *PostgresDatabase) GetData(t string, startDate, endDate time.Time) {}