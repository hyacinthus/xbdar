package main

import (
	"github.com/hyacinthus/xbdar/model"
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
