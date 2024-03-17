package server

import (
	"time"

	"github.com/MCTS/get-dat-money/utils"
)

func defaultTickerStartDate() string {
	return "1900-01-01"
}

func defaultTickerEndDate() string {
	now := time.Now()
	today := now.Format(utils.DateOnly)
	return today
}

func defaultTickerInterval() string {
	return "1d"
}
