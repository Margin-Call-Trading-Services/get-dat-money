package server

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/ryanlattanzi/go-hello-world/objects/db"
	"github.com/ryanlattanzi/go-hello-world/utils"
)

var (
	errMissingTicker = errors.New("Missing ticker param.")
)

func errorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func successResponse(arg interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   arg,
		"error":  nil,
	}
}

func GetTickerDataHandler(svc Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// TODO: put the API params in a struct or something for validation?

		type resp struct {
			Ticker string         `json:"ticker"`
			Data   []db.PriceData `json:"data"`
		}

		// Parse ticker
		ticker := c.Query("ticker")
		if ticker == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(errMissingTicker))
		}

		// Parse start
		start := c.Query("start_date", "1900-01-01")
		startDate, err := utils.ParseTimeStringDateOnly(start)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

		// Parse end
		currentTime := time.Now()
		today := currentTime.Format(utils.DateOnly)
		end := c.Query("end_date", today)
		endDate, err := utils.ParseTimeStringDateOnly(end)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

		// Parse interval
		interval := c.Query("interval", "1d")

		log.Printf("Recieved GET request for %s from %s to %s", ticker, start, end)

		data, err := svc.GetTickerData(c.Context(), ticker, startDate, endDate, interval)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(errorResponse(err))
		}
		log.Printf("Recieved %d price records.", len(data))

		return c.JSON(successResponse(resp{
			Ticker: ticker,
			Data:   data,
		}))
	}
}
