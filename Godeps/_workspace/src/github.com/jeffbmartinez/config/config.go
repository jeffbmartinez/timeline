package config

import (
	"encoding/json"
	"io/ioutil"
)

/*
Reads a config file into a general map structure. Makes no assumptions
about underlying structure of the config file that is read in.
*/
func ReadGeneral(filename string) (config map[string]interface{}, err error) {
	fileContents, err := ioutil.ReadFile(filename)

	if err == nil {
		json.Unmarshal(fileContents, &config)
	}

	return
}

/*
Reads a config file into a specific type. Must match the structure of the
type passed in in order to succeed.
*/
func ReadSpecific(filename string, configObject interface{}) error {
	fileContents, err := ioutil.ReadFile(filename)

	if err == nil {
		json.Unmarshal(fileContents, &configObject)
	}

	return err
}
