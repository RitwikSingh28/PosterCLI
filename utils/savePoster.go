package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func FetchAndSavePoster(poster *SearchResult) error {
	fmt.Println("Record Found.....")
	fmt.Println("Downloading Image.....")
	imgUrl := poster.Search[0].Poster
	movTitle := poster.Search[0].Title
	resp, err := http.Get(imgUrl)
	if err != nil {
		return fmt.Errorf("An error occurred when fetching image: %v\n", err)
	}
	defer resp.Body.Close()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Failed to get user's home directory: %v\n", err)
	}

	downloadsDir := filepath.Join(homeDir, "Downloads")
	filePath := filepath.Join(downloadsDir, movTitle+".jpg")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("An error occurred during file creation: %v\n", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("Couldn't create the file: %v", err)
	}
	fmt.Println("Successfully saved in ", filePath)
	return nil
}
