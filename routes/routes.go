package route

import (
	"net/http"

	handlers "github.com/djung460/cypress/handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"NuggetIndex",
		"GET",
		"/nuggets",
		handlers.NuggetIndex,
	},
	Route{
		"NuggetShow",
		"GET",
		"/nuggets/{todoId}",
		handlers.NuggetShow,
	},
}
