package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/MCTS/get-dat-money/objects/db"
)

func GetTickerDataHandler(svc Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		p := struct {
			ticker    string
			startDate string
			endDate   string
			interval  string
		}{
			ticker:    strings.ToUpper(c.Query("ticker")),
			startDate: c.Query("start_date", defaultTickerStartDate()),
			endDate:   c.Query("end_date", defaultTickerEndDate()),
			interval:  c.Query("interval", defaultTickerInterval()),
		}

		// Required ticker param
		if p.ticker == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(errMissingTicker))
		}

		if err := validateDateOnlyFmt(p.startDate); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

		if err := validateDateOnlyFmt(p.endDate); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

		if err := validateInterval(p.interval); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(errorResponse(err))
		}

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
