package movie

import (
	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	// "math/rand"
	"net/http"

)





func main() {
	r := mux.NewRouter()
	Initial()
	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Println("Starting at port 127.0.0.1:8000")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))

}
