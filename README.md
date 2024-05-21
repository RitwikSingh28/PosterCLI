# Movie Poster Fetcher

This is a simple Go application that fetches movie data from the OMDB API and downloads the poster of the first search result to your `Downloads` directory.

## Features

- Fetches movie data based on a provided title
- Downloads the poster of the first search result
- Saves the poster in the user's `Downloads` directory

## Installation

1. Ensure you have [Go](https://golang.org/dl/) installed on your machine.
2. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/movie-poster-fetcher.git
    cd movie-poster-fetcher
    ```

## Usage

To run the application, use the following command:

    ```sh
    go run main.go "movie title"
    ```
