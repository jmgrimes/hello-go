package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmgrimes/hello-go/handlers"
	"github.com/jmgrimes/hello-go/logging"
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
		var handler http.Handler = logging.LogHandler(route.HandlerFunc, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}

var routes = Routes{
	Route{
		"GetHelloEndpoint",
		"GET",
		"/hello-world",
		handlers.GetHelloEndpoint,
	},
}