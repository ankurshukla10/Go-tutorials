package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go is neat")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is about page!")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)
}
