package logging

import (
	"log"
	"net/http"
	"time"
)

func LogHandler(handler http.Handler, route string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		handler.ServeHTTP(writer, request)
		log.Printf(
			"%s\t%s\t%s\t%s",
			request.Method,
			request.RequestURI,
			route,
			time.Since(startTime),
		)
	})
}