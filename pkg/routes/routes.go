package routes

import (
	"github.com/Cristian-RM/system-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterEmantaRoutes = func(router *mux.Router) {
	router.HandleFunc("/authentication", controllers.Authenticate).Methods("POST")
}
