package router

import (
	"net/http"
	"syndio/handlers"
	"syndio/services"

	"github.com/gorilla/mux"
)

func InitRouter(dependencies services.ServerDependencies) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/employments", handlers.CreateEmployments(dependencies.SyndioService)).Methods(http.MethodPost)
	return router
}
