package influxdb

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/influxdb/influxdb/client"

	"github.com/jeffbmartinez/log"

	// "github.com/jeffbmartinez/timeline/storage"
)

func Initialize(host string, port int, dbname string, username string, password string) error {
	influxDbUrl, err := url.Parse(fmt.Sprintf("http://%s:%d", host, port))
	if err != nil {
		log.Error(err)
	}

	conf := client.Config{
		URL:      *influxDbUrl,
		Username: username,
		Password: password,
	}

	connection, err := client.NewClient(conf)
	if err != nil {
		errorMessage := fmt.Sprintf("Count not create influxDB client (%v)", err)
		log.Error(errorMessage)
		return errors.New(errorMessage)
	}

	_, _, err = connection.Ping()
	if err != nil {
		errorMessage := fmt.Sprintf("Count not ping the influxDB server (%v)", err)
		log.Error(errorMessage)
		return errors.New(errorMessage)
	}

	log.Info("Connection to InfluxDB successful")
	return nil
}
