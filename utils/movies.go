package utils

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	Poster string `json:"Poster"`
}

type SearchResult struct {
	Search []Movie `json:"Search"`
}
