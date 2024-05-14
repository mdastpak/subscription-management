package routes

import (
	"subscription-management/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/service", handlers.CreateServiceHandler).Methods("POST")
	adminRouter.HandleFunc("/service", handlers.UpdateServiceHandler).Methods("PUT")
	return router
}
