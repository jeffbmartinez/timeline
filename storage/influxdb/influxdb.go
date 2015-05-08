package influxdb

import (
	"errors"
	"fmt"
	"net/url"
	"sync"

	"github.com/influxdb/influxdb/client"
	"github.com/jeffbmartinez/config"
	"github.com/jeffbmartinez/log"
)

const (
	CONFIG_FILENAME = "config/influxdb.json"
)

type Config struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
}

var configFileContents = new(Config)

func ReadConfigFile() (*Config, error) {
	var err error

	var once sync.Once
	once.Do(func() {
		err = config.ReadSpecific(CONFIG_FILENAME, &configFileContents)

		if err != nil {
			errorMessage := fmt.Sprintf("Could not reading config file: %v", err)
			log.Error(errorMessage)
			err = errors.New(errorMessage)
			return
		}

		configProblems := verifyConfig(configFileContents)

		if len(configProblems) != 0 {
			errorMessage := fmt.Sprintf("Config has the following problems: %v", configProblems)
			log.Error(errorMessage)
			err = errors.New(errorMessage)
			return
		}

		err = nil
	})

	return configFileContents, err
}

func GetClient() (*client.Client, error) {
	serverConfig, err := ReadConfigFile()

	if err != nil {
		errorMessage := fmt.Sprintf("Could not read influxdb config file: %v", err)
		log.Error(errorMessage)
		return nil, errors.New(errorMessage)
	}

	influxDbUrl, err := url.Parse(fmt.Sprintf("http://%s:%s", serverConfig.Host, serverConfig.Port))
	if err != nil {
		errorMessage := fmt.Sprintf("Trouble building url from config file's host and port: %v", err)
		log.Error(errorMessage)
		return nil, errors.New(errorMessage)
	}

	clientConfig := client.Config{
		URL:      *influxDbUrl,
		Username: serverConfig.Username,
		Password: serverConfig.Password,
	}

	connection, err := client.NewClient(clientConfig)
	if err != nil {
		errorMessage := fmt.Sprintf("Count not create influxDB client (%v)", err)
		log.Error(errorMessage)
		return connection, errors.New(errorMessage)
	}

	return connection, nil
}

func TestConnection(connection *client.Client) error {
	_, _, err := connection.Ping()
	if err != nil {
		errorMessage := fmt.Sprintf("Could not ping the influxDB server (%v)", err)
		log.Error(errorMessage)
		return errors.New(errorMessage)
	}

	log.Info("Connection to InfluxDB successful")
	return nil
}

func StoreEvent(event interface{}) error {

	return nil
}

func verifyConfig(configuration *Config) []string {
	const NUMBER_OF_CHECKS = 10

	problems := make([]string, 0, NUMBER_OF_CHECKS)

	if configuration.Host == "" {
		problems = append(problems, "Config is missing host")
	}

	if configuration.Port == "" {
		problems = append(problems, "Config is missing port")
	}

	if configuration.DbName == "" {
		problems = append(problems, "Config is missing dbname")
	}

	if configuration.Username == "" {
		problems = append(problems, "Config is missing username")
	}

	if configuration.Password == "" {
		problems = append(problems, "Config is missing password")
	}

	return problems
}
