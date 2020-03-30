package main

import (
	"CallClient/callws"
	"CallClient/config"
	"CallClient/model"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const DefaultConfigPath = "./client.json"

func main() {
	configFilePath := DefaultConfigPath
	if len(os.Args) == 2 {
		configFilePath = os.Args[1]
	}

	err := config.LoadConfigurationFromPath(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	serverConfig := config.GetConfiguration().Server
	simConfig := config.GetConfiguration().Simulation

	var client = *callws.NewClient(serverConfig.Host, serverConfig.Port, serverConfig.Scheme)

	if simConfig.WipeOnStart{
		log.Println("Wiping all calls")
		_,err = client.RemoveCalls(*model.NewFilter())
		if err != nil{
			log.Fatal(err)
		}
	}

	var startDate,endDate time.Time

	startDate,err = time.Parse(time.RFC3339,simConfig.StartDate)
	if err != nil{
		log.Fatal(err)
	}

	endDate,err = time.Parse(time.RFC3339,simConfig.EndDate)
	if err != nil{
		log.Fatal(err)
	}

	if !startDate.Before(endDate){
		log.Fatal("Invalid Date Range")
	}

	log.Printf("Generating %d Random Phone Numbers", simConfig.NumberOfAgents)
	phoneNumbers := generatePhoneNumbers(simConfig.NumberOfAgents)
	log.Println(phoneNumbers)

	log.Printf("Generating %d Random Calls", simConfig.NumberOfCalls)
	callList := generateRandomCalls(simConfig.NumberOfCalls, phoneNumbers, startDate, endDate)
	log.Println(callList)

	_,err = client.AddCalls(callList)
	if err != nil{
		log.Fatal(err)
	}
	log.Println(client.GetMetadata())
}

func generateRandomCalls(n int, phoneNumbers []string, startDate, endDate time.Time) []model.Call {
	var callList []model.Call

	for i := 0; i < n; i++ {
		callerIdx := rand.Intn(len(phoneNumbers))
		calleeIdx := 0
		for {
			calleeIdx = rand.Intn(len(phoneNumbers))
			if callerIdx != calleeIdx {
				break
			}
		}
		callStartDate := randomDate(startDate, endDate)
		callEndDate := callStartDate.Add(time.Duration(rand.Intn(120)) * time.Minute)
		isInbound := rand.Intn(1) == 1

		callList = append(callList, model.Call{
			Caller:    phoneNumbers[callerIdx],
			Callee:    phoneNumbers[calleeIdx],
			StartTime: callStartDate,
			EndTime:   callEndDate,
			IsInbound: isInbound,
		})
	}

	return callList
}

func generatePhoneNumbers(n int) []string {
	var phoneNumberList []string

	seed := rand.Intn(999999)

	for i := 0; i < n; i++ {
		phoneNumberList = append(phoneNumberList, fmt.Sprint(seed+i))
	}

	return phoneNumberList
}

func randomDate(startDate, endDate time.Time) time.Time {
	min := startDate.Unix()
	max := endDate.Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
