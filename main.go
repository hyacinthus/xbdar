package main

import (
	"github.com/hyacinthus/x/page"
	"github.com/hyacinthus/x/xerr"
	"github.com/hyacinthus/x/xlog"
	"github.com/hyacinthus/xbdar/handler"
	"github.com/hyacinthus/xbdar/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// 全局变量
var (
	// config
	config = loadConfig()
	// Logger
	log = xlog.Get()
)

func init() {
	// logger
	if config.Debug {
		xlog.Debug()
	}

	// initialization
	model.Init(config.Debug, &config.DB)
}

func main() {
	defer clean()

	e := echo.New()
	e.HTTPErrorHandler = xerr.ErrorHandler

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(page.Middleware(config.APP.PageSize)) // 分页参数解析，在 pagination.go 定义

	if config.Debug {
		e.Debug = true
	}

	e.GET("/chart/:id", handler.GetChart)
	e.GET("/chart/:id/data", handler.FetchChartData)

	e.Logger.Fatal(e.Start(config.APP.Address))
}

func clean() {
	model.Clean()
}
