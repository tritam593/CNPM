package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test/book_store/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.Register_book_store_routes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("127.0.0.1:9010", r))
}
