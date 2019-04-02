package app

import (
	"github.com/hyacinthus/xbdar/app/model"
)

// Config app configs.
type Config struct {
	PageSize int `default:"10"`

	DB model.Config
}
