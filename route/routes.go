package route

import (
	"net/http"

	handlers "github.com/djung460/cypress/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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
