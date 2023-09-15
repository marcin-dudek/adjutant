package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Source      string
	Destination string
}

var cfg Config

func initConfig() {
	viper.SetConfigName("adjutant")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath(".")
	viper.SetDefault("source", "C:\\Users\\")
	viper.SetDefault("destination", "C:\\Users\\")
	viper.ReadInConfig()

	cfg = getConfig()
}

func saveDestination(path string) {
	viper.Set("destination", path)
	viper.SafeWriteConfig()
	cfg = getConfig()
}

func saveSource(path string) {
	viper.Set("source", path)
	viper.SafeWriteConfig()
	cfg = getConfig()
}

func getConfig() Config {
	cfg = Config{
		Source:      viper.GetString("source"),
		Destination: viper.GetString("destination"),
	}

	log.Info(log.Fields{
		"step":        "configuration",
		"source":      cfg.Source,
		"destination": cfg.Destination,
	})

	return cfg
}
