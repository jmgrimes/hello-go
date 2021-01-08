package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jmgrimes/hello-go/configs"
	"github.com/jmgrimes/hello-go/routes"
	_ "github.com/jmgrimes/hello-go/telemetry"
)



func main() {
	configurationFile := flag.String("config", "", "a yaml formatted configuration file")
	flag.Parse()
	if len(*configurationFile) > 0 {
		err := configs.Configure(*configurationFile)
		if (err != nil) {
			log.Fatal(err)
			return
		}
	}
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}