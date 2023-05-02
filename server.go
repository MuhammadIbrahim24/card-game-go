package main

import (
	"log"
	"net/http"

	"example.com/card-game/routes"

	"github.com/gorilla/mux"
)

func main() {
	//Initialize router
	router := mux.NewRouter()

	//Register Routes
	routes.RegisterRoutes(router)

	//Start server
	log.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
