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
        "/item": {
            "get": {
                "description": "名单项查找功能, 通过code, is_valid, 分页",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "名单项"
                ],
                "summary": "名单项查找功能",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Code",
                        "name": "code",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "是否生效",
                        "name": "isValid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "PageIndex",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "PageSize",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortedField",
                        "name": "sortedField",
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
                                                "$ref": "#/definitions/models.ListItem"
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
                "description": "名单项创建功能",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "名单项"
                ],
                "summary": "名单项创建功能",
                "parameters": [
                    {
                        "description": "名单属性",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/list_item.CreateListItemInfo"
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
                                            "$ref": "#/definitions/models.ListItem"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/item-hit/all": {
            "post": {
                "description": "名单项命中, 返回命中的所有名单项",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "名单项命中"
                ],
                "summary": "名单项命中",
                "parameters": [
                    {
                        "description": "预命中信息",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/list_item.ItemHitAllReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "正常回包",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httpx.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/list_item.ItemHitAllResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/item-hit/pre": {
            "post": {
                "description": "名单项预命中, 指定某个名单项, 能够判断数据是否能够命中该名单项, 如果不能, 输出不能命中的原因",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "名单项命中"
                ],
                "summary": "名单项预命中",
                "parameters": [
                    {
                        "description": "命中信息",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/list_item.ItemHitPreReq"
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
                                            "$ref": "#/definitions/list_item.ItemHitPreResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/item/{id}": {
            "delete": {
                "description": "名单项删除功能, 通过list_item_id删除",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "名单项"
                ],
                "summary": "名单项删除功能",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "名单项id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpx.JSONResult"
                        }
                    }
                }
            },
            "patch": {
                "description": "名单项修改功能, 通过code, is_valid, 分页",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "名单项"
                ],
                "summary": "名单项修改功能",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "名单项id",
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
                            "$ref": "#/definitions/list_item.UpdateListItemInfo"
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
                                            "$ref": "#/definitions/models.ListItem"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/namespace": {
            "get": {
                "description": "命名空间查找功能",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "命名空间"
                ],
                "summary": "命名空间查找功能",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Code 命名空间code（模糊查询）",
                        "name": "code",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Description 命名空间的描述（模糊查询）",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ID 命名空间ID（精确查询）",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "PageIndex 页码",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "PageSize 分页数量",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "SortedField 排序字段",
                        "name": "sortedField",
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
                                                "$ref": "#/definitions/models.Namespace"
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
                "description": "命名空间创建功能",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "命名空间"
                ],
                "summary": "命名空间创建功能",
                "parameters": [
                    {
                        "description": "名单属性",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/namespace.CreateNamespaceReq"
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
                                            "$ref": "#/definitions/models.Namespace"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/namespace/{id}": {
            "delete": {
                "description": "命名空间删除功能, 通过namespace_id删除",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "命名空间"
                ],
                "summary": "命名空间删除功能",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "命名空间id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpx.JSONResult"
                        }
                    }
                }
            },
            "patch": {
                "description": "命名空间修改功能",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "命名空间"
                ],
                "summary": "命名空间修改功能",
                "parameters": [
                    {
                        "minLength": 1,
                        "type": "string",
                        "description": "命名空间id",
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
                            "$ref": "#/definitions/namespace.UpdateNamespaceReq"
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
                                            "$ref": "#/definitions/models.Namespace"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
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
                        "type": "string",
                        "name": "code",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "name": "isValid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "namespace",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sortedField",
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
                "description": "名单类型创建功能",
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
                            "$ref": "#/definitions/httpx.JSONResult"
                        }
                    }
                }
            },
            "patch": {
                "description": "名单类型修改功能",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "名单类型"
                ],
                "summary": "名单类型修改功能",
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
                    "type": "string",
                    "example": "success"
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
                    "type": "string",
                    "example": "success"
                },
                "total": {
                    "description": "总数量",
                    "type": "integer"
                }
            }
        },
        "list_item.CreateListItemInfo": {
            "type": "object",
            "required": [
                "code",
                "is_valid",
                "namespace",
                "values"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "extra": {
                    "type": "object",
                    "additionalProperties": true
                },
                "is_valid": {
                    "type": "boolean"
                },
                "namespace": {
                    "type": "string"
                },
                "values": {
                    "description": "与list type中的fields一一对应的",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "list_item.ItemHitAllReq": {
            "type": "object",
            "required": [
                "data",
                "namespace"
            ],
            "properties": {
                "code_list": {
                    "description": "名单类型code, 如果不传, 或者传入的长度为0, 默认命中所有code",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "data": {
                    "description": "数据",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                }
            }
        },
        "list_item.ItemHitAllResp": {
            "type": "object",
            "properties": {
                "hit_list_item_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ListItem"
                    }
                }
            }
        },
        "list_item.ItemHitPreReq": {
            "type": "object",
            "required": [
                "data",
                "list_item_id"
            ],
            "properties": {
                "data": {
                    "description": "数据",
                    "type": "string"
                },
                "list_item_id": {
                    "description": "名单项ID",
                    "type": "string"
                }
            }
        },
        "list_item.ItemHitPreResp": {
            "type": "object",
            "properties": {
                "hit": {
                    "type": "boolean"
                },
                "list_item": {
                    "$ref": "#/definitions/models.ListItem"
                },
                "resource": {
                    "type": "string"
                }
            }
        },
        "list_item.UpdateListItemInfo": {
            "type": "object",
            "properties": {
                "extra": {
                    "description": "可自定义的结构, 不管控",
                    "type": "object",
                    "additionalProperties": true
                },
                "is_valid": {
                    "description": "是否生效",
                    "type": "boolean"
                }
            }
        },
        "list_type.CreateListTypeReq": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "名单类型编码",
                    "type": "string",
                    "example": "telephone"
                },
                "description": {
                    "description": "描述",
                    "type": "string",
                    "example": "description"
                },
                "fields": {
                    "description": "这类名单的值被构建的字段",
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "telephone",
                        "id_card"
                    ]
                },
                "is_valid": {
                    "description": "是否生效",
                    "type": "boolean"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string",
                    "example": "anti-fraud"
                }
            }
        },
        "list_type.UpdateListTypeReq": {
            "type": "object",
            "required": [
                "description",
                "is_valid"
            ],
            "properties": {
                "description": {
                    "description": "描述",
                    "type": "string",
                    "example": "description"
                },
                "is_valid": {
                    "description": "是否生效",
                    "type": "boolean"
                }
            }
        },
        "models.ListItem": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "名单类型编码",
                    "type": "string",
                    "example": "telephone__red"
                },
                "extra": {
                    "description": "可自定义的结构, 不管控",
                    "type": "object",
                    "additionalProperties": true
                },
                "id": {
                    "description": "主键id",
                    "type": "string",
                    "example": "60d2b17f70d9d2f0db53f866"
                },
                "is_valid": {
                    "description": "是否生效",
                    "type": "boolean"
                },
                "multi_value": {
                    "description": "多项值列表，如果名单是由多个字段构成的，可一一罗列出每个字段的值，如：[{\"key\":\"field1\",\"value\":\"value1\"},{\"key\":\"field2\",\"value\":\"value2\"}]",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MultiValueItem"
                    }
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string",
                    "example": "anti-fraud"
                },
                "value": {
                    "description": "名单项的值",
                    "type": "string",
                    "example": "13333333333"
                }
            }
        },
        "models.ListType": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "名单类型编码",
                    "type": "string",
                    "example": "telephone__red"
                },
                "description": {
                    "description": "描述",
                    "type": "string",
                    "example": "descriptions"
                },
                "fields": {
                    "description": "这类名单的值被构建的字段",
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "telephone",
                        "id_card"
                    ]
                },
                "id": {
                    "description": "主键id",
                    "type": "string",
                    "example": "60d2b17f70d9d2f0db53f866"
                },
                "is_valid": {
                    "description": "是否生效",
                    "type": "boolean"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string",
                    "example": "anti-fraud"
                }
            }
        },
        "models.MultiValueItem": {
            "type": "object",
            "properties": {
                "key": {
                    "description": "field name",
                    "type": "string",
                    "example": "telephone"
                },
                "value": {
                    "description": "value",
                    "type": "string",
                    "example": "13333333333"
                }
            }
        },
        "models.Namespace": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "命名空间code",
                    "type": "string",
                    "example": "anti-fraud"
                },
                "description": {
                    "description": "描述",
                    "type": "string",
                    "example": "anti fraud use"
                },
                "id": {
                    "description": "主键id",
                    "type": "string",
                    "example": "60d2b17f70d9d2f0db53f866"
                }
            }
        },
        "namespace.CreateNamespaceReq": {
            "type": "object",
            "required": [
                "code",
                "description"
            ],
            "properties": {
                "code": {
                    "description": "命名空间code",
                    "type": "string",
                    "example": "anti-fraud"
                },
                "description": {
                    "description": "描述",
                    "type": "string",
                    "example": "anti fraud use"
                }
            }
        },
        "namespace.UpdateNamespaceReq": {
            "type": "object",
            "required": [
                "description"
            ],
            "properties": {
                "description": {
                    "description": "描述",
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
	Description: "名单管理系统",
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
