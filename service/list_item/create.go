package list_item

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/models"
	listTypeRepo "github.com/siskinc/srv-name-list/repository/list_type"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListItemCreateInfo struct {
	Namespace string                 `json:"namespace" binding:"required"`
	Code      string                 `json:"code" binding:"required"`
	IsValid   bool                   `json:"is_valid" binding:"required"`
	Values    []string               `json:"values" binding:"required"` // 与list type中的fields一一对应的
	Extra     map[string]interface{} `json:"extra"`
}


func (service *ListItemService) Create(listItemCreateInfo *ListItemCreateInfo) (*models.ListItem, error) {
	if listItemCreateInfo == nil {
		return nil, fmt.Errorf("list item create info is nil")
	}

	makeListTypeQueryFilter := func(namespace, code string) bson.D {
		return bson.D{
			{
				Key:   "namespace",
				Value: namespace,
			},
			{
				Key:   "code",
				Value: code,
			},
		}
	}

	namespace := listItemCreateInfo.Namespace
	code := listItemCreateInfo.Code

	listTypeRepoObj := listTypeRepo.NewRepoListTypeMgo()
	listTypeQuery := makeListTypeQueryFilter(namespace, code)
	listTypeList, _, err := listTypeRepoObj.Query(listTypeQuery, 1, 10, "")
	if err != nil {
		return nil, err
	}
	if len(listTypeList) <= 0 {
		err = errorx.NewError(error_code.CustomForbiddenNotFoundListType, fmt.Errorf("未找到对应的名单类型，请添加"))
		return nil, err
	}
	listType := listTypeList[0]
	fields := listType.Fields
	values := listItemCreateInfo.Values
	if len(fields) != len(values) {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("值的数量和查找到的名单类型字段的数量不对等"))
	}
	multiValue := service.makeMultiValue(fields, values)
	value := service.makeValue(multiValue)
	listItem := &models.ListItem{
		Id:         primitive.NewObjectID(),
		Namespace:  namespace,
		Code:       code,
		Value:      value,
		MultiValue: multiValue,
		IsValid:    listItemCreateInfo.IsValid,
		Extra:      listItemCreateInfo.Extra,
	}
	err = service.listItemRepoObj.Create(listItem)
	if err != nil {
		return nil, err
	}
	return listItem, nil
}
