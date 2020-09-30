package main

import (
	"log"
	"net/http"

	"github.com/jmgrimes/hello-go/routes"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}