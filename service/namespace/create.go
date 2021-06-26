package namespace

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateNamespaceReq struct {
	Code        string `json:"code" example:"anti-fraud" binding:"required"`            // 命名空间code
	Description string `json:"description" example:"anti fraud use" binding:"required"` // 描述
}

func (service *Service) Create(info *CreateNamespaceReq) (npObj *models.Namespace, err error) {
	if info == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("create namespace info is nil"))
		return
	}
	var exist bool
	exist, err = service.CheckExist(info.Code)
	if err != nil {
		return
	}
	if exist {
		err = errorx.NewError(error_code.CustomForbiddenConflictNamespace, fmt.Errorf("namespace code is exist"))
		return
	}
	npObj = &models.Namespace{
		Id:          primitive.NewObjectID(),
		Code:        info.Code,
		Description: info.Description,
	}
	err = service.namespaceMongoRepo.Create(npObj)
	return
}
