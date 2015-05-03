package handler

import (
	"net/http"
)

func Single(response http.ResponseWriter, request *http.Request) {
	// store the data point

	WriteSimpleResponse(response, "single event recorded", http.StatusOK)
}
