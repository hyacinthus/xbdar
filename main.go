package main

import (
	"github.com/hyacinthus/x/page"
	"github.com/hyacinthus/x/xerr"
	"github.com/hyacinthus/x/xlog"
	"github.com/hyacinthus/xbdar/handler"
	"github.com/hyacinthus/xbdar/model"
	"github.com/hyacinthus/xbdar/utils/xconfig"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// 全局变量
var (
	// config
	config = new(Config)
	// Logger
	log = xlog.Get()
)

func init() {
	// config
	xconfig.Load(config)

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

	// routes
	e.GET("/charts/:id", handler.GetChart)
	e.GET("/charts/:id/data", handler.FetchChartData)

	// start
	e.Logger.Fatal(e.Start(config.APP.Address))
}

func clean() {
	model.Clean()
}
