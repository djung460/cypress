package route

import (
	"github.com/gorilla/mux"

	"github.com/djung460/cypress/handlers"
	"github.com/djung460/cypress/models"
)

func NewRouter(db models.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/", handlers.Index()).Methods("GET")
	router.Handle("/api/nuggets", handlers.UserNuggetIndex(db)).Methods("GET")
	router.Handle("/api/nugget/create", handlers.NuggetCreate(db)).Methods("POST")
	router.Handle("/api/categories", handlers.CategoryIndex(db)).Methods("GET")
	router.Handle("/api/category/create", handlers.CategoryCreate(db)).Methods("POST")

	return router
}
