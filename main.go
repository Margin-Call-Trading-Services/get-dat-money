package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MCTS/get-dat-money/app"
	"github.com/MCTS/get-dat-money/db/postgres"
	"github.com/MCTS/get-dat-money/fetcher/yahoofinance"
	"github.com/MCTS/get-dat-money/server"
)

func main() {

	db := postgres.NewDBFromEnv()
	log.Println("Successfully connected to DB.")

	fetcher := yahoofinance.NewFetcher()

	app := app.New(db, fetcher)

	server := server.New(app)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Gracefully shutting server down....")
		_ = server.Shutdown()
	}()

	if err := server.Listen(":8080"); err != nil {
		log.Panic(err)
	}

	log.Println("Server has been shutdown")
}
