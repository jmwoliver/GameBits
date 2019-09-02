package main

import (
	"log"
	"net/http"
)

const BASE_URL = "https://downloads.khinsider.com"

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
