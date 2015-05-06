package influxdb

import (
	"fmt"
	"net/url"

	"github.com/influxdb/influxdb/client"

	"github.com/jeffbmartinez/log"

	// "github.com/jeffbmartinez/timeline/storage"
)

func Initialize(host string, port int, dbname string, username string, password string) {
	influxDbUrl, err := url.Parse(fmt.Sprintf("http://%s:%d", host, port))
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL:      *influxDbUrl,
		Username: getInfluxDbUsername(),
		Password: getInfluxDbPassword(),
	}

	connection, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	_, _, err = connection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Connection to InfluxDB successful")
}

func getInfluxDbUsername() string {
	// username := os.Getenv("INFLUX_USER")
	return "root"
}

func getInfluxDbPassword() string {
	// password := os.Getenv("INFLUX_PWD")
	return "root"
}
