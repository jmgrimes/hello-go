package handlers

import(
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

func GetHelloEndpoint(writer http.ResponseWriter, request *http.Request) {
	var response HelloResponse
	var name = "Stranger"
	names, ok := request.URL.Query()["name"]
	if ok && len(names[0]) > 0 {
		 name = names[0]
	}
	response.Message = fmt.Sprintf("Hello, %s!", name)
	response.Version = "1.0.0"
	json.NewEncoder(writer).Encode(response)
}