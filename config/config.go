package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	configFilePath = "config.json"
)

type Config struct {
	HeaderMapping map[string]string `json:"header_mapping"`
}

var (
	defaulConfig = &Config{HeaderMapping: make(map[string]string)}
)

func LoadConfig() *Config {
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal("Config file not found")
		return defaulConfig
	}
	defer file.Close()
	config, err := parseConfig(file)
	if err != nil {
		log.Print(err)
		return defaulConfig
	}
	return config
}

func parseConfig(reader io.Reader) (*Config, error) {
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(buf, &config)

	if err != nil {
		return nil, err
	}
	return &config, nil
}
