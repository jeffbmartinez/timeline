package handler

import (
	"encoding/json"
	"net/http"
	"net/url"

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

func GetAnyMissingArgs(urlArgs url.Values, requiredArgs []string) []string {
	missingArgs := make([]string, 0, len(urlArgs))

	for _, requiredArg := range requiredArgs {
		if urlArgs.Get(requiredArg) == "" {
			missingArgs = append(missingArgs, requiredArg)
		}
	}

	return missingArgs
}
