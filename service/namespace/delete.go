package namespace

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteNamespaceReq struct {
	Oid string `uri:"id" binding:"required"`
}

func (service *Service) Delete(req *DeleteNamespaceReq) (err error) {
	if req == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("delete namespace params is empty"))
		return
	}
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(req.Oid)
	if err != nil {
		return
	}
	err = service.namespaceMongoRepo.Delete(oid)
	return
}
