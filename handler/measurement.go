package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/influxdb/influxdb/client"
	"github.com/jeffbmartinez/log"

	"github.com/jeffbmartinez/timeline/storage/influxdb"
)

func Measurement(response http.ResponseWriter, request *http.Request) {
	handler := methodNotAllowed

	switch request.Method {
	case "GET":
		handler = getMeasurement
	case "POST":
		handler = postMeasurement
	}

	handler(response, request)
}

func getMeasurement(response http.ResponseWriter, request *http.Request) {
	notImplemented(response, request)
}

func postMeasurement(response http.ResponseWriter, request *http.Request) {
	urlArgs, err := parseUrlArgs(request.URL.Query())
	if err != nil {
		errorMessage := err.Error()
		log.Infof(errorMessage)
		WriteSimpleResponse(response, errorMessage, http.StatusBadRequest)
		return
	}

	log.Infof("Measurement storage request with args: %v", urlArgs)

	REQUIRED_ARGS := []string{
		"name",
		"value",
	}

	missingArgs := GetMissingArgs(urlArgs, REQUIRED_ARGS)

	if len(missingArgs) > 0 {
		errorMessage := fmt.Sprintf("Missing required arguments: %v", missingArgs)
		log.Infof(errorMessage)
		WriteSimpleResponse(response, errorMessage, http.StatusBadRequest)
		return
	}

	tags := getTagsFromUrlArgs(urlArgs, REQUIRED_ARGS)
	measurement := newMeasurement(urlArgs["name"], urlArgs["value"], tags)

	influxDbConfig, err := influxdb.GetConfiguration()
	if err != nil {
		log.Errorf("Problem reading influxdb configuration: %v", err)
		WriteSimpleResponse(response, "Unable to store measurement", http.StatusInternalServerError)
		return
	}

	const RETENTION_POLICY string = "default"

	err = influxdb.StorePoint(measurement, influxDbConfig.DbName, RETENTION_POLICY)
	if err != nil {
		log.Errorf("Problem writing to influxdb: %v", err)
		WriteSimpleResponse(response, "Unable to store measurement", http.StatusInternalServerError)
		return
	}

	WriteSimpleResponse(response, "measurement recorded", http.StatusOK)
}

func newMeasurement(name string, value interface{}, tags map[string]string) client.Point {
	return client.Point{
		Measurement: name,
		Tags:        tags,
		Time:        time.Now(),
		Fields: map[string]interface{}{
			"value": value,
		},
		Precision: "s",
	}
}
