package apiroutes

import (
	"api/handler"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	Route{
		"GetAllUsers",
		"GET",
		"/user",
		handler.GetAllUsers,
	},
	Route{
		"GetByID",
		"GET",
		"/user/{Id}",
		handler.GetByID,
	},
	Route{
		"CreateUser",
		"POST",
		"/user",
		handler.CreateUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/user/{Id}",
		handler.UpdateUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/user/{Id}",
		handler.DeleteUser,
	},
}

//NewRouter():
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		//handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
