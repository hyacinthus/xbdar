package main

import (
	"github.com/hyacinthus/xbdar/app"
)

// Config is the all settings of this command
// 与主程序配置文件结构保持一致，可以复用主程序的配置文件
type Config struct {
	Debug bool `default:"false"`

	APP app.Config
}
