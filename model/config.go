package model

// Config database configs.
type Config struct {
	Dialect    string `default:"sqlite3"`
	ConnString string `default:"test.db"`
	Lifetime   int    `default:"3000"`
}
