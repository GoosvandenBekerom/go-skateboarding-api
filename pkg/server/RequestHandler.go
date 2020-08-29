package server

import (
	"fmt"
	"log"
	"net/http"
)

type RequestHandler struct {
	pathHandlers map[string]func(request *http.Request, response http.ResponseWriter)
}

func (handler RequestHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	path := request.URL.Path

	if handleRequest, ok := handler.pathHandlers[path]; ok {
		handleRequest(request, response)
	} else {
		response.WriteHeader(http.StatusNotFound)
		_, err := fmt.Fprintf(response, "The resource you were looking at %s for could not be found.", path)

		if err != nil {
			log.Println(err)
		}
	}
}
