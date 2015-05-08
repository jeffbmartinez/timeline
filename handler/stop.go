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

	// event := &storage.Event{
	// 	Owner:    urlArgs.Get("owner"),
	// 	Category: urlArgs.Get("category"),
	// }

	WriteSimpleResponse(response, "stop event recorded", http.StatusOK)
}
