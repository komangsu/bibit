package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"restapi/models"
	"restapi/repository"
	"strings"
)

type ErrResponse struct {
	Message string `json:"message"`
}

func GetMovieDetail(w http.ResponseWriter, r *http.Request) {

	// get param
	imdbID := strings.TrimPrefix(r.URL.Path, "/detail-movie/")
	errResponse := ErrResponse{}
	if imdbID == "" {
		errResponse.Message = "Missing url parameter id."
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	url := repository.GetMovieDetail(imdbID)
	fmt.Println("debug:", url)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := models.DetailMovieResponse{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)

}
