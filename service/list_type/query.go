package list_type

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QueryListTypeReq struct {
	Namespace   *string `form:"namespace"`
	IsValid     *bool   `form:"is_valid"`
	Code        *string `form:"code"`
	SortedField *string `form:"sorted_field"`
	PageIndex   int64   `form:"page_index"`
	PageSize    int64   `form:"page_size"`
}

func (service *Service) QueryListType(req *QueryListTypeReq) (result []*models.ListType, total int64, err error) {
	if req == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("query list type req is nil"))
		return
	}

	filter := bson.D{}
	if req.IsValid != nil {
		filter = append(filter, bson.E{Key: "is_valid", Value: *req.IsValid})
	}
	if req.Code != nil && *req.Code != "" {
		regexFilter := bson.E{
			Key:   "code",
			Value: bson.D{{"$regex", primitive.Regex{Pattern: fmt.Sprintf("%s", *req.Code), Options: "i"}}},
		}
		filter = append(filter, regexFilter)
	}
	if req.Namespace != nil && *req.Namespace != "" {
		filter = append(filter, bson.E{Key: "namespace", Value: *req.Namespace})
	}
	sortedField := "-_id"
	if req.SortedField != nil && *req.SortedField != "" {
		sortedField = *req.SortedField
	}
	result, total, err = service.listTypeRepoObj.Query(filter, req.PageIndex, req.PageSize, sortedField)
	return
}

func (service *Service) QueryAll(filter bson.D) (listTypeList []*models.ListType, err error) {
	listTypeList, _, err = service.listTypeRepoObj.Query(filter, 0, 0, "")
	return
}
