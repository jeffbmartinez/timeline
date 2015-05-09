package handler

import (
	"fmt"
	"net/http"

	"github.com/jeffbmartinez/log"
)

func Stop(response http.ResponseWriter, request *http.Request) {
	urlArgs := request.URL.Query()

	log.Info(urlArgs)

	REQUIRED_ARGS := []string{
		"startRequestId",
	}

	missingArgs := GetAnyMissingArgs(urlArgs, REQUIRED_ARGS)

	if len(missingArgs) > 0 {
		errorMessage := fmt.Sprintf("Missing required arguments: %v", missingArgs)
		log.Infof(errorMessage)
		WriteSimpleResponse(response, errorMessage, http.StatusBadRequest)
		return
	}

	// todo

	WriteSimpleResponse(response, "stop point recorded", http.StatusOK)
}
