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
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data Status Code:%d\tStatus:%s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from response: %v ", err)
	}

	var poster SearchResult

	err = json.Unmarshal(data, &poster)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %v ", err)
	}

	return &poster, nil
}
