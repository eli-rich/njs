package main

import (
	"encoding/json"
	"log"
	"os"
)

func loadConfig() Config {
	configFile, err := os.ReadFile("config.json")
	if os.IsNotExist(err) {
		log.Fatalln("config.json not found")
	} else if err != nil {
		log.Fatalln(err)
	}

	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		log.Fatalln(err)
	}
	return config
}
