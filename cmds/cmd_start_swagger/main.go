package main

import (
	"encoding/json"
	"os"
	"net/http"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/swaggo/echo-swagger"
	"github.com/hyacinthus/xbdar/docs"
)

var (
	debug bool
	dev bool
	address string
	docFile string
	host string
)

func init() {
	// flag must before xconfig load as it use flag too.
	flag.BoolVar(&debug, "debug", false, "debug?")
	flag.BoolVar(&dev, "dev", false, "in dev mode?")
	flag.StringVar(&address, "a", ":7070", "listening address.")
	flag.StringVar(&docFile, "d", "./docs/swagger.json", "swagger.json file path.")
	flag.StringVar(&host, "h", "localhost:5000", "api host.")
	flag.Parse()

	// change to real host when use 'doc.json'.
	docs.SwaggerInfo.Host = host
}

func main() {
	e := echo.New()
	e.Debug = debug

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	e.GET("/docs/swagger.json", func(c echo.Context) error {
		f, err := os.Open(docFile)
		if err != nil {
			return err
		}
		defer f.Close()

		dec := json.NewDecoder(f)
		data := make(map[string]interface{}, 0)
		if err := dec.Decode(&data); err != nil {
			return err
		}
		data["host"] = host
		return c.JSON(http.StatusOK, data)
	})
	if dev {
		e.GET("/swagger/*", echoSwagger.EchoWrapHandler(echoSwagger.URL("/docs/swagger.json")))
	} else {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	e.Logger.Fatal(e.Start(address))
}