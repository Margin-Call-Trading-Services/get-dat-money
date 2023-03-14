package api

import (
	"time"

	"github.com/MCTS/get-dat-money/utils"
)

func defaultTickerStartDate() string {
	return "1900-01-01"
}

func defaultTickerEndDate() string {
	currentTime := time.Now()
	today := currentTime.Format(utils.DateOnly)
	return today
}

func defaultTickerInterval() string {
	return "1d"
}
