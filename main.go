package main

import (
	"go-microservice-mongodb/routers"
	"log"
	"net/http"
)

func main() {

	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
