package model

import (
	"github.com/hyacinthus/x/model"
)

// Dashboard 报表
type Dashboard struct {
	model.Entity
	Key        *string     `json:"key" gorm:"type:varchar(64);index"`
	Title      string      `json:"title" gorm:"type:varchar(128);not null"`
	LayoutJSON JSONObject  `json:"layout_json" gorm:"type:text;not null"`
	Children   []Dashboard `json:"children,omitempty" gorm:"ForeignKey:ParentID"`
	ParentID   *string     `json:"parent_id" gorm:"type:varchar(20)"`
	Parent     *Dashboard  `json:"parent,omitempty" gorm:"ForeignKey:ParentID"`
	Charts     []Chart     `json:"charts,omitempty" gorm:"many2many:dashboard_charts"`
}
