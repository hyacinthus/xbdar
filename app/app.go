package app

import (
	"github.com/hyacinthus/xbdar/app/handler"
	"github.com/hyacinthus/xbdar/app/model"
	"github.com/hyacinthus/xbdar/app/utils/xerr"
	"github.com/hyacinthus/xbdar/app/utils/xpage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	config *Config
	e      *echo.Echo
)

// Init app, database, stores...
func Init(debug bool, conf *Config) {
	config = conf

	// db
	model.Init(debug, &config.DB)

	// echo
	e = echo.New()
	e.Debug = debug

	e.HTTPErrorHandler = xerr.ErrorHandler

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(xpage.Middleware(config.PageSize))

	// routes
	// dashboard
	e.GET("/dashboards", handler.GetDashboards)
	e.GET("/dashboards/:id", handler.GetDashboard)

	// chart
	e.GET("/charts", handler.GetCharts)
	e.GET("/charts/:id", handler.GetChart)
	e.GET("/charts/:id/data", handler.FetchChartData)
}

// Run start echo server.
func Run(address string) {
	defer clean()
	e.Logger.Fatal(e.Start(address))
}

func clean() {
	model.Clean()
}
