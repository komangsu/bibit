package repository

import (
	"fmt"
)

var (
	API_URL = "http://www.omdbapi.com/?apikey="
	API_KEY = "faf7e5bb"
)

func GetMovieDetail(id string) string {
	BASE_URL := API_URL + API_KEY
	return fmt.Sprintf("%s&i=%s", BASE_URL, id)
}

func SearchMovie(s string, page int) string {
	BASE_URL := API_URL + API_KEY
	return fmt.Sprintf("%s&s=%s&page=%d", BASE_URL, s, page)
}
