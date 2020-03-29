package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Dbname   string `json:"dbname"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"database"`

	Server struct {
		Port             string `json:"port"`
		PhoneNumberRegex string `json:"phone_number_regex"`
		CallCost         struct {
			InboundPrice1          uint32 `json:"inbound_price_1"`
			InboundPrice2          uint32 `json:"inbound_price_2"`
			InboundPriceThreshold  uint16 `json:"inbound_price_threshold"`
			OutboundPrice1         uint32 `json:"outbound_price_1"`
			OutboundPrice2         uint32 `json:"outbound_price_2"`
			OutboundPriceThreshold uint16 `json:"outbound_price_threshold"`
		} `json:"call_cost"`
	} `json:"server"`
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
