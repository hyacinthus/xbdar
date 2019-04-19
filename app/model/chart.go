package model

import (
	"github.com/hyacinthus/x/model"
	"github.com/webee/x/xpage"
	"github.com/webee/x/xpage/xgorm"
)

type (
	// Entity 实体common
	Entity = model.Entity
)

// Chart 图表
type Chart struct {
	Entity
	Name             string            `json:"name" gorm:"type:varchar(128);not null" example:"基础折线图"`
	Type             string            `json:"type" gorm:"type:varchar(30);not null" example:"line"`
	DatasourceDomain string            `json:"datasource_domain" gorm:"type:varchar(20);not null" example:"db"`
	DatasourceType   string            `json:"datasource_type" gorm:"type:varchar(20);not null" example:"mysql"`
	Datasource       *Datasource       `json:"datasource,omitempty" gorm:"ForeignKey:DatasourceID"`
	DatasourceID     string            `json:"datasource_id" gorm:"type:varchar(20);not null" example:"xxxx"`
	DataParamJSON    JSONObject        `json:"data_param_json" gorm:"type:text;not null"`
	ChartParamJSON   JSONObject        `json:"chart_param_json" gorm:"type:text;not null"`
	Dashboards       []*Dashboard      `json:"dashboards,omitempty" gorm:"many2many:dashboard_charts"`
	ChartDashboards  []*DashboardChart `json:"chart_dashboards,omitempty" gorm:"ForeignKey:ChartID"`
}

// services

// GetCharts 分页获取chart信息
func GetCharts(page, perPage int) (*xpage.Pagination, error) {
	charts := make([]Chart, 0)
	return xgorm.NewPaginator(db.Model(&Chart{}).Order("id"), &charts).Paginate(page, perPage)
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
