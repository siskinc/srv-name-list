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
	Code        *string `form:"code"`         // Code 命名空间code
	PageIndex   int64   `form:"page_index"`   // PageIndex 页码
	PageSize    int64   `form:"page_size"`    // PageSize 分页数量
	SortedField string `form:"sorted_field"` // SortedField 排序字段
}

func (service *Service) Query(req *QueryNamespaceReq) (result []*models.Namespace, total int64, err error) {
	if req == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("query namespace info is empty"))
		return
	}
	filter := bson.D{}
	if req.Code != nil && *req.Code != "" {
		regexFilter := bson.E{
			Key:   "code",
			Value: bson.D{{"$regex", primitive.Regex{Pattern: fmt.Sprintf("%s", *req.Code), Options: "i"}}},
		}
		filter = append(filter, regexFilter)
	}
	result, total, err = service.namespaceMongoRepo.Query(filter, req.PageIndex, req.PageSize, req.SortedField)
	return
}
