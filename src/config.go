package main

import (
	"encoding/json"
	"io/ioutil"
)

var config interface{}

func FromFile() map[string]interface{} {

	if config != nil {
		return config.(map[string]interface{})
	}

	configJson, err := ioutil.ReadFile("private/config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(configJson, &config)
	return config.(map[string]interface{})
}

func ClientId() string {

	config := FromFile()
	return config["client_id"].(string)
}

func ClientSecret() string {
	config := FromFile()
	return config["client_secret"].(string)
}
