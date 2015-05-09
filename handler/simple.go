package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/influxdb/influxdb/client"
	"github.com/jeffbmartinez/log"

	"github.com/jeffbmartinez/timeline/storage/influxdb"
)

func Simple(response http.ResponseWriter, request *http.Request) {
	urlArgs := request.URL.Query()
	log.Infof("Simple point request with args: %v", urlArgs)

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

	point := newSimplePoint(urlArgs)

	err := influxdb.StorePoint(point, "test_timeline", "default")
	if err != nil {
		log.Errorf("Problem writing to influxdb: %v", err)
		WriteSimpleResponse(response, "Unable to store point", http.StatusInternalServerError)
		return
	}

	WriteSimpleResponse(response, "simple point recorded", http.StatusOK)
}

func newSimplePoint(urlArgs url.Values) client.Point {
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
