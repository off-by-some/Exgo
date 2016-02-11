package resources

import (
	"net/http"

	L "github.com/Pholey/Exgo/logger"
	sessionResource "github.com/Pholey/Exgo/resources/session"

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

		// Set up logging for each request
		handler := L.Logger(route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"CreateUser",
		"POST",
		"/user",
		sessionResource.Create,
	},
	Route{
		"Auth",
		"POST",
		"/auth",
		sessionResource.Auth,
	},
}
