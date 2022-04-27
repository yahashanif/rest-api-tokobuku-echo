package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func GetConfig() Configuration {

	jsonFile, _ := os.Open("config/config.json")
	jsonString, _ := ioutil.ReadAll(jsonFile)

	config := Configuration{}
	err := json.Unmarshal(jsonString, &config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config
}
