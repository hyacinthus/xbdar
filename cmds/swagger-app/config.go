package swaggerapp

// Config app configs.
type Config struct {
	Dev         bool   `default:"false"`
	DocFile     string `default:"./docs/swagger.json"`
	DocPath     string `default:"/docs/swagger.json"`
	Host        string `default:"localhost:5000"`
	SwaggerPath string `default:"/swagger/"`
}
