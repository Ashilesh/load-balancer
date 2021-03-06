package config

import (
	"encoding/json"
	"os"

	"github.com/Ashilesh/load-balancer/logs"
	"github.com/Ashilesh/load-balancer/utils"
)

var config *Configuration

type Configuration struct {
	Host                 string   `json:"host"`
	Port                 string   `json:"port"`
	Protocol             string   `json:"protocol"`
	NetworkType          string   `json:"networkType"`
	PersistentConnection bool     `json:"persistentConnection"`
	Nodes                []string `json:"nodes"`
}

func GetConfig() Configuration {
	if config != nil {
		return *config
	}

	c := createNewConfig()
	setConfig(c)

	return *config
}

func createNewConfig() *Configuration {
	// TODO: create struct to store strings for command fields ex. config = "-config"
	arg, err := utils.GetCmdArgs("-config")
	if err != nil {
		logs.Fatal("Configuration file path argument not found")
	}

	file, err := os.Open(arg)
	if err != nil {
		logs.Fatal("cannot open config file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var configuration Configuration

	if err := decoder.Decode(&configuration); err != nil {
		logs.Fatal("cannot decode config file")
	}

	return &configuration
}

func setConfig(conf *Configuration) {
	config = conf
}
