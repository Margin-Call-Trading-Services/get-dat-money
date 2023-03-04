package db

import (
	"database/sql"
	"fmt"
	"os"
)

type DatabaseConfig interface {
	Connect() *sql.DB
}

func NewPostgresConfig() PostgresDatabaseConfig {
	return PostgresDatabaseConfig{
		host:     os.Getenv("POSTGRES_HOST"),
		port:     os.Getenv("POSTGRES_PORT"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbname:   os.Getenv("POSTGRES_DB"),
	}
}

type PostgresDatabaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func (pgcfg PostgresDatabaseConfig) Connect() *sql.DB {
	connStr := pgcfg.GetConnectionStr()
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}

func (pgcfg PostgresDatabaseConfig) GetConnectionStr() string {
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pgcfg.host, pgcfg.port, pgcfg.user, pgcfg.password, pgcfg.dbname,
	)
	return connStr
}
