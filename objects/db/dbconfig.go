package db

import (
	"fmt"
	"os"
)

type DatabaseConfig interface {
	GetConnectionStr() string
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

func (pgcfg PostgresDatabaseConfig) GetConnectionStr() string {
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pgcfg.host, pgcfg.port, pgcfg.user, pgcfg.password, pgcfg.dbname,
	)
	return connStr
}
