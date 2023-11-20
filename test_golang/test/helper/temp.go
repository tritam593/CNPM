package helper

import (
	"fmt"
	"log"
	"net/http"
)

func run() {
	Print_string("hello")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", FormHandler)
	http.HandleFunc("/hello", HelloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
