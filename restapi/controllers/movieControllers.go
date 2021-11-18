package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"restapi/models"
	"restapi/repository"
	"strconv"
	"strings"
)

type ErrResponse struct {
	Response string `json:"Response"`
	Error    string `json:"Error"`
}

func GetMovieDetail(w http.ResponseWriter, r *http.Request) {

	// get param
	imdbID := strings.TrimPrefix(r.URL.Path, "/detail-movie/")
	errResponse := ErrResponse{}

	// handle validation parameter
	switch {
	case imdbID == "":
		errResponse.Response = "False"
		errResponse.Error = "Incorrect IMDb ID."
		json.NewEncoder(w).Encode(errResponse)
		return
	case len(imdbID) != 9:
		errResponse.Response = "False"
		errResponse.Error = "Incorrect IMDb ID."
		json.NewEncoder(w).Encode(errResponse)
		return
	case string(imdbID[0:2]) != "tt":
		errResponse.Response = "False"
		errResponse.Error = "Incorrect IMDb ID."
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	url := repository.GetMovieDetail(imdbID)

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

	// return to json
	result := models.DetailMovieResponse{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func SearchMovie(w http.ResponseWriter, r *http.Request) {
	// get query param
	paramSearch := r.URL.Query().Get("s")
	page := r.URL.Query().Get("page")

	paramPage, _ := strconv.Atoi(page)

	url := repository.SearchMovie(paramSearch, paramPage)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to server.")
		return
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// return to json
	result := models.SearchMoviesResponse{}
	err = json.Unmarshal(responseBody, &result)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
