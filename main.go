package main

import (
	"log"
	"net/http"

	route "github.com/djung460/cypress/route"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
