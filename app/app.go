package app

import (
	"github.com/hyacinthus/xbdar/app/handler"
	"github.com/hyacinthus/xbdar/app/model"
	"github.com/hyacinthus/xbdar/app/utils/xerr"
	"github.com/hyacinthus/xbdar/app/utils/xpage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// App is app object
type App struct {
	config *Config
	e      *echo.Echo
}

// Run start echo server.
func (app *App) Run(address string) {
	app.e.Logger.Fatal(app.e.Start(address))
}

// Destroy destroy this app.
func (app *App) Destroy() {
	model.Clean()
}

// CreateApp create a app object.
func CreateApp(debug bool, config *Config) *App {
	app := &App{config, echo.New()}

	// db
	model.Init(debug, &config.DB)

	// echo
	initEcho(app.e, debug, config)

	return app
}

func initEcho(e *echo.Echo, debug bool, config *Config) {
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
