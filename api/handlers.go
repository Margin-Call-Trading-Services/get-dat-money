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
		"data":  "",
		"error": err.Error(),
	}
}

func successResponse(arg interface{}) *fiber.Map {
	return &fiber.Map{
		"data":  arg,
		"error": nil,
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

func DefaultTickerInterval() string {
	return "1d"
}

func GetTickerDataHandler(svc Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		p := struct {
			ticker    string
			startDate string
			endDate   string
			interval  string
		}{
			ticker:    strings.ToUpper(c.Query("ticker")),
			startDate: c.Query("start_date", DefaultTickerStartDate()),
			endDate:   c.Query("end_date", DefaultTickerEndDate()),
			interval:  c.Query("interval", DefaultTickerInterval()),
		}

		// Required ticker param
		if p.ticker == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(errMissingTicker))
		}

		// Validate startDate dateOnly format
		if _, err := utils.ParseTimeStringDateOnly(p.startDate); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

		// Validate endDate dateOnly format
		if _, err := utils.ParseTimeStringDateOnly(p.endDate); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

		// Parse interval

		log.Printf("Recieved GET request for %s (%s) from %s to %s", p.ticker, p.interval, p.startDate, p.endDate)

		data, err := svc.GetTickerData(c.Context(), p.ticker, p.startDate, p.endDate, p.interval)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(errorResponse(err))
		}
		log.Printf("Recieved %d price records.", len(data))

		return c.JSON(successResponse(struct {
			Ticker    string         `json:"ticker"`
			PriceData []db.PriceData `json:"price_data"`
		}{
			Ticker:    p.ticker,
			PriceData: data,
		}))
	}
}
