package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/GoosvandenBekerom/skateboarding-api/pkg/database"
	"log"
	"net/http"
)

type TrickController struct{}

func (c TrickController) HandleRoot(request *http.Request, response http.ResponseWriter) {
	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if name, ok := request.URL.Query()["name"]; ok {
		writeGetByNameResponse(name[0], response)
	} else {
		writeGetAllResponse(response)
	}
}

func writeGetAllResponse(w http.ResponseWriter) {
	tricks := database.GetDatabase().GetAllTricks()

	_, err2 := fmt.Fprint(w, toJson(tricks))

	if err2 != nil {
		log.Println(err2)
	}
}

func writeGetByNameResponse(name string, w http.ResponseWriter) {
	trick, notFound := database.GetDatabase().GetTrickByName(name)

	if notFound != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(notFound.Error()))
	} else {
		_, err := fmt.Fprint(w, toJson(trick))

		if err != nil {
			log.Println(err)
		}
	}
}

func toJson(v ...interface{}) string {
	encoded, err := json.Marshal(v)

	if err != nil {
		log.Println(err)
	}

	return string(encoded)
}
