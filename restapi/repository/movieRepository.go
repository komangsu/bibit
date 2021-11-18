package repository

import (
	"fmt"
)

func GetMovieDetail(id string) string {
	API_URL := "http://www.omdbapi.com/?apikey="
	API_KEY := "faf7e5bb"

	BASE_URL := API_URL + API_KEY
	return fmt.Sprintf("%s&i=%s", BASE_URL, id)
}
