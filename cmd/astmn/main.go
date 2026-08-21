package main

import (
	"os"

	"astmn/internal/config"
	"astmn/internal/db"
	"astmn/internal/log"
	"astmn/internal/ui"
)

var c *config.Config

func main() {
	var err error
	c, err = config.Load("./astmn.conf.yml")
	if err != nil {
		ui.PError("unable to load config" + err.Error())
		os.Exit(1)
	}

	logFile, err := log.InitLogger(c.LogsDir)
	if err != nil {
		ui.PError("unable init logger: " + err.Error())
		os.Exit(1)
	}
	defer logFile.Close()
	log.Info("logger initialized")

	log.Info("starting...")
	log.Info("opening db pool...")
	if err := db.OpenPool(c.DBPath); err != nil {
		log.Errorf("error opening db pool: %v", err)
		os.Exit(1)
	}
	defer db.ClosePool()
	log.Info("done!")

	Execute()
}
