package model

import (
	"CallServer/config"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestDurationAndCost(t *testing.T) {

	err := config.LoadConfigurationFromString(`{
  "server": {
    "port": "",
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
		t.Fatal(err)
	}

	outBoundCall := &Call{
		StartTime: time.Now().Add(time.Duration(-60) * time.Minute),
		EndTime:   time.Now(),
		IsInbound: false,
	}
	outBoundCall.ComputeDurationAndCost()

	assert.Equal(t, outBoundCall.Duration, uint16(60))
	assert.Equal(t, outBoundCall.Cost, uint32(5750000))

	inBoundCall := &Call{
		StartTime: time.Now().Add(time.Duration(-2) * time.Hour),
		EndTime:   time.Now(),
		IsInbound: true,
	}
	inBoundCall.ComputeDurationAndCost()

	assert.Equal(t, inBoundCall.Duration, uint16(120))
	assert.Equal(t, inBoundCall.Cost, uint32(0))
}

func TestCallValidation(t *testing.T) {

	err := config.LoadConfigurationFromString(`{
  "server": {
    "port": "",
    "phone_number_regex": "^(\\+|00)[0-9]{2,}|[0-9]+$"
  }
}`)

	if err != nil {
		t.Fatal(err)
	}

	call := &Call{
		Caller:    "123",
		Callee:    "321",
		StartTime: time.Now().Add(time.Duration(-60) * time.Minute),
		EndTime:   time.Now(),
		IsInbound: false,
	}

	validationErrors := call.Validate()

	assert.Equal(t, len(validationErrors), 0)

	call.Caller = "aaaa"
	call.Callee = "aaaa"
	call.EndTime = call.StartTime.Add(time.Duration(-60) * time.Minute)
	validationErrors = call.Validate()

	assert.Equal(t, len(validationErrors), 4)
}
