package list_type

import (
	"fmt"

	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/mongox"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateListTypeReq struct {
	OId         string `json:"-" swaggerignore:"true"`
	IsValid     bool   `json:"is_valid"`                                             // 是否生效
	Description string `json:"description" example:"description" binding:"required"` // 描述
}

func (service *Service) Update(req *UpdateListTypeReq) (listType *models.ListType, err error) {
	if req == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("update list type req is nil"))
		return
	}
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(req.OId)
	if err != nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("id %s is invalid", req.OId))
		return
	}
	if mongox.EmptyOid(oid) {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("id %s is empty", req.OId))
		return
	}
	listType, err = service.listTypeRepoObj.Update(oid, req.IsValid, req.Description)
	return
}
