package model

import (
	"CallServer/config"
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

	if outBoundCall.Duration != 60 {
		t.Errorf("Invalid outbound duration, expected %d got %d", 60, outBoundCall.Duration)
	}

	if outBoundCall.Cost != 5750000 {
		t.Errorf("Invalid outbound cost, expected %d got %d", 5750000, outBoundCall.Cost)
	}

	inBoundCall := &Call{
		StartTime: time.Now().Add(time.Duration(-2) * time.Hour),
		EndTime:   time.Now(),
		IsInbound: true,
	}
	inBoundCall.ComputeDurationAndCost()

	if inBoundCall.Duration != 120 {
		t.Errorf("Invalid outbound duration, expected %d got %d", 120, inBoundCall.Duration)
	}

	if inBoundCall.Cost != 0 {
		t.Errorf("Invalid outbound cost, expected %d got %d", 0, inBoundCall.Cost)
	}
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

	if len(validationErrors) > 0 {
		t.Errorf("Unexpected validation errors found, expected none got %s", validationErrors)
	}

	call.Caller = "aaaa"
	call.Callee = "aaaa"
	call.EndTime = call.StartTime.Add(time.Duration(-60) * time.Minute)
	validationErrors = call.Validate()

	if len(validationErrors) != 4 {
		t.Errorf("Unexpected validation erros found, expected %d got %s", 4, validationErrors)
	}
}
