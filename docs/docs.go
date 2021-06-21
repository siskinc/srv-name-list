// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/type": {
            "get": {
                "description": "名单类型查找功能, 通过code, is_valid, 分页",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "名单类型"
                ],
                "summary": "名单类型查找功能",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "是否生效",
                        "name": "is_valid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "名单类型编码",
                        "name": "code",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "页码",
                        "name": "page_index",
                        "in": "query"
                    },
                    {
                        "minimum": 10,
                        "type": "integer",
                        "default": 10,
                        "description": "分页大小",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "正常回包, 回复查询成功的名单类型数据",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResultPaged"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.ListType"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "名单类型查找功能, 通过code, is_valid, 分页",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "名单类型"
                ],
                "summary": "名单类型创建功能",
                "parameters": [
                    {
                        "description": "名单属性",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/list_type.CreateListTypeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "正常回包, 回复创建成功的名单类型数据",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.ListType"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/type/{id}": {
            "delete": {
                "description": "名单类型删除功能, 通过list_type_id删除",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "名单类型"
                ],
                "summary": "名单类型删除功能",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "名单类型id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "正常回包, 回复删除成功的名单类型数据",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.ListType"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "patch": {
                "description": "名单类型查找功能, 通过code, is_valid, 分页",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "名单类型"
                ],
                "summary": "名单类型创建功能",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "名单类型id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "名单属性",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/list_type.UpdateListTypeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "正常回包, 回复更新成功的名单类型数据",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.ListType"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpx.JSONResult": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "回包code，表明是否正确，在code == 0时，表明服务正常",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "object"
                },
                "message": {
                    "description": "回报message，在code != 0时，展示给前端",
                    "type": "string"
                }
            }
        },
        "httpx.JSONResultPaged": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "回包code，表明是否正确，在code == 0时，表明服务正常",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "object"
                },
                "message": {
                    "description": "回报message，在code != 0时，展示给前端",
                    "type": "string"
                },
                "total": {
                    "description": "总数量",
                    "type": "integer"
                }
            }
        },
        "list_type.CreateListTypeReq": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "名单类型编码",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "fields": {
                    "description": "这类名单的值被构建的字段",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "is_valid": {
                    "description": "是否生效",
                    "type": "boolean"
                }
            }
        },
        "list_type.UpdateListTypeReq": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "is_valid": {
                    "description": "是否生效",
                    "type": "boolean"
                }
            }
        },
        "models.ListType": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "名单类型编码",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "fields": {
                    "description": "这类名单的值被构建的字段",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "description": "主键id",
                    "type": "string"
                },
                "is_valid": {
                    "description": "是否生效",
                    "type": "boolean"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8000",
	BasePath:    "/name-list",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server Petstore server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}