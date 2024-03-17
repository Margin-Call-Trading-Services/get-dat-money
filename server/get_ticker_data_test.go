package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/MCTS/get-dat-money/model"
)

type appMock struct {
	mock.Mock
}

func (sm *appMock) GetTickerData(ctx context.Context, ticker, startDate, endDate, interval string) ([]model.PriceData, error) {
	args := sm.Called(ctx, ticker, startDate, endDate, interval)

	return args.Get(0).([]model.PriceData), args.Error(1)
}

func setupTestApp() (*appMock, *fiber.App) {
	app := &appMock{}
	server := fiber.New()
	route(server, app)

	return app, server
}

func TestGetTickerDataHandler(t *testing.T) {
	var (
		app, server = setupTestApp()
	)

	t.Run("Fail_MissingTicker", func(t *testing.T) {

		resp, err := server.Test(httptest.NewRequest(http.MethodGet, getTickerDataEndpoint, nil))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Fail_WrongStartDateFmt", func(t *testing.T) {

		target := fmt.Sprint(getTickerDataEndpoint, "?ticker=AAPL&start_date=20210101")

		resp, err := server.Test(httptest.NewRequest(http.MethodGet, target, nil))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Fail_WrongEndDateFmt", func(t *testing.T) {

		target := fmt.Sprint(getTickerDataEndpoint, "?ticker=AAPL&end_date=01/01/2021")

		resp, err := server.Test(httptest.NewRequest(http.MethodGet, target, nil))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Fail_InvalidInterval", func(t *testing.T) {

		target := fmt.Sprint(getTickerDataEndpoint, "?ticker=AAPL&interval=oneday")

		resp, err := server.Test(httptest.NewRequest(http.MethodGet, target, nil))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Success_OnlyStartDate", func(t *testing.T) {

		app.On(
			"GetTickerData",
			mock.AnythingOfType("*fasthttp.RequestCtx"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return([]model.PriceData{}, nil)

		target := fmt.Sprint(getTickerDataEndpoint, "?ticker=AAPL")

		resp, err := server.Test(httptest.NewRequest(http.MethodGet, target, nil))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
