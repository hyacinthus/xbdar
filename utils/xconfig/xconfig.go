package xconfig

import (
	"flag"
	"os"

	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
)

// Load config settings for config object(should be a struct pointer).
// TODO: 添加dump完整配置文件结构和生成默认结构文件的功能, 作为sample config file.
func Load(config interface{}) {
	var (
		envFile    string
		configFile string
	)
	flag.StringVar(&envFile, "e", ".env", "env file")
	flag.StringVar(&configFile, "c", ".config.yml", "config file")
	flag.Parse()
	godotenv.Load(envFile)
	os.Setenv("CONFIGOR_ENV_PREFIX", "-")
	if fileInfo, err := os.Stat(configFile); err == nil && fileInfo.Mode().IsRegular() {
		configor.Load(config, configFile)
	} else {
		configor.Load(config)
	}
}
