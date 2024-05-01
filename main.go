package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thjis is the about page!")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the homepage!")
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)

	http.ListenAndServe(port, nil)
}
