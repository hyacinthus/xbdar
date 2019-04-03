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
	// config
	config *Config
	// Logger
	log = xlog.Get()
	// gorm mysql db connection
	db *gorm.DB
)

// Init model, database, stores...
func Init(debug bool, conf *Config) {
	config = conf
	var err error
	for {
		db, err = gorm.Open(config.Dialect, config.DSN)
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

	if config.AutoMigrate {
		// async create tables if needed.
		go CreateTables()
	}
}

// Clean up model resources.
func Clean() {
	db.Close()
}

// CreateTables gorm auto migrate tables
func CreateTables() {
	db.AutoMigrate(&Datasource{}, &Chart{}, &Dashboard{}, &DashboardChart{})
	initData()
}

func initData() {
}

// GetDB returns db object in ops environment.
func GetDB() *gorm.DB {
	if config != nil && config.IsOps {
		return db
	}
	return nil
}

// CleanDB drops all tables.
func CleanDB() {
	db.DropTableIfExists(&Dashboard{}, &Chart{}, &Datasource{}, &DashboardChart{})
}
