package main

import (
	app "github.com/hyacinthus/xbdar/cmds/swagger-app"
)

// Config is the all settings of this command
type Config struct {
	Debug   bool   `default:"false"`
	Address string `default:":7070"`

	APP app.Config
}
