package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/MCTS/get-dat-money/model"
)

func getTickerData(app App) fiber.Handler {
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
			return c.JSON(model.ErrorResponse(model.ErrMissingTicker))
		}

		if err := validateDateOnlyFmt(p.startDate); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(model.ErrorResponse(err))
		}

		if err := validateDateOnlyFmt(p.endDate); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(model.ErrorResponse(err))
		}

		if err := validateInterval(p.interval); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(model.ErrorResponse(err))
		}

		log.Printf("Recieved GET request for %s (%s) from %s to %s", p.ticker, p.interval, p.startDate, p.endDate)

		data, err := app.GetTickerData(c.Context(), p.ticker, p.startDate, p.endDate, p.interval)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(model.ErrorResponse(err))
		}
		log.Printf("Recieved %d price records.", len(data))

		return c.JSON(model.SuccessResponse(struct {
			Ticker    string            `json:"ticker"`
			PriceData []model.PriceData `json:"price_data"`
		}{
			Ticker:    p.ticker,
			PriceData: data,
		}))
	}
}
