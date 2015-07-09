package handler

import (
	"net/http"
)

func notImplemented(response http.ResponseWriter, request *http.Request) {
	WriteSimpleResponse(response, "not implemented", http.StatusNotImplemented)
}
