package db

import "gorm.io/gorm"

type DatabaseConfig interface {
	Connect() *gorm.DB
}
