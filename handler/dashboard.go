package handler

import (
	"fmt"
	"net/http"

	"github.com/hyacinthus/x/xerr"
	"github.com/hyacinthus/xbdar/model"
	"github.com/labstack/echo"
)

// GetDashboards 分页获取dashboard info
func GetDashboards(c echo.Context) error {
	page := c.Get("page").(int)
	limit := c.Get("limit").(int)
	charts, err := model.GetDashboards(page, limit)
	if err != nil {
		return xerr.New(400, "request error", fmt.Sprintf("get dashboards: %v", err))
	}
	return c.JSON(http.StatusOK, charts)
}

// GetDashboard 获取dashboard info(递归包含所有子报表)
func GetDashboard(c echo.Context) error {
	id := c.Param("id")
	dashboard, err := model.GetDashboardByID(id)
	if err != nil {
		return xerr.New(400, "request error", fmt.Sprintf("get dashboard #%s: %v", id, err))
	}
	return c.JSON(http.StatusOK, dashboard)
}
