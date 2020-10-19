package exampleAPI

import (
	"net/http"

	"github.com/adeturner/exampleAPI/utils"
	"github.com/gorilla/mux"
)

// Router defines the required methods for retrieving api routes
type Router interface {
	Routes() utils.Routes
}

// NewRouter creates a new router for any number of api routers
func NewRouter(routers ...Router) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, api := range routers {
		for _, route := range api.Routes() {
			var handler http.Handler
			handler = route.HandlerFunc
			handler = utils.Logger(handler, route.Name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}

	return router
}
