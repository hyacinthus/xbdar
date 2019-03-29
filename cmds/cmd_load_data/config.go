package main

import (
	"github.com/hyacinthus/xbdar/model"
)

// Config is the all settings of this command
type Config struct {
	Debug bool `default:"false"`

	DB model.Config
}
