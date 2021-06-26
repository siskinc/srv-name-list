package list_type

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/mongox"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteListTypeReq struct {
	OId string `uri:"id"`
}

func (service *Service) Delete(req *DeleteListTypeReq) (err error) {
	if req == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("delete list type req is nil"))
		return
	}
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(req.OId)
	if err != nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("id is valid"))
		return
	}
	if mongox.EmptyOid(oid) {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("id is zore"))
		return
	}
	err = service.listTypeRepoObj.Delete(oid)
	return
}
