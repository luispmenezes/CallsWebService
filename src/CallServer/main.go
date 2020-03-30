package main

import (
	"CallServer/api"
	"CallServer/config"
	"CallServer/persistence"
	"log"
	"os"
)

const DefaultConfigPath = "./server.json"

func main() {
	configFilePath := DefaultConfigPath
	if len(os.Args) == 2 {
		configFilePath = os.Args[1]
	}

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
