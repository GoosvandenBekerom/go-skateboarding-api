package server

import (
	"fmt"
	"log"
	"net/http"
)

func StartWebServer() {
	port := 8080
	log.Printf("Webserver running on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), RequestHandler{}))
}
