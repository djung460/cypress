package main

import (
	"log"
	"net/http"

	routes "github.com/djung460/cypress/routes"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
