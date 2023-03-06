package api

import (
	"github.com/gofiber/fiber/v2"
)

func ApiRouter(app fiber.Router, svc Service) {
	app.Get("/prices", GetTickerDataHandler(svc))
}
