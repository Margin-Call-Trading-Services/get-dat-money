package postgres

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/MCTS/get-dat-money/model"
	"github.com/MCTS/get-dat-money/utils"
)

func newConfigFromEnv() databaseConfig {
	return databaseConfig{
		host:      os.Getenv("POSTGRES_HOST"),
		port:      os.Getenv("POSTGRES_PORT"),
		user:      os.Getenv("POSTGRES_USER"),
		password:  os.Getenv("POSTGRES_PASSWORD"),
		dbname:    os.Getenv("POSTGRES_DB"),
		batchSize: os.Getenv("POSTGRES_BATCH_SIZE"),
	}
}

type databaseConfig struct {
	host      string
	port      string
	user      string
	password  string
	dbname    string
	batchSize string
}

func (cfg databaseConfig) Connect() *gorm.DB {
	connStr := cfg.getConnectionStr()
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: model.PriceTableSchema + ".",
		},
		CreateBatchSize: int(utils.StrToInt(cfg.batchSize)),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func (cfg databaseConfig) getConnectionStr() string {
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.host, cfg.port, cfg.user, cfg.password, cfg.dbname,
	)
	return connStr
}
