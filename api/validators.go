package api

import (
	"fmt"

	"github.com/MCTS/get-dat-money/utils"
)

func validateDateOnlyFmt(timeStr string) error {
	_, err := utils.ParseTimeStringDateOnly(timeStr)
	if err != nil {
		return fmt.Errorf("invalid date: `%s`. ensure fmt YYYY-MM-DD", timeStr)
	}
	return nil
}

func validateInterval(interval string) error {
	valid := []string{
		"1d",
	}

	c := utils.SliceContains(valid, interval)
	if c {
		return nil
	}
	return fmt.Errorf("invalid interval. valid values are %v", valid)
}
