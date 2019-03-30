package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// DataFetcher fetch data
type DataFetcher interface {
	Fetch() (interface{}, error)
}

// NewDataFetcher create a data fetcher.
func NewDataFetcher(dsDomain, dsType string, dsParam, dataParam map[string]interface{}) (DataFetcher, error) {
	var df DataFetcher
	switch dsDomain {
	case "file":
		df = &DatasourceFile{Type: dsType}
	case "db":
		df = &DatasourceDB{Driver: dsType}
	default:
		return nil, fmt.Errorf("unsupported datasource domain: %s", dsDomain)
	}
	if err := mapstructure.Decode(dsParam, df); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(dataParam, df); err != nil {
		return nil, err
	}
	return df, nil
}

// DatasourceFile is a type file as datasource
type DatasourceFile struct {
	Type  string
	Path  string
	Param struct {
		KeyPath string `mapstructure:"key_path"`
	} `mapstructure:",squash"`
}

// Decoder decode to v
type Decoder interface {
	Decode(v interface{}) error
}

// Fetch DataFetcher interface
func (ds *DatasourceFile) Fetch() (interface{}, error) {
	reader, err := openPathForRead(ds.Path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var dec Decoder
	switch ds.Type {
	case "json":
		dec = json.NewDecoder(reader)
	case "yaml":
		dec = yaml.NewDecoder(reader)
	default:
		return nil, fmt.Errorf("unsupported file type: %s", ds.Type)
	}

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

// DatasourceDB is a database as datasource
type DatasourceDB struct {
	Driver     string
	ConnString string `mapstructure:"conn_string"`
	Param      struct {
		sql string
	} `mapstructure:",squash"`
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
