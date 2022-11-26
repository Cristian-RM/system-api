package main

import (
	"log"
	"net/http"

	"github.com/Cristian-RM/system-api/pkg/models"
	"github.com/Cristian-RM/system-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	models.Init()
	r := mux.NewRouter()
	routes.RegisterEmantaRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
