package utils

import (
	"encoding/csv"
	"net/http"
)

func ReadCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}
