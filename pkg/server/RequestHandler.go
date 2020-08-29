package server

import (
	"github.com/GoosvandenBekerom/skateboarding-api/pkg/api/controllers"
	"net/http"
)

type RequestHandler struct{}

func (handler RequestHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	var controller controllers.Controller

	switch request.URL.Path {
	case "/":
		controller = controllers.RootController{}
	default:
		controller = controllers.NotFoundController{}
	}

	controller.Handle(request, response)
}
