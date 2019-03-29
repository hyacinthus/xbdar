package model

import (
	"github.com/hyacinthus/x/model"
)

// Dashboard 报表
type Dashboard struct {
	model.Entity
	Key        string     `gorm:"type:varchar(64);index"`
	Title      string     `gorm:"type:varchar(128);not null"`
	LayoutJSON JSONObject `json:"layout_json" gorm:"type:text;not null"`
}
