package main

import (
	"github.com/hyacinthus/x/xlog"
	"github.com/hyacinthus/xbdar/app"
	"github.com/hyacinthus/xbdar/utils/xconfig"
)

// @title 雪豹商情报表系统API
// @version 0.1.0
// @description 展示商情数据报表

// @contact.name webee
// @contact.url https://github.com/webee
// @contact.email webee.yw@gmail.com

// schemes http https

// @BasePath /

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

	// initialize app
	app.Init(config.Debug, &config.APP)
}

func main() {
	// start
	app.Run(config.Address)
}
