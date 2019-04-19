package handler

import (
	// for swag
	_ "github.com/webee/x/xpage"

	"fmt"
	"net/http"

	"github.com/webee/x/xerr"

	"github.com/hyacinthus/xbdar/app/model"
	"github.com/hyacinthus/xbdar/app/service"

	"github.com/labstack/echo/v4"
)

// GetCharts 分页获取图表基本信息
// @Summary 分页获取图表基本信息
// @ID get-charts
// @Tags Chart
// @Accept  json
// @Produce json
// @Param page query int false "第几页"
// @Param per_page query int false "每页多少"
// @Success 200 {object} xpage.Pagination
// @Router /charts [get]
func GetCharts(c echo.Context) error {
	page := c.Get("page").(int)
	perPage := c.Get("perPage").(int)
	charts, err := model.GetCharts(page, perPage)
	if err != nil {
		return xerr.New(400, "request error", fmt.Sprintf("get charts: %v", err))
	}
	return c.JSON(http.StatusOK, charts)
}

// GetChart 获取图表基本信息
// @Summary 获取图表基本信息
// @ID get-chart-by-id
// @Tags Chart
// @Accept  json
// @Produce json
// @Param id path string true "Chart ID"
// @Success 200 {object} model.Chart
// @Router /charts/{id} [get]
func GetChart(c echo.Context) error {
	id := c.Param("id")
	chart, err := model.GetChartByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, chart)
}

// FetchChartData 拉取图表数据
// GetChart 拉取图表数据
// @Summary 拉取图表数据
// @ID fetch-chart-data-by-id
// @Tags Chart
// @Accept  json
// @Produce json
// @Param id path string true "Chart ID"
// @Success 200 {string} string "任意类型数据"
// @Router /charts/{id}/data [get]
func FetchChartData(c echo.Context) error {
	id := c.Param("id")

	data, err := service.FetchChartData(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}
