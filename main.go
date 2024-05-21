// Todo: make it modular, and robust
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	title := os.Args[1:]
	movieTitle := ""
	for _, partTitle := range title {
		movieTitle += partTitle + " "
	}
	movieTitle = movieTitle[:len(movieTitle)-1]
	encodedTitle := url.QueryEscape(movieTitle)
	url := "http://www.omdbapi.com/?s=" + encodedTitle + "&apikey=1b8a3f2d"

	fmt.Println("Fetching Data....")
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

	fmt.Fprint(os.Stdin, "")
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

	if len(poster.Search) == 0 {
		log.Fatal("No records found")
		return
	}

	fmt.Println("Record Found.....")
	fmt.Println("Downloading Image.....")
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
