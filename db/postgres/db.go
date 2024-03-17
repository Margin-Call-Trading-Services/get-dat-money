package postgres

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func NewDBFromEnv() *Database {
	cfg := newConfigFromEnv()
	conn := cfg.Connect()

	return &Database{
		conn: conn,
	}
}

type Database struct {
	conn *gorm.DB
}
