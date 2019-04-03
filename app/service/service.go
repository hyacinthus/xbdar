package service

import (
	"github.com/hyacinthus/xbdar/app/model"
)

// FetchDashboardChartData fetch a dashboard's chart's data
func FetchDashboardChartData(dashboardID, chartID string) (interface{}, error) {
	chart, err := getDashboardChart(dashboardID, chartID, true)
	if err != nil {
		return nil, err
	}
	return fetchChartData(chart)
}

// FetchChartData fetch a chart's data
func FetchChartData(chartID string) (interface{}, error) {
	chart, err := model.GetChartWithDatasourceByID(chartID)
	if err != nil {
		return nil, err
	}
	return fetchChartData(chart)
}

func fetchChartData(chart *model.Chart) (interface{}, error) {
	ds := chart.Datasource
	dsFetcher, err := NewDataFetcher(ds.Domain, ds.Type, ds.ParamJSON, chart.DataParamJSON)
	if err != nil {
		return nil, err
	}
	return dsFetcher.Fetch()
}

func copyMap(dest map[string]interface{}, sources ...map[string]interface{}) {
	for _, src := range sources {
		for k, v := range src {
			dest[k] = v
		}
	}
}

func getDashboardChart(dashboardID, chartID string, withDatasource bool) (*model.Chart, error) {
	dashboardChart, err := model.GetDashboardChartByID(dashboardID, chartID)
	if err != nil {
		return nil, err
	}

	var chart *model.Chart
	if withDatasource {
		chart, err = model.GetChartWithDatasourceByID(chartID)
		if err != nil {
			return nil, err
		}
	} else {
		chart, err = model.GetChartByID(chartID)
	}

	copyMap(chart.DataParamJSON, dashboardChart.DataParamJSON)

	return chart, nil
}
