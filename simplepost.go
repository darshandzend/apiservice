package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/jsonpost", postHandler)

	http.ListenAndServe(":8080", nil)
}
