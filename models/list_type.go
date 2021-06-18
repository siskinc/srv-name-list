package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ListType 名单类型
type ListType struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`                  // 主键id
	Code        string             `json:"code" bson:"code"`               // 名单类型编码
	Fields      []string           `json:"fields" bson:"fields"`           // 这类名单的值被构建的字段
	IsValid     bool               `json:"is_valid" bson:"is_valid"`       // 是否生效
	Description string             `json:"description" bson:"description"` // 描述
}
