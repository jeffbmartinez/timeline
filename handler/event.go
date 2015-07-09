package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/influxdb/influxdb/client"
	"github.com/jeffbmartinez/log"

	"github.com/jeffbmartinez/timeline/storage/influxdb"
)

func Event(response http.ResponseWriter, request *http.Request) {
	handler := methodNotAllowed

	switch request.Method {
	case "GET":
		handler = getEvent
	case "POST":
		handler = postEvent
	}

	handler(response, request)
}

func getEvent(response http.ResponseWriter, request *http.Request) {
	notImplemented(response, request)
}

func postEvent(response http.ResponseWriter, request *http.Request) {
	urlArgs, err := parseUrlArgs(request.URL.Query())
	if err != nil {
		errorMessage := err.Error()
		log.Infof(errorMessage)
		WriteSimpleResponse(response, errorMessage, http.StatusBadRequest)
		return
	}

	log.Infof("Event storage request with args: %v", urlArgs)

	REQUIRED_ARGS := []string{
		"name",
	}

	missingArgs := GetMissingArgs(urlArgs, REQUIRED_ARGS)

	if len(missingArgs) > 0 {
		errorMessage := fmt.Sprintf("Missing required arguments: %v", missingArgs)
		log.Infof(errorMessage)
		WriteSimpleResponse(response, errorMessage, http.StatusBadRequest)
		return
	}

	tags := getTagsFromUrlArgs(urlArgs, REQUIRED_ARGS)
	event := newEvent(urlArgs["name"], tags)

	influxDbConfig, err := influxdb.GetConfiguration()
	if err != nil {
		log.Errorf("Problem reading influxdb configuration: %v", err)
		WriteSimpleResponse(response, "Unable to store event", http.StatusInternalServerError)
		return
	}

	const RETENTION_POLICY string = "default"

	err = influxdb.StorePoint(event, influxDbConfig.DbName, RETENTION_POLICY)
	if err != nil {
		log.Errorf("Problem writing to influxdb: %v", err)
		WriteSimpleResponse(response, "Unable to store event", http.StatusInternalServerError)
		return
	}

	WriteSimpleResponse(response, "event recorded", http.StatusOK)
}

// Influxdb doesn't support the storage of a measurement without fields
// attached to it. In this case, I'm just using a field of value=1
// in order to store an event. The 1 is arbitrary, except that perhaps
// it could be used to sum up the number of events that occurred within
// a timeframe.
func newEvent(name string, tags map[string]string) client.Point {
	return client.Point{
		Measurement: name,
		Tags:        tags,
		Time:        time.Now(),
		Fields: map[string]interface{}{
			"value": 1,
		},
		Precision: "s",
	}
}
