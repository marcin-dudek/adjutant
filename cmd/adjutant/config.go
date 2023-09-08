package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type config struct {
	source      string
	destination string
}

func initConfig() config {
	viper.SetConfigName("adjutant")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath(".")
	viper.SetDefault("source", "C:\\Users\\Avanade\\music-test")
	viper.SetDefault("destination", "C:\\Users\\Avanade\\music-output")
	viper.ReadInConfig()

	cfg := config{
		source:      viper.GetString("source"),
		destination: viper.GetString("destination"),
	}

	log.Info(log.Fields{
		"step":        "configuration",
		"source":      cfg.source,
		"destination": cfg.destination,
	})
	return cfg
}
