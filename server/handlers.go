package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanlattanzi/go-hello-world/fetchers"
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

		type resp struct {
			ticker string               `json:"ticker"`
			data   []fetchers.PriceData `json:"data"`
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
		today := currentTime.Format(time.DateOnly)
		end := c.Query("end_date", today)
		endDate, err := utils.ParseTimeStringDateOnly(end)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

		// Parse interval
		interval := c.Query("interval", "1d")

		data, err := svc.GetTickerData(c, ticker, startDate, endDate, interval)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(errorResponse(err))
		}

		return c.JSON(successResponse(resp{
			ticker: ticker,
			data:   data,
		}))
	}
}
