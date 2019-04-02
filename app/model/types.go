package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONObject is json object type
type JSONObject map[string]interface{}

// String 转换为string类型
func (j JSONObject) String() string {
	data, err := json.Marshal(j)
	if err != nil {
		return "{}"
	}
	return string(data)
}

// Scan implements the Scanner interface.
func (j *JSONObject) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	*j = make(map[string]interface{}, 0)
	data, ok := src.([]byte)
	if !ok {
		return errors.New("Read JSON from db failed")
	}

	return json.Unmarshal(data, j)
}

// Value implements the driver Valuer interface.
func (j JSONObject) Value() (driver.Value, error) {
	return j.String(), nil
}
