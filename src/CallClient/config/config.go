package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Server struct {
		Host   string `json:"host"`
		Port   string `json:"port"`
		Scheme string `json:"scheme"`
	} `json:"server"`
	Simulation struct {
		WipeOnStart    bool   `json:"wipe_on_start"`
		NumberOfAgents int    `json:"num_agents"`
		NumberOfCalls  int    `json:"num_calls"`
		StartDate      string `json:"start_date"`
		EndDate        string `json:"end_date"`
	} `json:"simulation"`
}

var configuration Config

func LoadConfigurationFromPath(confFilePath string) error {
	log.Printf("Loading Configuration file (%s)", confFilePath)

	absConfPath, err := filepath.Abs(confFilePath)
	if err != nil {
		return err
	}
	configFile, err := os.Open(absConfPath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	return jsonParser.Decode(&configuration)
}

func LoadConfigurationFromString(confJSONStr string) error {
	log.Printf("Loading Configuration String (%s)", confJSONStr)
	return json.Unmarshal([]byte(confJSONStr), &configuration)
}

func GetConfiguration() Config {
	return configuration
}
