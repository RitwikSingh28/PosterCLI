package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	title := os.Args[1:]
	url := "http://www.omdbapi.com/?s=" + title[0] + "&apikey=1b8a3f2d"

	fmt.Printf("%s\n", url)
	resp, err := http.Get(url)
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

	type SearchResult struct {
		Search []struct {
			Title  string `json:"Title"`
			Year   string `json:"Year"`
			Poster string `json:"Poster"`
		} `json:"Search"`
	}
	var poster SearchResult

	err = json.Unmarshal(data, &poster)
	if err != nil {
		log.Fatal(err)
	}

	imgUrl := poster.Search[0].Poster
	movTitle := poster.Search[0].Title
	resp, err = http.Get(imgUrl)
	if err != nil {
		log.Fatal("An error occurred when fetching image: ", err)
	}
	defer resp.Body.Close()

	filePath := "/home/ritro015/Downloads/" + movTitle + ".jpg"
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
