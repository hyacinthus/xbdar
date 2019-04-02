package model

import (
	"github.com/hyacinthus/x/model"
)

// Datasource 数据源
type Datasource struct {
	model.Entity
	Name      string     `json:"name" gorm:"type:varchar(128);not null;unique_index"`
	Domain    string     `json:"domain" gorm:"type:varchar(20);not null"`
	Type      string     `json:"type" gorm:"type:varchar(20);not null"`
	ParamJSON JSONObject `json:"param_json" gorm:"type:text;not null"`
	Charts    []*Chart    `json:"charts,omitempty" gorm:"ForeignKey:DatasourceID"`
}
