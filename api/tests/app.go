package api_tests

import (
	"github.com/MCTS/get-dat-money/api"
	"github.com/gofiber/fiber/v2"
)

const (
	apiGroupName     = "/api"
	versionGroupName = "/v1"

	getTickerData = "/prices"
)

func setupTestApp(svc api.Service) *fiber.App {
	app := fiber.New()
	apiGroup := app.Group(apiGroupName)
	versionGroup := apiGroup.Group(versionGroupName)

	// Every new endpoint and associated handler needs to be added here.
	versionGroup.Get(getTickerData, api.GetTickerDataHandler(svc))

	return app
}
