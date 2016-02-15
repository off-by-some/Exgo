package routes

import (
	sessionResource "github.com/Pholey/Exgo/resources/session"
	"github.com/gin-gonic/gin"
)

// Route - Struct containing all the info to initialize a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}

// Routes - Array of routes
var Routes = []Route{
	Route{
		"CreateUser",
		"POST",
		"/user",
		sessionResource.Create,
	},
}
