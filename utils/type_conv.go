package utils

import (
	"log"
	"strconv"
	"time"
)

const (
	// Mimicking new time package constant
	DateOnly = "2006-01-02"
)

func StrToFloat(fs string) float64 {
	f, err := strToFloat(fs)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func StrToInt(is string) int64 {
	i, err := strToInt(is)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func strToFloat(fs string) (float64, error) {
	f, err := strconv.ParseFloat(fs, 8)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func strToInt(is string) (int64, error) {
	i, err := strconv.ParseInt(is, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func ParseTimeStringDateOnly(timeStr string) (time.Time, error) {
	t, err := time.Parse(DateOnly, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func ParseTimeToEpoch(t time.Time) int64 {
	epoch := t.Unix()
	return epoch
}
