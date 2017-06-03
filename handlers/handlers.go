package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/djung460/cypress/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func NuggetIndex(w http.ResponseWriter, r *http.Request) {
	nuggets := models.Nuggets{
		models.Nugget{Username: "djung", Title: "Sample Nugget", Category: "REST", Body: "Sample Rest"},
		models.Nugget{Username: "djung", Title: "Sample Nugget 2", Category: "REST", Body: "Sample Rest"},
	}

	setHeader(w)

	if err := json.NewEncoder(w).Encode(nuggets); err != nil {
		panic(err)
	}
}

func NuggetShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nuggetid := vars["nuggetid"]
	fmt.Fprintln(w, "Nugget show:", nuggetid)
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}
