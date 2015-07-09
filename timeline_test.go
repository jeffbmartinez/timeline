package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jeffbmartinez/timeline/handler"
)

func TestEventEndpoint_PostWithNoParams(t *testing.T) {
	request, _ := http.NewRequest("POST", "http://example.com", nil)
	response := httptest.NewRecorder()

	handler.Event(response, request)

	const EXPECTED_STATUS_CODE = http.StatusBadRequest

	if response.Code != EXPECTED_STATUS_CODE {
		t.Fatalf("Should have received status code %v", EXPECTED_STATUS_CODE)
	}
}

func TestEventEndpoint_PostWithMissingNameParam(t *testing.T) {
	request, _ := http.NewRequest("POST", "http://localhost?key1=value1&key2=value", nil)
	response := httptest.NewRecorder()

	handler.Event(response, request)

	const EXPECTED_STATUS_CODE = http.StatusBadRequest

	if response.Code != EXPECTED_STATUS_CODE {
		t.Fatalf("Should have received status code %v", EXPECTED_STATUS_CODE)
	}
}

func TestMeasurementEndpoint_PostWithNoParams(t *testing.T) {
	request, _ := http.NewRequest("POST", "http://example.com", nil)
	response := httptest.NewRecorder()

	handler.Measurement(response, request)

	const EXPECTED_STATUS_CODE = http.StatusBadRequest

	if response.Code != EXPECTED_STATUS_CODE {
		t.Fatalf("Should have received status code %v", EXPECTED_STATUS_CODE)
	}
}

func TestMeasurementEndpoint_PostWithMissingNameParam(t *testing.T) {
	request, _ := http.NewRequest("POST", "http://localhost?key1=value1&key2=value", nil)
	response := httptest.NewRecorder()

	handler.Measurement(response, request)

	const EXPECTED_STATUS_CODE = http.StatusBadRequest

	if response.Code != EXPECTED_STATUS_CODE {
		t.Fatalf("Should have received status code %v", EXPECTED_STATUS_CODE)
	}
}
