package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchMovieData(title string) (*SearchResult, error) {
	fmt.Println("Fetching Data....")
	resp, err := http.Get(title)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch data: %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch data\nStatus Code:%d\tStatus:%s\n", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read data from response: %v\n", err)
	}

	var poster SearchResult

	err = json.Unmarshal(data, &poster)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal data: %v\n", err)
	}

	return &poster, nil
}
