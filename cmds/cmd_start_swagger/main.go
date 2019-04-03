package main

import (
	"flag"

	"github.com/hyacinthus/xbdar/docs"

	app "github.com/hyacinthus/xbdar/cmds/swagger-app"
	"github.com/hyacinthus/xbdar/utils/xconfig"
)

var (
	config = new(Config)
	a      *app.App
)

func init() {
	rv := new(xconfig.ResetValues)
	xconfig.Load(config, func() {
		flagConfig := new(Config)
		flag.BoolVar(rv.AddBool(&flagConfig.Debug, &config.Debug), "debug", false, "debug?")
		flag.StringVar(rv.AddString(&flagConfig.Address, &config.Address), "a", ":7070", "listening address.")
		flag.BoolVar(rv.AddBool(&flagConfig.APP.Dev, &config.APP.Dev), "dev", false, "in dev mode?")
		flag.StringVar(rv.AddString(&flagConfig.APP.DocFile, &config.APP.DocFile), "d", "./docs/swagger.json", "swagger.json file path.")
		flag.StringVar(rv.AddString(&flagConfig.APP.Host, &config.APP.Host), "host", "localhost:5000", "api host.")
	})
	rv.Reset()

	// change to real host when use 'doc.json'.
	docs.SwaggerInfo.Host = config.APP.Host
	// app
	a = app.CreateApp(config.Debug, &config.APP)
}

func main() {
	defer a.Destroy()
	a.Run(config.Address)
}
