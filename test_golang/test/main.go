package main

import (
	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	// "math/rand"
	"net/http"
	// "strconv"
	"test/movie"
)


// var movies []Movie
var movies  = movie.Movies


func main() {
	r := mux.NewRouter()
	movie.Initial()
	r.HandleFunc("/movies", movie.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", movie.GetMovie).Methods("GET")
	r.HandleFunc("/movies", movie.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", movie.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", movie.DeleteMovie).Methods("DELETE")

	fmt.Println("Starting at port 127.0.0.1:8000")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))

}
