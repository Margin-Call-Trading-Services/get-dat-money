package api

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/ryanlattanzi/get-dat-money/objects/db"
	"github.com/ryanlattanzi/get-dat-money/utils"
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

func DefaultTickerStartDate() string {
	return "1900-01-01"
}

func DefaultTickerEndDate() string {
	currentTime := time.Now()
	today := currentTime.Format(utils.DateOnly)
	return today
}

func GetTickerDataHandler(svc Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// TODO: put the API params in a struct or something for validation?

		type resp struct {
			Ticker string         `json:"ticker"`
			Data   []db.PriceData `json:"data"`
		}

		// Parse ticker
		ticker := strings.ToUpper(c.Query("ticker"))
		if ticker == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(errMissingTicker))
		}

		// Parse start; validating dateOnly format
		start := c.Query("start_date", DefaultTickerStartDate())
		if _, err := utils.ParseTimeStringDateOnly(start); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

		// Parse end; validating dateOnly format
		end := c.Query("end_date", DefaultTickerEndDate())
		if _, err := utils.ParseTimeStringDateOnly(end); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

		// Parse interval
		interval := c.Query("interval", "1d")

		log.Printf("Recieved GET request for %s from %s to %s", ticker, start, end)

		data, err := svc.GetTickerData(c.Context(), ticker, start, end, interval)
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
