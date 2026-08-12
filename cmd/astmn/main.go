package main

import (
	"astmn/internal/config"
	"astmn/internal/db"
	"astmn/internal/log"
	"astmn/internal/ui"
)

var c *config.Config

func main() {
	_, err := log.InitLogger(c.LogsDir)
	if err != nil {
		ui.PError("unable init logger: " + err.Error())
		return
	}

	log.Info("starting...")

	log.Info("loading data from config...")
	c, err = config.Load("./astmn.conf.yml")
	if err != nil {
		log.Errorf("unable to load config: %v", err)
	}
	log.Info("done!")

	log.Info("opening db pool...")
	db.OpenPool(c.DBPath)
	log.Info("done!")
}
