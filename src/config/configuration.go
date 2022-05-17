package config

import (
	"encoding/json"
	"github.com/Ashilesh/balancer/src/utils"
	"os"
)

var config *Configuration

type Configuration struct {
	Host                 string   `json:"host"`
	Port                 string   `json:"port"`
	NetworkType          string   `json:"networkType"`
	PersistentConnection bool     `json:"persistentConnection"`
	Nodes                []string `json:"nodes"`
}

func GetConfig() Configuration {
	if config != nil {
		return *config
	}

	c := createNewConfig()
	setConfig(&c)

	return *config
}

func createNewConfig() Configuration {
	// TODO: create struct to store strings for command fields ex. config = "-config"
	arg, err := utils.GetCmdArgs("-config")
	if err != nil {
		panic("Configuration file path argument not found")
	}

	file, err := os.Open(arg)
	if err != nil {
		panic("cannot open config file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var configuration Configuration

	if err := decoder.Decode(&configuration); err != nil {
		panic("cannot decode config file")
	}

	return configuration
}

func setConfig(conf *Configuration) {
	config = conf
}
