package list_item

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson"
)

type ListItemQueryInfo struct {
	// 命名空间
	Namespace   string  `form:"namespace" binding:"required"`
	// 是否生效
	IsValid     *bool   `form:"is_valid"`
	// Code
	Code        *string `form:"code"`
	// PageIndex
	PageIndex   int64   `form:"page_index"`
	// PageSize
	PageSize    int64   `form:"page_size"`
	// SortedField
	SortedField *string `form:"sorted_field"`
}

func (service *ListItemService) Query(info *ListItemQueryInfo) (results []*models.ListItem, total int64, err error) {
	if info == nil {
		err = errorx.NewError(error_code.CustomForbiddenConflictListType, fmt.Errorf("ListItemQueryInfo is nil"))
		return
	}
	filter := bson.D{
		bson.E{Key: "namespace", Value: info.Namespace},
	}
	if info.IsValid != nil {
		filter = append(filter, bson.E{Key: "is_valid", Value: *info.IsValid})
	}
	if info.Code != nil {
		filter = append(filter, bson.E{Key: "code", Value: *info.Code})
	}
	sortedField := ""
	if info.SortedField != nil {
		sortedField = *info.SortedField
	}
	if sortedField == "" {
		sortedField = "-_id"
	}
	results, total, err = service.listItemRepoObj.Query(filter, info.PageIndex, info.PageSize, sortedField)
	return
}
