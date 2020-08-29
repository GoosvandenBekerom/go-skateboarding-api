package controllers

import "net/http"

type Controller interface {
	Handle(request *http.Request, response http.ResponseWriter)
}
