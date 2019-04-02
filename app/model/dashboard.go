package model

import (
	"github.com/hyacinthus/x/model"
)

// Dashboard 报表
type Dashboard struct {
	model.Entity
	Key        *string      `json:"key" gorm:"type:varchar(64);index"`
	Title      string       `json:"title" gorm:"type:varchar(128);not null"`
	LayoutJSON JSONObject   `json:"layout_json" gorm:"type:text;not null"`
	Children   []*Dashboard `json:"children,omitempty" gorm:"ForeignKey:ParentID"`
	ParentID   *string      `json:"parent_id" gorm:"type:varchar(20)"`
	Parent     *Dashboard   `json:"parent,omitempty" gorm:"ForeignKey:ParentID"`
	Order      int          `json:"order" gorm:"type:samllint;default:0"`
	Charts     []*Chart     `json:"charts,omitempty" gorm:"many2many:dashboard_charts"`
}
// TODO: 在dashboard_charts中添加覆盖chart中的数据和图表参数

// services

// GetDashboards 分页获取报表基本信息(不包含子报表)
func GetDashboards(page, perPage int) (*Pagination, error) {
	dashboards := make([]Dashboard, 0)
	return Paginate(db.Model(&Dashboard{}).Where("parent_id is ?", nil).Order("id"), page, perPage, &dashboards)
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
	if err := db.Model(&Dashboard{}).Where("parent_id = ?", d.ID).Find(&d.Children).Error; err != nil {
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
