package routes

import (
	"github.com/bispoman/go-APIs/controller"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter().StrictSlash(false)

	routes.HandleFunc("/Hello", controller.Hello).Methods("GET")

	return routes
}
