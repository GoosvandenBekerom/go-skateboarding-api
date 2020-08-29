package controllers

import (
	"fmt"
	"log"
	"net/http"
)

type RootController struct{}

func (c RootController) HandleRoot(request *http.Request, response http.ResponseWriter) {
	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, err := fmt.Fprint(response, "Welcome to my go skateboarding api")

	if err != nil {
		log.Println(err)
	}
}
