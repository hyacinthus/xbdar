package service

import (
	"github.com/hyacinthus/xbdar/model"
)

// FetchChartData fetch a chart's data
func FetchChartData(id string) (interface{}, error) {
	chart, err := model.GetChartWithDatasourceByID(id)
	if err != nil {
		return nil, err
	}
	ds := chart.Datasource
	dsFetcher, err := NewDataFetcher(ds.Type, ds.ParamJSON, chart.DataParamJSON)
	if err != nil {
		return nil, err
	}
	return dsFetcher.Fetch()
}
