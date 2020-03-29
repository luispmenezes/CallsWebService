package main

import (
	"CallServer/api"
	"CallServer/config"
	"CallServer/persistence"
	"log"
)

const DefaultConfigPath = "./config.json"

func main() {
	configFilePath := DefaultConfigPath

	err := config.LoadConfigurationFromPath(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	conf := config.GetConfiguration()

	persistenceManager := persistence.NewPGManager(conf.Database.Host, conf.Database.Port, conf.Database.Dbname,
		conf.Database.User, conf.Database.Password)
	apiController := api.NewBaseController(persistenceManager)
	err = apiController.Start(conf.Server.Port)

	if err != nil {
		log.Fatal(err)
	}
}
