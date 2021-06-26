package list_type

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/models"
	namespaceService "github.com/siskinc/srv-name-list/service/namespace"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateListTypeReq struct {
	Namespace   string   `json:"namespace" example:"anti-fraud"`     // 命名空间
	Code        string   `json:"code" example:"telephone"`           // 名单类型编码
	Fields      []string `json:"fields" example:"telephone,id_card"` // 这类名单的值被构建的字段
	IsValid     bool     `json:"is_valid"`                           // 是否生效
	Description string   `json:"description" example:"description"`  // 描述
}

func (service *Service) Create(req *CreateListTypeReq) (listType *models.ListType, err error) {
	if req == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("create list type req is nil"))
		return
	}
	namespaceServiceObj := namespaceService.NewService()
	var exist bool
	exist, err = namespaceServiceObj.CheckExist(req.Namespace)
	if err != nil {
		return
	}
	if !exist {
		err = errorx.NewError(error_code.CustomForbiddenNotFoundNamespace, fmt.Errorf("not found namespace %s", req.Namespace))
		return
	}
	listType = &models.ListType{
		Id:          primitive.NewObjectID(),
		Namespace:   req.Namespace,
		Code:        req.Code,
		Fields:      req.Fields,
		IsValid:     req.IsValid,
		Description: req.Description,
	}
	err = service.listTypeRepoObj.Create(listType)
	return
}
