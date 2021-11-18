package models

type DetailMovieResponse struct {
	Title    string `json:"Title"`
	Poster   string `json:"Poster"`
	Type     string `json:"Type"`
	Movie    string `json:"Movie"`
	Year     string `json:"Year"`
	ImdbID   string `json:"imdbID"`
	Plot     string `json:"Plot"`
	Rated    string `json:"Rated"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Language string `json:"Language"`
	Released string `json:"Released"`
	Actors   string `json:"Actors"`
	Writer   string `json:"Writer"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
}

type SearchMovieResponse struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
