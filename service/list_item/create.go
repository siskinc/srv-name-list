package list_item

import (
	"fmt"

	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/models"
	listTypeService "github.com/siskinc/srv-name-list/service/list_type"
	namespaceService "github.com/siskinc/srv-name-list/service/namespace"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateListItemInfo struct {
	Namespace string                 `json:"namespace" binding:"required"`
	Code      string                 `json:"code" binding:"required"`
	IsValid   bool                   `json:"is_valid"`
	Values    []string               `json:"values" binding:"required"` // 与list type中的fields一一对应的
	Extra     map[string]interface{} `json:"extra"`
}

func (service *Service) Create(listItemCreateInfo *CreateListItemInfo) (*models.ListItem, error) {
	if listItemCreateInfo == nil {
		return nil, fmt.Errorf("list item create info is nil")
	}
	if listItemCreateInfo.Extra == nil {
		listItemCreateInfo.Extra = make(map[string]interface{})
	}

	namespaceServiceObj := namespaceService.NewService()
	exist, err := namespaceServiceObj.CheckExist(listItemCreateInfo.Namespace)
	if err != nil {
		return nil, fmt.Errorf("find %s namespace have an err: %v", listItemCreateInfo.Namespace, err)
	}

	if !exist {
		return nil, errorx.NewError(error_code.CustomForbiddenNotFoundNamespace, fmt.Errorf("not found %s namespace", listItemCreateInfo.Namespace))
	}

	namespace := listItemCreateInfo.Namespace
	code := listItemCreateInfo.Code

	listTypeServiceObj := listTypeService.NewService()
	req := &listTypeService.QueryListTypeReq{
		Namespace: &namespace,
		Code:      &code,
	}
	listTypeList, _, err := listTypeServiceObj.QueryListType(req)
	if err != nil {
		return nil, err
	}
	if len(listTypeList) <= 0 {
		err = errorx.NewError(error_code.CustomForbiddenNotFoundListType, fmt.Errorf("not found list type %s", code))
		return nil, err
	}
	listType := listTypeList[0]
	fields := listType.Fields
	values := listItemCreateInfo.Values
	if len(fields) != len(values) {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("values len not eqaul list type fields len, fields: %+v", listType.Fields))
		return nil, err
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
