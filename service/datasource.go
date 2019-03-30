package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// DataFetcher fetch data
type DataFetcher interface {
	Fetch() (interface{}, error)
}

// DatasourceJSONFile is a json file as datasource
type DatasourceJSONFile struct {
	Path  string
	Param struct {
		KeyPath string `mapstructure:"key_path"`
	} `mapstructure:",squash"`
}

// NewDataFetcher create a data fetcher.
func NewDataFetcher(dsType string, dsParam, dataParam map[string]interface{}) (DataFetcher, error) {
	var df DataFetcher
	switch dsType {
	case "file.json":
		df = new(DatasourceJSONFile)
	case "file.yaml":
		df = new(DatasourceYAMLFile)
	case "db.sqlite3", "db.mysql", "db.postgres":
		df = new(DatasourceDB)
	default:
		return nil, fmt.Errorf("unsupported datasource type: %s", dsType)
	}
	if err := mapstructure.Decode(dsParam, df); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(dataParam, df); err != nil {
		return nil, err
	}
	return df, nil
}

// Fetch DataFetcher interface
func (ds *DatasourceJSONFile) Fetch() (interface{}, error) {
	reader, err := openPathForRead(ds.Path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	dec := json.NewDecoder(reader)
	var data interface{}
	if err := dec.Decode(&data); err != nil {
		return nil, err
	}
	keys := strings.Split(ds.Param.KeyPath, ".")
	for _, key := range keys {
		if key == "" {
			continue
		}

		m, ok := data.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("json data structure error, no path: '%s'", ds.Param.KeyPath)
		}
		data = m[key]
	}
	return data, nil
}

// DatasourceYAMLFile is a yaml file as datasource
type DatasourceYAMLFile struct {
	Path  string
	Param struct {
		KeyPath string `mapstructure:"key_path"`
	}
}

// Fetch DataFetcher interface
func (ds *DatasourceYAMLFile) Fetch() (interface{}, error) {
	return nil, nil
}

// DatasourceDB is a database as datasource
type DatasourceDB struct {
	driver string
	Path   string
	Param  struct {
		sql string
	}
}

// Fetch DataFetcher interface
func (ds *DatasourceDB) Fetch() (interface{}, error) {
	return nil, nil
}

func openPathForRead(path string) (interface {
	io.Reader
	io.Closer
}, error) {
	if strings.HasPrefix(path, "http") {
		res, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		return res.Body, nil
	}
	return os.Open(path)
}
