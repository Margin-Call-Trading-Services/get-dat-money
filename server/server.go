package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/MCTS/get-dat-money/model"
)

type App interface {
	GetTickerData(ctx context.Context, ticker, startDate, endDate, interval string) ([]model.PriceData, error)
}

func New(app App) *fiber.App {
	fbr := fiber.New()
	fbr.Use(logger.New())
	fbr.Use(requestid.New())

	route(fbr, app)

	return fbr
}
