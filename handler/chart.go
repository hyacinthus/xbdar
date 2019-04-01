package handler

import (
	"fmt"
	"net/http"

	"github.com/hyacinthus/x/xerr"

	"github.com/hyacinthus/xbdar/model"
	"github.com/hyacinthus/xbdar/service"

	"github.com/labstack/echo"
)

// GetDashboards 分页获取dashboard info
func GetDashboards(c echo.Context) error {
	// TODO
	c.JSON(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
	return nil
}

// GetDashboard 获取dashboard info
func GetDashboard(c echo.Context) error {
	// TODO
	c.JSON(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
	return nil
}

// GetCharts 分页获取chart info
func GetCharts(c echo.Context) error {
	page := c.Get("page").(int)
	limit := c.Get("limit").(int)
	charts, err := model.GetCharts(page, limit)
	if err != nil {
		return xerr.New(400, "request error", fmt.Sprintf("get charts: %v", err))
	}
	c.JSON(http.StatusOK, charts)
	return nil
}

// GetChart get a chart's info.
func GetChart(c echo.Context) error {
	id := c.Param("id")
	chart, err := model.GetChartByID(id)
	if err != nil {
		return xerr.New(400, "request error", fmt.Sprintf("fetch chart data #%s: %v", id, err))
	}
	c.JSON(http.StatusOK, chart)
	return nil
}

// FetchChartData fetch a chart's data.
func FetchChartData(c echo.Context) error {
	id := c.Param("id")
	data, err := service.FetchChartData(id)
	if err != nil {
		return xerr.New(400, "request error", fmt.Sprintf("fetch chart data #%s: %v", id, err))
	}
	c.JSON(http.StatusOK, data)
	return nil
}
