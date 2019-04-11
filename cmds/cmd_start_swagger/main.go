package main

import (
	"github.com/hyacinthus/xbdar/docs"
	"github.com/webee/x/xswagger/cmd"
)

func main() {
	docs.SwaggerInfo.Host = cmd.GetConfig().APP.Host
	cmd.Start()
}
