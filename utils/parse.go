package utils

import (
	"fmt"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func ParseMovieTitle(title []string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", fmt.Errorf("couldn't load the API Key: %v", err)
	}
	apiKey := os.Getenv("API_KEY")

	movieTitle := ""
	for _, partTitle := range title {
		movieTitle += partTitle + " "
	}
	movieTitle = movieTitle[:len(movieTitle)-1]
	encodedTitle := url.QueryEscape(movieTitle)
	url := "http://www.omdbapi.com/?s=" + encodedTitle + "&apikey=" + apiKey
	return url, nil
}
