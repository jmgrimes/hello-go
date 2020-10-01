package handlers

import(
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmgrimes/hello-go/configs"
)

type HelloResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

func GetHelloEndpoint(writer http.ResponseWriter, request *http.Request) {
	var response HelloResponse
	var name = configs.Configuration.DefaultName
	names, ok := request.URL.Query()["name"]
	if ok && len(names[0]) > 0 {
		 name = names[0]
	}
	response.Message = fmt.Sprintf(configs.Configuration.Template, name)
	response.Version = configs.Configuration.Version
	json.NewEncoder(writer).Encode(response)
}