package handler

import (
	"fmt"
	"net/http"

	"github.com/hyacinthus/xbdar/app/model"
	"github.com/hyacinthus/xbdar/app/service"
	"github.com/hyacinthus/xbdar/app/utils/xerr"
	"github.com/labstack/echo/v4"
)

// GetDashboards 分页获取dashboard info
func GetDashboards(c echo.Context) error {
	page := c.Get("page").(int)
	limit := c.Get("limit").(int)
	charts, err := model.GetDashboards(page, limit)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, charts)
}

// GetDashboard 获取dashboard info(递归包含所有子报表)
func GetDashboard(c echo.Context) error {
	id := c.Param("id")
	dashboard, err := model.GetDashboardByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dashboard)
}

// GetDashboardCharts 分页获取报表的图表基本信息
func GetDashboardCharts(c echo.Context) error {
	id := c.Param("id")
	page := c.Get("page").(int)
	perPage := c.Get("perPage").(int)

	dashboardCharts, err := model.GetDashboardCharts(id, page, perPage)
	if err != nil {
		return xerr.New(400, "request error", fmt.Sprintf("get charts: %v", err))
	}
	return c.JSON(http.StatusOK, dashboardCharts)
}

// GetDashboardChart 获取dashboard所属的chart
func GetDashboardChart(c echo.Context) error {
	id := c.Param("id")
	chartID := c.Param("chart_id")

	dashboardChart, err := model.GetDashboardChartByID(id, chartID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dashboardChart)
}

// GetDashboardChartData 获取dashboard所属的chart的数据
func GetDashboardChartData(c echo.Context) error {
	id := c.Param("id")
	chartID := c.Param("chart_id")

	data, err := service.FetchDashboardChartData(id, chartID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}
