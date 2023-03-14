package api

import (
	"fmt"

	"github.com/MCTS/get-dat-money/utils"
)

func validateDateOnlyFmt(timeStr string) error {
	_, err := utils.ParseTimeStringDateOnly(timeStr)
	if err != nil {
		return err
	}
	return nil
}

func validateInterval(interval string) error {
	valid := []string{
		"1m",
		"2m",
		"5m",
		"15m",
		"30m",
		"60m",
		"90m",
		"1h",
		"1d",
		"5d",
		"1wk",
		"1mo",
		"3mo",
	}

	c := utils.SliceContains(valid, interval)
	if c {
		return nil
	}
	return fmt.Errorf("invalid interval. valid values are %v", valid)
}
