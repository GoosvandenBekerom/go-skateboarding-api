package controllers

import (
	"fmt"
	"net/http"
)

type NotFoundController struct{}

func (c NotFoundController) Handle(request *http.Request, response http.ResponseWriter) {
	response.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprintf(response, "The resource you were looking at %s for could not be found.", request.URL.Path)

	if err != nil {
		fmt.Println(err)
	}
}
