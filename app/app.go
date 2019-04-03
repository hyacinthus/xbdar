package app

import (
	"net/http"

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
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	xpageMiddleware := xpage.Middleware(config.PageSize)

	// routes
	// app
	e.GET("/status", getStatus)

	dashboards := e.Group("/dashboards")
	charts := e.Group("/charts")
	// dashboard
	dashboards.GET("", handler.GetDashboards)
	dashboards.GET("/:id", handler.GetDashboard)
	dashboards.GET("/:id/charts", handler.GetDashboardCharts, xpageMiddleware)
	dashboards.GET("/:id/charts/:chart_id", handler.GetDashboardChart)
	dashboards.GET("/:id/charts/:chart_id/data", handler.GetDashboardChartData)

	// chart
	charts.GET("", handler.GetCharts, xpageMiddleware)
	charts.GET("/:id", handler.GetChart)
	charts.GET("/:id/data", handler.FetchChartData)
}

// API状态 成功204 失败500
func getStatus(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
