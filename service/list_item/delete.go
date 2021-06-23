package list_item

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/mongox"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *ListItemService) Delete(oid primitive.ObjectID) (err error) {
	if mongox.EmptyOid(oid) {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("oid: %s is a empty oid", oid.Hex()))
		return
	}
	return service.listItemRepoObj.DeleteById(oid)
}
