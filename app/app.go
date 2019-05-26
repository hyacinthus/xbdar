package app

import (
	"net/http"
	"time"

	"github.com/hyacinthus/xbdar/app/handler"
	"github.com/hyacinthus/xbdar/app/model"
	"github.com/webee/x/xerr"
	"github.com/webee/x/xpage/xecho"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	// App is app object
	App struct {
		config *Config
		e      *echo.Echo
		status Status
	}

	// Status is app's status
	Status struct {
		Startup time.Time `json:"startup"`
	}
)

// Run start echo server.
func (app *App) Run(address string) {
	app.status.Startup = time.Now()
	app.e.Logger.Fatal(app.e.Start(address))
}

// Destroy destroy this app.
func (app *App) Destroy() {
	model.Clean()
}

// CreateApp create an app object.
func CreateApp(debug bool, config *Config) *App {
	app := &App{config: config, e: echo.New()}

	// db
	model.Init(debug, &config.DB)

	// echo
	initEcho(app, app.e, debug, config)

	return app
}

func initEcho(app *App, e *echo.Echo, debug bool, config *Config) {
	e.Debug = debug

	e.HTTPErrorHandler = xerr.ErrorHandler

	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	xpageMiddleware := xecho.Middleware(config.PageSize)

	// routes
	// app
	e.GET("/status", app.getStatus)

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

// getStatus return app's status info
// @Summary 获取应用状态信息
// @ID get-status
// @Tags App
// @Produce json
// @Success 200 {object} app.Status
// @Router /status [get]
func (app *App) getStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, app.status)
}
