package config

import (
	"encoding/json"
	"io/ioutil"
)

var config interface{}

func NewConfigFromFile() map[string]interface{} {

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

	config := NewConfigFromFile()
	return config["client_id"].(string)
}

func ClientSecret() string {
	config := NewConfigFromFile()
	return config["client_secret"].(string)
}

func Port() string {
	config := NewConfigFromFile()
	return config["server_port"].(string)
}

func DbUsername() string {
	config := NewConfigFromFile()
	return config["db_username"].(string)
}

func DbPassword() string {
	config := NewConfigFromFile()
	return config["db_password"].(string)
}

func SessionAuthenticationKey() string {
	config := NewConfigFromFile()
	return config["session_authentication_key"].(string)
}

func SessionEncryptionKey() string {
	config := NewConfigFromFile()
	return config["session_encryption_key"].(string)
}
