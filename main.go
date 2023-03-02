package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/ryanlattanzi/go-hello-world/db"
	"github.com/ryanlattanzi/go-hello-world/fetchers"
	"github.com/ryanlattanzi/go-hello-world/server"
)

func main() {

	db := db.NewPostgresDatabase()
	fetcher := fetchers.NewYahooFinanceFetcher()
	svc := server.NewService()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(requestid.New())

	// api := app.Group("/api")
	// v1 := api.Group("/v1")

	// pgConfig := db.PostgresConfig()
	// pgDB := db.PostgresDatabase{
	// 	Config: pgConfig,
	// }

	// ticker := "AAPL"
	// interval := "1d"
	// start := "1900-01-01"
	// end := "2023-01-01"

	// fmt.Printf("Collected %d days of price data for %s\n", len(priceData), ticker)
	// fmt.Printf("Sample data: %s, %f, %f\n", priceData[10].Date, priceData[10].Open, priceData[10].Close)
}
