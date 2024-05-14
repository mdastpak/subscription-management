package main

import (
	"log"
	"net/http"
	"subscription-management/routes"
)

func main() {
	router := routes.InitRoutes()
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
