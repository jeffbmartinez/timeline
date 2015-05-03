package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jeffbmartinez/log"
)

type SimpleResponse struct {
	Message string `json:"message"`
}

func WriteSimpleResponse(response http.ResponseWriter, message string, statusCode int) {
	msg := &SimpleResponse{Message: message}
	WriteResponse(response, msg, statusCode)
}

func WriteResponse(response http.ResponseWriter, message interface{}, statusCode int) {
	responseString, err := json.Marshal(message)

	if err != nil {
		log.Errorf("Couldn't marshal json: %v", err)

		response.WriteHeader(statusCode)
		response.Write([]byte(""))
		return
	}

	response.WriteHeader(statusCode)
	response.Write([]byte(responseString))
}
