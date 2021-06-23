package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ListType 名单类型
type ListType struct {
	Id          primitive.ObjectID `json:"id" bson:"_id" example:"60d2b17f70d9d2f0db53f866"`      // 主键id
	Namespace   string             `json:"namespace" bson:"namespace" example:"anti-fraud"`       // 命名空间
	Code        string             `json:"code" bson:"code" example:"telephone__red"`             // 名单类型编码
	Fields      []string           `json:"fields" bson:"fields" example:"telephone,id_card"`      // 这类名单的值被构建的字段
	IsValid     bool               `json:"is_valid" bson:"is_valid"`                              // 是否生效
	Description string             `json:"description" bson:"description" example:"descriptions"` // 描述
}
