package main

import (
	"log"
	"os"

	"github.com/RitwikSingh28/poster_cli/utils"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a movie title")
		return
	}

	title, err := utils.ParseMovieTitle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	poster, err := utils.FetchMovieData(title)
	if err != nil {
		log.Fatal(err)
	}

	if len(poster.Search) == 0 {
		log.Fatal("No records found")
		return
	}
	if err = utils.FetchAndSavePoster(poster); err != nil {
		log.Fatal(err)
	}
}
