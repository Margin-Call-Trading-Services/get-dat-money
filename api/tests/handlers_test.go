package api_tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	getTickerDataEndpoint = fmt.Sprint(apiGroupName, versionGroupName, getTickerData)
)

func TestGetTickerDataHandler(t *testing.T) {
	var (
		svc = &serviceMock{}
		app = setupTestApp(svc)
	)

	t.Run("Fail_MissingTicker", func(t *testing.T) {

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, getTickerDataEndpoint, nil))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Fail_WrongStartDateFmt", func(t *testing.T) {

		target := fmt.Sprint(getTickerDataEndpoint, "?ticker=AAPL&start_date=20210101")

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, target, nil))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Fail_WrongEndDateFmt", func(t *testing.T) {

		target := fmt.Sprint(getTickerDataEndpoint, "?ticker=AAPL&end_date=01/01/2021")

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, target, nil))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Fail_InvalidInterval", func(t *testing.T) {

		target := fmt.Sprint(getTickerDataEndpoint, "?ticker=AAPL&interval=oneday")

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, target, nil))

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Success_OnlyStartDate", func(t *testing.T) {})
}
