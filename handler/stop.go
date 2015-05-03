package handler

import (
	"net/http"
)

func Stop(response http.ResponseWriter, request *http.Request) {
	// store the data point

	WriteSimpleResponse(response, "stop event recorded", http.StatusOK)
}
