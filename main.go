package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	server "github.com/ryanlattanzi/go-hello-world/api"
	"github.com/ryanlattanzi/go-hello-world/objects/db"
	"github.com/ryanlattanzi/go-hello-world/objects/fetchers"
)

func main() {

	dbCfg := db.NewPostgresConfig()
	dbConn := dbCfg.Connect()

	db := db.NewPostgresDatabase(dbConn)
	log.Println("Established PostgresDatabase struct.")

	log.Println("Successfully connected to DB.")
	defer dbConn.Close()

	fetcher := fetchers.NewYahooFinanceFetcher()

	svc := server.NewService(db, dbConn, fetcher)

	app := fiber.New()

	app.Use(logger.New())
	app.Use(requestid.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	server.ApiRouter(v1, svc)

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
