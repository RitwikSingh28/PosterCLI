package utils

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	Poster string `json:"Poster"`
}

type SearchResult struct {
	Search []Movie `json:"Search"`
}

const apiKey = "1b8a3f2d"
