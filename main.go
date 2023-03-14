package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/MCTS/get-dat-money/api"
	"github.com/MCTS/get-dat-money/objects/db"
	"github.com/MCTS/get-dat-money/objects/fetchers"
)

func main() {

	dbCfg := db.NewPostgresConfig()
	dbConn := dbCfg.Connect()

	db := db.NewPostgresDatabase(dbConn)
	log.Println("Successfully connected to DB.")

	fetcher := fetchers.NewYahooFinanceFetcher()

	svc := api.NewService(db, fetcher)

	app := fiber.New()

	app.Use(logger.New())
	app.Use(requestid.New())

	apiGroup := app.Group("/api")
	v1Group := apiGroup.Group("/v1")

	api.ApiRouter(v1Group, svc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-c
		log.Println("Gracefully shutting server down....")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":8080"); err != nil {
		log.Panic(err)
	}

	log.Println("Server has been shutdown")
}
