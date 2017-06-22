package main

import (
	"log"
	"net/http"
	"os"

	"github.com/djung460/cypress/models"
	route "github.com/djung460/cypress/route"
)

type Server struct {
	Hostname string `json:"hostname"`
	UseHTTP  bool   `json:"http"`
	UseHTTPS bool   `json:"https"`
	Port     int    `json:"port"`
	SSLPort  int    `json:"sslport"`
	CertFile int    `json:"CertFile"`
	KeyFile  int    `json:"KeyFile"`
}

func main() {
	var db models.DB
	dev := true
	if dev {
		d, err := models.Init()
		if err != nil {
			log.Fatalf("error: [%s]", err)
			os.Exit(1)
		}
		db = d
	}

	router := route.NewRouter(db)

	log.Fatal(http.ListenAndServe(":8080", router))
}
