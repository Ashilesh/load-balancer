package config

import (
	"encoding/json"
	"os"
)

var config *Configuration

type Configuration struct {
	Host        string   `json:"host"`
	Port        string   `json:"port"`
	NetworkType string   `json:"networkType"`
	Nodes       []string `json:"nodes"`
}

func GetConfig() Configuration {
	if config != nil {
		return *config
	}

	file, err := os.Open("config.json")
	if err != nil {
		panic("cannot open config file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var configuration Configuration

	if err := decoder.Decode(&configuration); err != nil {
		panic("cannot decode config file")
	}

	config = &configuration

	return *config
}
