package namespace

import (
	"fmt"

	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QueryNamespaceReq struct {
	ID          *string `form:"id"`           // ID 命名空间ID（精确查询）
	Code        *string `form:"code"`         // Code 命名空间code（模糊查询）
	Description *string `form:"description"`  // Description 命名空间的描述（模糊查询）
	PageIndex   int64   `form:"page_index"`   // PageIndex 页码
	PageSize    int64   `form:"page_size"`    // PageSize 分页数量
	SortedField string  `form:"sorted_field"` // SortedField 排序字段
}

func (service *Service) Query(req *QueryNamespaceReq) (result []*models.Namespace, total int64, err error) {
	if req == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("query namespace info is empty"))
		return
	}
	filter := bson.D{}
	if req.ID != nil && *req.ID != "" {
		var oid primitive.ObjectID
		oid, err = primitive.ObjectIDFromHex(*req.ID)
		if err != nil {
			err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("id is invalid"))
			return
		}
		filter = append(filter, bson.E{
			Key:   "_id",
			Value: oid,
		})
	}
	if req.Code != nil && *req.Code != "" {
		regexFilter := bson.E{
			Key:   "code",
			Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: *req.Code, Options: "i"}}},
		}
		filter = append(filter, regexFilter)
	}
	if req.Description != nil && *req.Description != "" {
		regexFilter := bson.E{
			Key:   "description",
			Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: *req.Description, Options: "i"}}},
		}
		filter = append(filter, regexFilter)
	}
	result, total, err = service.namespaceMongoRepo.Query(filter, req.PageIndex, req.PageSize, req.SortedField)
	return
}
