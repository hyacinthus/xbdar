package main

import (
	"github.com/hyacinthus/xbdar/app"
)

// Config is the all settings of this project
// TODO: create a cmd to dump a sample config file.
type Config struct {
	Name    string `default:"xuebao dashboard"`
	Debug   bool   `default:"false"`
	Address string `default:":8080"`

	APP app.Config
}
