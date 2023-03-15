package api_tests

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/MCTS/get-dat-money/objects/db"
)

type serviceMock struct {
	mock.Mock
}

func (sm *serviceMock) GetTickerData(ctx context.Context, ticker, startDate, endDate, interval string) ([]db.PriceData, error) {
	args := sm.Called(ctx, ticker, startDate, endDate, interval)

	return args.Get(0).([]db.PriceData), args.Error(1)
}
