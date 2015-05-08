package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/influxdb/influxdb/client"
	"github.com/jeffbmartinez/log"

	"github.com/jeffbmartinez/timeline/storage/influxdb"
)

func Simple(response http.ResponseWriter, request *http.Request) {
	urlArgs := request.URL.Query()
	log.Infof("Simple event request with args: %v", urlArgs)

	REQUIRED_ARGS := []string{
		"series",
	}

	missingArgs := GetAnyMissingArgs(urlArgs, REQUIRED_ARGS)

	if len(missingArgs) > 0 {
		errorMessage := fmt.Sprintf("Missing required arguments: %v", missingArgs)
		log.Infof(errorMessage)
		WriteSimpleResponse(response, errorMessage, http.StatusBadRequest)
		return
	}

	influxDbClient, err := influxdb.GetClient()
	if err != nil {
		fmt.Errorf("Could not get an influxdb client: %v", err)
		WriteSimpleResponse(response, "Unable to store event", http.StatusInternalServerError)
		return
	}

	pts := make([]client.Point, 1)
	pts[0] = client.Point{
		Name: "test",
		Fields: map[string]interface{}{
			"name": "jeff",
			"age":  "50",
		},
		Timestamp: time.Now(),
		Precision: "s",
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        "test_timeline",
		RetentionPolicy: "default",
	}

	// event := newSimpleEvent(urlArgs)

	_, err = influxDbClient.Write(bps)

	if err != nil {
		log.Errorf("Had an error writing to influxdb: %v", err)
		WriteSimpleResponse(response, "Unable to store event", http.StatusInternalServerError)
		return
	}

	WriteSimpleResponse(response, "simple event recorded", http.StatusOK)
}

func newSimpleEvent(urlArgs url.Values) client.Point {
	seriesName := urlArgs.Get("series")

	fields := make(map[string]interface{}, len(urlArgs))

	for key := range urlArgs {
		if key != "series" {
			fields[key] = urlArgs.Get(key)
		}
	}

	log.Info(fields)

	return client.Point{
		Name:   seriesName,
		Fields: fields,
	}
}
