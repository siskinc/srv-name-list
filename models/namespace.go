package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Namespace 命名空间
type Namespace struct {
	Id          primitive.ObjectID `json:"id" bson:"_id" example:"60d2b17f70d9d2f0db53f866"`        // 主键id
	Code        string             `json:"code" bson:"code" example:"anti-fraud"`                   // 命名空间code
	Description string             `json:"description" bson:"description" example:"anti fraud use"` // 描述
}
