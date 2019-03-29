package model

// Config database configs.
type Config struct {
	// 运维模式
	IsOps bool `default:"false"`

	// db info
	Dialect    string `default:"sqlite3"`
	ConnString string `default:"test.db"`
	Lifetime   int64  `default:"3000"`
}

// SetIsOps set ops env, !!use this only in ops condition
func (c *Config) SetIsOps(v bool) {
	c.IsOps = v
}
