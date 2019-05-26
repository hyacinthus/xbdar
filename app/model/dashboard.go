package model

import (
	"time"

	"github.com/webee/x/xpage"
	"github.com/webee/x/xpage/xgorm"
)

// Dashboard 报表
type Dashboard struct {
	Entity
	Key             *string           `json:"key" gorm:"type:varchar(64);index"`
	Title           string            `json:"title" gorm:"type:varchar(128);not null"`
	LayoutJSON      JSONObject        `json:"layout_json" gorm:"type:text;not null"`
	Children        []*Dashboard      `json:"children,omitempty" gorm:"ForeignKey:ParentID"`
	ParentID        *string           `json:"parent_id" gorm:"type:varchar(20)"`
	Parent          *Dashboard        `json:"parent,omitempty" gorm:"ForeignKey:ParentID"`
	Order           int               `json:"order" gorm:"type:samllint;default:0"`
	Charts          []*Chart          `json:"charts,omitempty" gorm:"many2many:dashboard_charts"`
	DashboardCharts []*DashboardChart `json:"dashboard_charts,omitempty" gorm:"ForeignKey:DashboardID"`
}

// DashboardChart 报表-图表关联表
type DashboardChart struct {
	Dashboard     *Dashboard `json:"dashboard" gorm:"ForeighKey:DashboardID"`
	DashboardID   string     `json:"dashboard_id" gorm:"type:varchar(20);primary_key"`
	Chart         *Chart     `json:"chart" gorm:"Foreign:ChartID"`
	ChartID       string     `json:"chart_id" gorm:"type:varchar(20);primary_key"`
	DataParamJSON JSONObject `json:"data_param_json" gorm:"type:text"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 最后更新时间
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName set the gorm model table name.
func (*DashboardChart) TableName() string {
	return "dashboard_charts"
}

// services

// GetDashboards 分页获取报表基本信息(不包含子报表)
func GetDashboards(page, perPage int) (*xpage.Pagination, error) {
	dashboards := make([]Dashboard, 0)
	return xgorm.NewPaginator(db.Model(&Dashboard{}).Where("parent_id is ?", nil).Order("id"), &dashboards).Paginate(page, perPage)
}

// GetDashboardByID 通过id获取报表基本信息（递归包含所有子报表）
func GetDashboardByID(id string) (*Dashboard, error) {
	dashboard := new(Dashboard)
	if err := db.Find(dashboard, "id=?", id).Error; err != nil {
		return nil, err
	}

	if err := fetchDashboardChildren(dashboard); err != nil {
		return nil, err
	}
	return dashboard, nil
}

func fetchDashboardChildren(d *Dashboard) error {
	d.Children = make([]*Dashboard, 0)
	if err := db.Model(&Dashboard{}).Where("parent_id=?", d.ID).Find(&d.Children).Error; err != nil {
		return err
	}
	// TODO: 并行获取
	for _, dc := range d.Children {
		if err := fetchDashboardChildren(dc); err != nil {
			return err
		}
	}
	return nil
}

// GetDashboardCharts 分页获取报表chart信息
func GetDashboardCharts(dashboardID string, page, perPage int) (*xpage.Pagination, error) {
	dashboardCharts := make([]DashboardChart, 0)
	return xgorm.NewPaginator(db.Model(&DashboardChart{}).Preload("Chart").Where("dashboard_id=?", dashboardID).Order("id"), &dashboardCharts).Paginate(page, perPage)
}

// GetDashboardChartByID 通过id获取报表-图表信息
func GetDashboardChartByID(dashboardID, chartID string) (*DashboardChart, error) {
	dashboardChart := new(DashboardChart)
	return dashboardChart, db.Preload("Chart").Where("dashboard_id=?", dashboardID).Where("chart_id=?", chartID).Find(dashboardChart).Error
}
