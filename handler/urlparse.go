package handler

import (
	"errors"
	"net/url"
)

// url.Values is a defined as map[string][]string
// since I don't want multiple values per key, this function
// ensures keys are unique and flattens the map to return
// a map[string]string instead of map[string][]string
//
// The purpose is to expicitly disallow the following otherwise legal
// query params:
//  - multiple values for one key:
//      example.com?key=value1&key=value2
//  - blank value for a key
//      example.com?name=jeff&age=&city=seattle
func parseUrlArgs(urlArgs url.Values) (map[string]string, error) {
	flatUrlArgs := make(map[string]string)

	for key, value := range urlArgs {
		if len(value) != 1 {
			return nil, errors.New("Url arguments must have exactly one value per key: Duplicates and empty values are not allowed")
		}

		flatUrlArgs[key] = value[0]
	}

	return flatUrlArgs, nil
}

func getTagsFromUrlArgs(urlArgs map[string]string, nonTagKeys []string) map[string]string {
	tags := make(map[string]string)

	nonTagKeysMap := convertToMap(nonTagKeys)

	for key, value := range urlArgs {
		_, ok := nonTagKeysMap[key]
		if !ok {
			tags[key] = value
		}
	}

	return tags
}

// Creates a map[string]bool from the elements in a slice as keys, which can
// be used as a set. It allows me to quickly check if an element exists in
// the original set or not.
func convertToMap(slice []string) map[string]bool {
	setmap := make(map[string]bool)

	for _, element := range slice {
		setmap[element] = true
	}

	return setmap
}

// If you expect to check if an element exists in a the same slice
// multiple times, don't use this. It makes more sense to make
// a map with the elements of the slice and just check if the key exists.
func existsInSlice(value string, slice []string) bool {
	for _, element := range slice {
		if value == element {
			return true
		}
	}

	return false
}
