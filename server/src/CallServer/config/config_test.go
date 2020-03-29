package config

import (
	"testing"
)

func TestLoadConfigurationFromString(t *testing.T) {

	err := LoadConfigurationFromString(`{
  "database": {
    "host": "host",
    "port": "5432",
    "dbname": "dbname",
    "user": "username",
    "password": "pwd"
  },
  "server": {
    "port": "8989",
    "phone_number_regex": "^(\\+|00)[0-9]{2,}|[0-9]+$",
    "call_cost": {
      "inbound_price_1": 0,
      "inbound_price_2": 0,
      "inbound_price_threshold": 0,
      "outbound_price_1": 50000,
      "outbound_price_2": 100000,
      "outbound_price_threshold": 5
    }
  }
}`)

	if err != nil {
		t.Errorf("Failed to parse configuration string (error: %s)", err)
	}

	if GetConfiguration().Server.PhoneNumberRegex != "^(\\+|00)[0-9]{2,}|[0-9]+$" {
		t.Errorf("Failed to reading phone number regex expected %s got %s)", "^(\\+|00)[0-9]{2,}|[0-9]+$", GetConfiguration().Server.PhoneNumberRegex)
	}
}
