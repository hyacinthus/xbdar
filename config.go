package main

import (
	"flag"
	"os"

	"github.com/hyacinthus/xbdar/model"

	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
)

// Config is the all settings of this project
type Config struct {
	Debug bool `default:"false"`

	APP struct {
		Name     string `default:"xuebao dashboard"`
		Address  string `default:":8080"`
		PageSize int    `default:"10"`
	}

	DB model.Config
}

func loadConfig() *Config {
	config := new(Config)
	envFile := flag.String("e", ".env", "env file")
	configFile := flag.String("c", ".config.yml", "config file")
	flag.Parse()
	godotenv.Load(*envFile)
	os.Setenv("CONFIGOR_ENV_PREFIX", "-")
	if fileInfo, err := os.Stat(*configFile); err == nil && fileInfo.Mode().IsRegular() {
		configor.Load(config, *configFile)
	} else {
		configor.Load(config)
	}

	return config
}
