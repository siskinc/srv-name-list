package namespace

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateNamespaceReq struct {
	Oid         string `json:"-" swaggerignore:"true"` // 命名空间id
	Description string `json:"description" uri:"-" binding:"required"`             // 描述
}

func (service *Service) Update(req *UpdateNamespaceReq) (npObj *models.Namespace, err error) {
	if req == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("update namespace info is empty"))
		return
	}
	oid, err := primitive.ObjectIDFromHex(req.Oid)
	if err != nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("id %s is invalid", req.Oid))
		return
	}
	npObj, err = service.namespaceMongoRepo.Update(oid, req.Description)
	return
}
