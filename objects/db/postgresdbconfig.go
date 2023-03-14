package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/MCTS/get-dat-money/utils"
)

func NewPostgresConfig() PostgresDatabaseConfig {
	return PostgresDatabaseConfig{
		host:      os.Getenv("POSTGRES_HOST"),
		port:      os.Getenv("POSTGRES_PORT"),
		user:      os.Getenv("POSTGRES_USER"),
		password:  os.Getenv("POSTGRES_PASSWORD"),
		dbname:    os.Getenv("POSTGRES_DB"),
		batchSize: os.Getenv("POSTGRES_BATCH_SIZE"),
	}
}

type PostgresDatabaseConfig struct {
	host      string
	port      string
	user      string
	password  string
	dbname    string
	batchSize string
}

func (pgcfg PostgresDatabaseConfig) Connect() *gorm.DB {
	connStr := pgcfg.GetConnectionStr()
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: priceTableSchema + ".",
		},
		CreateBatchSize: int(utils.StrToInt(pgcfg.batchSize)),
	})
	if err != nil {
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
