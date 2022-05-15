package main

import (
	"fmt"
	"get/http"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:S"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/(id)", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/(id)", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/(id)", deleteMovie).Methods("DELETE")

	fmt.Printf(("Starting Server At PORT 8080\n"))
	log.Fatal(http.ListenAndServe(":8080", r))
}
