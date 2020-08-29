package server

import (
	"fmt"
	. "github.com/GoosvandenBekerom/skateboarding-api/pkg/api/controllers"
	"log"
	"net/http"
)

func StartWebServer() {
	port := 8080
	log.Printf("Webserver running on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), RequestHandler{
		pathHandlers: map[string]func(request *http.Request, response http.ResponseWriter){
			"/":       RootController{}.HandleRoot,
			"/tricks": TrickController{}.HandleRoot,
			//"/tricks/{name}": func(request *http.Request, response http.ResponseWriter) {
			//	parts := strings.Split(request.URL.Path, "/")
			//	TrickController{}.HandleGetByName(parts[1], request, response)
			//},
		},
	}))
}
