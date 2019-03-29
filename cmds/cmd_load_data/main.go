package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/hyacinthus/xbdar/model"
	"github.com/hyacinthus/xbdar/utils/xconfig"

	"github.com/jinzhu/gorm"
)

var (
	config   = new(Config)
	dataFile string

	// 希望直接操作db
	db *gorm.DB
)

func init() {
	// flag must before xconfig load as it use flag too.
	flag.StringVar(&dataFile, "d", "-", "data json file, '-' for reset database.")
	xconfig.Load(config)
	flag.Parse()

	config.DB.SetIsOps(true)

	// initialization
	model.Init(config.Debug, &config.DB)
	model.CleanDB()
	model.CreateTables()
	db = model.GetDB()
}

func main() {
	defer clean()

	if dataFile == "-" {
		return
	}

	f, err := os.Open(dataFile)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	var data = new(Data)
	if err := dec.Decode(data); err != nil {
		log.Println(err)
		panic(err)
	}

	loadData(data)
}

func loadData(data *Data) {
	// datasources
	datasources := make(map[string]*model.Datasource)
	for _, d := range data.Datasources {
		id := d.ID
		d.ID = ""
		datasources[id] = d
	}
	for _, v := range datasources {
		db.Create(v)
	}
	// charts
	charts := make(map[string]*model.Chart)
	for _, d := range data.Charts {
		id := d.ID
		d.ID = ""
		d.DatasourceID = datasources[d.DatasourceID].ID
		charts[id] = d
	}
	for _, v := range charts {
		db.Create(v)
	}
	// dashboards
	dashboards := make(map[string]*model.Dashboard)
	for _, d := range data.Dashboards {
		id := d.ID
		v := &d.Dashboard
		dashboards[id] = v

		db.Create(v)
		for _, chartID := range d.ChartIDs {
			db.Model(v).Association("Charts").Append(charts[chartID])
		}
	}
}

func clean() {
	model.Clean()
}

// Data 数据文件格式
type Data struct {
	Datasources []*model.Datasource
	Charts      []*model.Chart
	Dashboards  []*struct {
		model.Dashboard
		ChartIDs []string `json:"chart_ids"`
	}
}
