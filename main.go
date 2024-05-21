package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	Poster string `json:"Poster"`
}

type SearchResult struct {
	Search []Movie `json:"Search"`
}

const apiKey = "1b8a3f2d"

func parseMovieTitle(title []string) string {
	movieTitle := ""
	for _, partTitle := range title {
		movieTitle += partTitle + " "
	}
	movieTitle = movieTitle[:len(movieTitle)-1]
	encodedTitle := url.QueryEscape(movieTitle)
	url := "http://www.omdbapi.com/?s=" + encodedTitle + "&apikey=" + apiKey
	return url

}

func fetchMovieData(title string) *SearchResult {
	fmt.Println("Fetching Data....")
	resp, err := http.Get(title)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status error code: %d %s\n", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var poster SearchResult

	err = json.Unmarshal(data, &poster)
	if err != nil {
		log.Fatal(err)
	}

	return &poster
}

func fetchAndSavePoster(poster *SearchResult) {
	fmt.Println("Record Found.....")
	fmt.Println("Downloading Image.....")
	imgUrl := poster.Search[0].Poster
	movTitle := poster.Search[0].Title
	resp, err := http.Get(imgUrl)
	if err != nil {
		log.Fatal("An error occurred when fetching image: ", err)
	}
	defer resp.Body.Close()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Failed to get user's home directory: ", err)
	}

	downloadsDir := filepath.Join(homeDir, "Downloads")
	filePath := filepath.Join(downloadsDir, movTitle+".jpg")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("An error occurred during file creation: ", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println("Successfully saved in ", filePath)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a movie title")
		return
	}

	title := parseMovieTitle(os.Args[1:])
	poster := fetchMovieData(title)

	if len(poster.Search) == 0 {
		log.Fatal("No records found")
		return
	}
	fetchAndSavePoster(poster)
}
