package handler

import (
	"net/http"
)

func methodNotAllowed(response http.ResponseWriter, request *http.Request) {
	WriteSimpleResponse(response, "method not allowed", http.StatusMethodNotAllowed)
}
