package swaggerapp

import (
	"encoding/json"
	"os"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/swaggo/echo-swagger"
)

// App is app object
type App struct {
	config *Config
	e *echo.Echo
}

// Run start server
func (app *App)Run(address string) {
	app.e.Logger.Fatal(app.e.Start(address))
}

// Destroy destroy this app.
func (app *App) Destroy() {
}

// CreateApp create a app object
func CreateApp(debug bool, config *Config) *App {
	app := &App{config, echo.New()}

	app.e.Debug = debug
	// routers

	app.e.Use(middleware.Logger())
	app.e.Use(middleware.Recover())
	app.e.Use(middleware.CORS())

	app.e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	app.e.GET(config.DocPath, func(c echo.Context) error {
		f, err := os.Open(config.DocFile)
		if err != nil {
			return err
		}
		defer f.Close()

		dec := json.NewDecoder(f)
		data := make(map[string]interface{}, 0)
		if err := dec.Decode(&data); err != nil {
			return err
		}
		data["host"] = config.Host
		return c.JSON(http.StatusOK, data)
	})
	if config.Dev {
		app.e.GET(config.SwaggerPath + "*", echoSwagger.EchoWrapHandler(echoSwagger.URL(config.DocPath)))
	} else {
		app.e.GET(config.SwaggerPath + "*", echoSwagger.WrapHandler)
	}

	return app
}