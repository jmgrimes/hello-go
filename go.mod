module github.com/jmgrimes/hello-go

go 1.15

require github.com/gorilla/mux v1.8.0

replace github.com/jmgrimes/hello-go/handlers => ./handlers
replace github.com/jmgrimes/hello-go/logging => ./logging
replace github.com/jmgrimes/hello-go/routes => ./routes