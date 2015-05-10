package config

import (
	"testing"
)

func TestReadKeyExists(t *testing.T) {
	config, err := ReadGeneral("testconfig/simple.json")

	if err != nil {
		t.Fatal(err)
	}

	if config["key"] != "value" {
		t.Fatal("Unexpected config value")
	}
}

func TestReadKeyMissing(t *testing.T) {
	config, err := ReadGeneral("testconfig/simple.json")

	if err != nil {
		t.Fatal(err)
	}

	if config["does not exist"] != nil {
		t.Fatal("Unexpected config value")
	}
}

func TestFileNotFound(t *testing.T) {
	_, err := ReadGeneral("non existing filename")

	if err == nil {
		t.Fatal("err should not have been nil because file does not exist")
	}
}

func TestEmptyConfigFile(t *testing.T) {
	config, err := ReadGeneral("testconfig/empty.json")

	if err != nil {
		t.Fatal("Returned nil, but shouldn't have")
	}

	if len(config) != 0 {
		t.Fatal("Expected empty map")
	}
}

func TestReadSpecificSimple(t *testing.T) {
	var simpleConfig SimpleConfig
	err := ReadSpecific("testconfig/simple.json", &simpleConfig)

	if err != nil {
		t.Fatalf("Config read failed with error: %v", err)
	}

	if simpleConfig.Key != "value" {
		t.Fatal("Config not read properly")
	}
}

func TestReadSpecificComplex(t *testing.T) {
	var person Person
	err := ReadSpecific("testconfig/persons.json", &person)

	if err != nil {
		t.Fatalf("Config read failed with error: %v", err)
	}

	if person.Name != "jeff" || person.Age != 30 || len(person.Friends) != 2 {
		t.Fatal("Config read incorrect values")
	}
}

type SimpleConfig struct {
	Key string
}

type Person struct {
	Name    string
	Age     uint
	Friends []Person
}
