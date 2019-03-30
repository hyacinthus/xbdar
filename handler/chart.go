package handler

import (
	"fmt"
	"net/http"

	"github.com/hyacinthus/xbdar/model"
	"github.com/hyacinthus/xbdar/service"

	"github.com/labstack/echo"
)

// GetChart get a chart's info.
func GetChart(c echo.Context) error {
	id := c.Param("id")
	chart, err := model.GetChartByID(id)
	if err != nil {
		return fmt.Errorf("get chart #%s: %v", id, err)
	}
	c.JSON(http.StatusOK, chart)
	return nil
}

// FetchChartData fetch a chart's data.
func FetchChartData(c echo.Context) error {
	id := c.Param("id")
	data, err := service.FetchChartData(id)
	if err != nil {
		return fmt.Errorf("fetch chart data #%s: %v", id, err)
	}
	c.JSON(http.StatusOK, data)
	return nil
}
