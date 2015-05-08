package handler

import (
	"fmt"
	"net/http"

	"github.com/jeffbmartinez/log"

	// "github.com/jeffbmartinez/timeline/storage"
)

func Simple(response http.ResponseWriter, request *http.Request) {
	urlArgs := request.URL.Query()

	log.Info(urlArgs)

	REQUIRED_ARGS := []string{
		"owner",
		"category",
	}

	missingArgs := make([]string, 0, len(REQUIRED_ARGS))

	for _, expectedArg := range REQUIRED_ARGS {
		if argument := urlArgs.Get(expectedArg); argument == "" {
			missingArgs = append(missingArgs, expectedArg)
		}
	}

	if len(missingArgs) > 0 {
		errorMessage := fmt.Sprintf("Missing arguments: %v", missingArgs)
		log.Infof(errorMessage)
		WriteSimpleResponse(response, errorMessage, http.StatusBadRequest)
		return
	}

	// event := &storage.Event{
	// 	Owner:    urlArgs.Get("owner"),
	// 	Category: urlArgs.Get("category"),
	// }

	WriteSimpleResponse(response, "simple event recorded", http.StatusOK)
}
