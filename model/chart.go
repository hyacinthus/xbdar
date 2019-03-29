package model

import (
	"github.com/hyacinthus/x/model"
)

// Chart 图表
type Chart struct {
	model.Entity
	Name           string      `json:"name" gorm:"type:varchar(128);not null"`
	Type           string      `json:"type" gorm:"type:varchar(30);not null"`
	DatasourceType string      `json:"datasource_type" gorm:"type:varchar(20);not null"`
	Datasource     *Datasource `json:"datasource,omitempty" gorm:"ForeignKey:DatasourceID"`
	DatasourceID   string      `json:"datasource_id" gorm:"type:varchar(20);not null"`
	DataParamJSON  JSONObject  `json:"data_param_json" gorm:"type:text;not null"`
	ChartParamJSON JSONObject  `json:"chart_param_json" gorm:"type:text;not null"`
	Dashboards     []Dashboard `json:"dashboards,omitempty" gorm:"many2many:dashboard_charts"`
}
