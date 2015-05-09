package handler

import (
	"fmt"
	"net/http"

	"github.com/jeffbmartinez/log"
)

func Start(response http.ResponseWriter, request *http.Request) {
	urlArgs := request.URL.Query()

	log.Info(urlArgs)

	REQUIRED_ARGS := []string{
		"series",
	}

	missingArgs := GetAnyMissingArgs(urlArgs, REQUIRED_ARGS)

	if len(missingArgs) > 0 {
		errorMessage := fmt.Sprintf("Missing required arguments: %v", missingArgs)
		log.Infof(errorMessage)
		WriteSimpleResponse(response, errorMessage, http.StatusBadRequest)
		return
	}

	// todo

	WriteSimpleResponse(response, "start point recorded", http.StatusOK)
}
