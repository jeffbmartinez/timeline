package handler

import (
	"net/http"
)

func Start(response http.ResponseWriter, request *http.Request) {
	// store the data point

	WriteSimpleResponse(response, "start event recorded", http.StatusOK)
}
