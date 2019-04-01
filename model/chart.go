package model

import (
	"github.com/hyacinthus/x/model"
)

// Chart 图表
type Chart struct {
	model.Entity
	Name             string      `json:"name" gorm:"type:varchar(128);not null"`
	Type             string      `json:"type" gorm:"type:varchar(30);not null"`
	DatasourceDomain string      `json:"datasource_domain" gorm:"type:varchar(20);not null"`
	DatasourceType   string      `json:"datasource_type" gorm:"type:varchar(20);not null"`
	Datasource       *Datasource `json:"datasource,omitempty" gorm:"ForeignKey:DatasourceID"`
	DatasourceID     string      `json:"datasource_id" gorm:"type:varchar(20);not null"`
	DataParamJSON    JSONObject  `json:"data_param_json" gorm:"type:text;not null"`
	ChartParamJSON   JSONObject  `json:"chart_param_json" gorm:"type:text;not null"`
	Dashboards       []Dashboard `json:"dashboards,omitempty" gorm:"many2many:dashboard_charts"`
}

// services

// GetCharts 分页获取chart信息
func GetCharts(offset, limit int) ([]Chart, error) {
	charts := make([]Chart, 0)
	q := db.Offset(offset).Limit(limit).Order("")
	return charts, q.Find(&charts).Error
}

// GetChartByID 通过id获取图表基本信息
func GetChartByID(id string) (*Chart, error) {
	chart := new(Chart)
	return chart, db.Find(chart, "id=?", id).Error
}

// GetChartWithDatasourceByID 通过id获取图表基本信息和其数据源信息
func GetChartWithDatasourceByID(id string) (*Chart, error) {
	chart := new(Chart)
	if err := db.Find(chart, "id=?", id).Error; err != nil {
		return nil, err
	}
	chart.Datasource = new(Datasource)
	if err := db.Model(chart).Related(chart.Datasource).Error; err != nil {
		return nil, err
	}
	return chart, nil
}
