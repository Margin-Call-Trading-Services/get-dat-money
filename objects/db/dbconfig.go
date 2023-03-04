package db

import (
	"database/sql"
)

type DatabaseConfig interface {
	Connect() *sql.DB
}
