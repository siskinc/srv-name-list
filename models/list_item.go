package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MultiValueItem struct {
	Key   string `json:"key" bson:"key" example:"telephone"`       // field name
	Value string `json:"value" bson:"value" example:"13333333333"` // value
}

// ListItem 名单项
type ListItem struct {
	Id         primitive.ObjectID     `json:"id" bson:"_id" example:"60d2b17f70d9d2f0db53f866"` // 主键id
	Namespace  string                 `json:"namespace" bson:"namespace" example:"anti-fraud"`  // 命名空间
	Code       string                 `json:"code" bson:"code" example:"telephone__red"`        // 名单类型编码
	Value      string                 `json:"value" bson:"value" example:"13333333333"`         // 名单项的值
	MultiValue []MultiValueItem       `json:"multi_value" bson:"multi_value"`                   // 多项值列表，如果名单是由多个字段构成的，可一一罗列出每个字段的值，如：[{"key":"field1","value":"value1"},{"key":"field2","value":"value2"}]
	IsValid    bool                   `json:"is_valid" bson:"is_valid"`                         // 是否生效
	Extra      map[string]interface{} `json:"extra" bson:"extra"`                               // 可自定义的结构, 不管控
}
