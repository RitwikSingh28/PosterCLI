package utils

import "net/url"

func ParseMovieTitle(title []string) string {
	movieTitle := ""
	for _, partTitle := range title {
		movieTitle += partTitle + " "
	}
	movieTitle = movieTitle[:len(movieTitle)-1]
	encodedTitle := url.QueryEscape(movieTitle)
	url := "http://www.omdbapi.com/?s=" + encodedTitle + "&apikey=" + apiKey
	return url
}
