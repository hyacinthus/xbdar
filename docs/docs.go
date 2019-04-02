// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-04-03 02:39:51.066792 +0800 CST m=+0.032045170

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "展示商情数据报表",
        "title": "雪豹商情报表系统API",
        "contact": {
            "name": "webee",
            "url": "https://github.com/webee",
            "email": "webee.yw@gmail.com"
        },
        "license": {},
        "version": "0.1.0"
    },
    "host": "{{.Host}}",
    "basePath": "/",
    "paths": {
        "/charts": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chart"
                ],
                "summary": "分页获取图表基本信息",
                "operationId": "get-charts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "第几页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页多少",
                        "name": "per_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Pagination"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/xerr.Error"
                        }
                    }
                }
            }
        },
        "/charts/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chart"
                ],
                "summary": "获取图表基本信息",
                "operationId": "get-chart-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chart ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Chart"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/xerr.Error"
                        }
                    }
                }
            }
        },
        "/charts/{id}/data": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chart"
                ],
                "summary": "拉取图表数据",
                "operationId": "fetch-chart-data-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chart ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "任意类型数据",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/xerr.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Chart": {
            "type": "object",
            "properties": {
                "chart_param_json": {
                    "type": "object",
                    "$ref": "#/definitions/model.JSONObject"
                },
                "dashboards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Dashboard"
                    }
                },
                "data_param_json": {
                    "type": "object",
                    "$ref": "#/definitions/model.JSONObject"
                },
                "datasource": {
                    "type": "object",
                    "$ref": "#/definitions/model.Datasource"
                },
                "datasource_domain": {
                    "type": "string",
                    "example": "db"
                },
                "datasource_id": {
                    "type": "string",
                    "example": "xxxx"
                },
                "datasource_type": {
                    "type": "string",
                    "example": "mysql"
                },
                "name": {
                    "type": "string",
                    "example": "基础折线图"
                },
                "type": {
                    "type": "string",
                    "example": "line"
                }
            }
        },
        "model.Dashboard": {
            "type": "object",
            "properties": {
                "charts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Chart"
                    }
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Dashboard"
                    }
                },
                "key": {
                    "type": "string"
                },
                "layout_json": {
                    "type": "object",
                    "$ref": "#/definitions/model.JSONObject"
                },
                "order": {
                    "type": "integer"
                },
                "parent": {
                    "type": "object",
                    "$ref": "#/definitions/model.Dashboard"
                },
                "parent_id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.Datasource": {
            "type": "object",
            "properties": {
                "charts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Chart"
                    }
                },
                "domain": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "param_json": {
                    "type": "object",
                    "$ref": "#/definitions/model.JSONObject"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.JSONObject": {
            "type": "object",
            "additionalProperties": {
                "type": "object"
            }
        },
        "model.Pagination": {
            "type": "object",
            "properties": {
                "has_next": {
                    "type": "boolean"
                },
                "has_prev": {
                    "type": "boolean"
                },
                "items": {
                    "type": "object"
                },
                "next_page": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "pages": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "prev_page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "xerr.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "description": "错误代码，为英文字符串，前端可用此判断大的错误类型。",
                    "type": "string"
                },
                "message": {
                    "description": "错误消息，为详细错误描述，前端可选择性的展示此字段。",
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
