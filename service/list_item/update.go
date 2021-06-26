package list_item

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/mongox"
	"github.com/siskinc/srv-name-list/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateListItemInfo struct {
	IsValid *bool                   `json:"is_valid"` // 是否生效
	Extra   *map[string]interface{} `json:"extra"`    // 可自定义的结构, 不管控
}

func (service *Service) makeUpdater(info *UpdateListItemInfo) bson.M {
	value := bson.M{}
	if info.IsValid != nil {
		value["is_valid"] = *info.IsValid
	}
	if info.Extra != nil {
		value["extra"] = *info.Extra
	}
	if len(value) == 0 {
		value = nil
	}
	return bson.M{
		"$set": value,
	}
}

func (service *Service) Update(oid primitive.ObjectID, info *UpdateListItemInfo) (listItem *models.ListItem, err error) {
	if mongox.EmptyOid(oid) {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("oid: %s is a empty oid", oid.Hex()))
		return
	}
	updater := service.makeUpdater(info)
	if updater["$set"] == nil {
		return
	}
	listItem, err = service.listItemRepoObj.UpdateById(oid, updater)
	if err != nil {
		logrus.Errorf("update by id have an err: %v", err)
		return
	}
	return
}
