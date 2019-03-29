package model

import (
	"time"

	"github.com/hyacinthus/x/xlog"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// Logger
	log = xlog.Get()
	// gorm mysql db connection
	db *gorm.DB
)

// Init model, database, stores...
func Init(debug bool, config *Config) {
	var err error
	for {
		db, err = gorm.Open(config.Dialect, config.ConnString)
		if err != nil {
			log.WithError(err).Warn("waiting for connect to database")
			time.Sleep(time.Second * 2)
			continue
		}
		db.DB().SetConnMaxLifetime(time.Duration(config.Lifetime) * time.Second)
		log.Info("database connect successful.")
		break
	}

	// gorm debug log
	if debug {
		db.LogMode(true)
	}
}

// Clean up model resources.
func Clean() {
	db.Close()
}

// createTable gorm auto migrate tables
func createTables() {
	db.AutoMigrate(&Datasource{})
	initData()
}

func initData() {
}
