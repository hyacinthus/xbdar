package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetChart get a chart's info.
func GetChart(c echo.Context) error {
	id := c.Param("id")
	c.JSON(http.StatusOK, id)
	return nil
}

// FetchChartData fetch a chart's data.
func FetchChartData(c echo.Context) error {
	id := c.Param("id")
	c.JSON(http.StatusOK, id)
	return nil
}
