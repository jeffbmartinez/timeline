package handler

import (
	"net/url"
	"testing"
)

func TestGetMissingArgs_NoRequiredArgs(t *testing.T) {
	urlArgs, _ := url.ParseQuery("key1=value1&key2=value2")

	NO_REQUIRED_ARGS := []string{}

	if missingArgs := GetMissingArgs(urlArgs, NO_REQUIRED_ARGS); len(missingArgs) != 0 {
		t.Fatal("Should not have returned any missing arguments")
	}
}

func TestGetMissingArgs_OneRequiredArg(t *testing.T) {
	urlArgs, _ := url.ParseQuery("key1=value1&key2=value2")

	NO_REQUIRED_ARGS := []string{"key1"}

	if missingArgs := GetMissingArgs(urlArgs, NO_REQUIRED_ARGS); len(missingArgs) != 0 {
		t.Fatal("Should not have returned any missing arguments")
	}
}

func TestGetMissingArgs_MissingRequiredArgs(t *testing.T) {
	urlArgs, _ := url.ParseQuery("key1=value1&key2=value2")

	NO_REQUIRED_ARGS := []string{"missingArg1", "missingArg2"}

	missingArgs := GetMissingArgs(urlArgs, NO_REQUIRED_ARGS)
	if len(missingArgs) != 2 || missingArgs[0] != "missingArg1" || missingArgs[1] != "missingArg2" {
		t.Fatal("Did not detect the missing arguments")
	}
}
