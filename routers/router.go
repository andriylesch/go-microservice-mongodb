package routers

import (
	"go-microservice-mongodb/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes []Route

func init() {

	routes = append(routes, Route{
		Name:        "UsersGet",
		Method:      "GET",
		Pattern:     "/users",
		HandlerFunc: handlers.GetUsers,
	})

	routes = append(routes, Route{
		Name:        "UserGetById",
		Method:      "GET",
		Pattern:     "/users/{userId}",
		HandlerFunc: handlers.GetUserById,
	})

	routes = append(routes, Route{
		Name:        "UserCreate",
		Method:      "POST",
		Pattern:     "/users",
		HandlerFunc: handlers.PostUser,
	})

	routes = append(routes, Route{
		Name:        "UserUpdate",
		Method:      "PUT",
		Pattern:     "/users/{userId}",
		HandlerFunc: handlers.PutUser,
	})

	routes = append(routes, Route{
		Name:        "UserDelete",
		Method:      "DELETE",
		Pattern:     "/users/{userId}",
		HandlerFunc: handlers.DeleteUser,
	})

}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
