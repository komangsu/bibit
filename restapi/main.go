package main

import (
	"log"
	"net/http"
	"restapi/controllers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/detail-movie/", controllers.GetMovieDetail)
	mux.HandleFunc("/search-movie", controllers.SearchMovie)

	// if not avaliable route then return 404
	mux.Handle("/resources", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":3002", mux))
}
