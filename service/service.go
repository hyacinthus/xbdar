package service

import (
	"fmt"

	"github.com/hyacinthus/xbdar/model"
	"github.com/mitchellh/mapstructure"
)

// FetchChartData fetch a chart's data
func FetchChartData(id string) (interface{}, error) {
	chart, err := model.GetChartWithDatasourceByID(id)
	if err != nil {
		return nil, err
	}
	datasource := chart.Datasource
	dsType := datasource.Type
	var dsFetcher DataFetcher
	switch dsType {
	case "file.json":
		ds := new(DatasourceJSONFile)
		mapstructure.Decode(datasource.ParamJSON, ds)
		mapstructure.Decode(chart.DataParamJSON, &(ds.Param))
		dsFetcher = ds
	case "file.yaml":
	case "db.sqlite3", "db.mysql", "db.postgres":
	default:
		return nil, fmt.Errorf("unsupported datasource type: %s", dsType)
	}
	return dsFetcher.Fetch()
}
