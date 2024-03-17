package server

import "github.com/gofiber/fiber/v2"

const (
	getTickerDataEndpoint = "/api/v1/prices"
)

func route(fbr fiber.Router, app App) {
	fbr.Get(getTickerDataEndpoint, getTickerData(app))
}
