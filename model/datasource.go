package model

import (
	"github.com/hyacinthus/x/model"
)

// Datasource 数据源
type Datasource struct {
	model.Entity
	Name      string     `gorm:"type:varchar(128);not null;unique_index"`
	Type      string     `gorm:"type:varchar(20);not null"`
	ParamJSON JSONObject `json:"param_json" gorm:"type:text;not null"`
	Charts    []Chart    `gorm:"ForeignKey:DatasourceID"`
}
