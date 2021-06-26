package list_item

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/mongox"
	"github.com/siskinc/srv-name-list/models"
	listTypeService "github.com/siskinc/srv-name-list/service/list_type"
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type ItemHitPreReq struct {
	ListItemId string `json:"list_item_id" binding:"required" minlength:"1"` // 名单项ID
	Data       string `json:"data" binding:"required" minlength:"2"`         // 数据
}

type ItemHitPreResp struct {
	Hit      bool             `json:"hit"`
	Resource string           `json:"resource"`
	ListItem *models.ListItem `json:"list_item"`
}

func (service *Service) ItemHitPre(req *ItemHitPreReq) (resp *ItemHitPreResp, err error) {
	if req == nil {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("list item hit req is nil"))
		return
	}
	var (
		oid      primitive.ObjectID
		listItem *models.ListItem
		listType *models.ListType
	)

	resp = &ItemHitPreResp{}
	oid, err = primitive.ObjectIDFromHex(req.ListItemId)
	if mongox.EmptyOid(oid) {
		err = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("list item id is empty"))
		return
	}

	listItem, err = service.listItemRepoObj.FindById(oid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = errorx.NewError(error_code.CustomForbiddenNotFoundListItem, fmt.Errorf("not found list item %s", req.ListItemId))
		}
		return
	}

	resp.ListItem = listItem

	namespace, code := listItem.Namespace, listItem.Code
	listTypeServiceObj := listTypeService.NewService()
	listType, err = listTypeServiceObj.FindOne(namespace, code)
	if err != nil {
		return
	}

	if listType == nil {
		resp.Resource = fmt.Sprintf("not found list type, namespace: %s, code: %s", namespace, code)
		return
	}

	if !listType.IsValid {
		resp.Resource = "list_type's is_valid is closed"
		return
	}

	if !listItem.IsValid {
		resp.Resource = "is_valid is closed"
		return
	}

	results := gjson.GetMany(req.Data, listType.Fields...)
	values := make([]string, len(results))
	var notFoundFieldList []string
	for i := range results {
		if !results[i].Exists() {
			notFoundFieldList = append(notFoundFieldList, listType.Fields[i])
		}
		values[i] = results[i].Str
	}

	if len(notFoundFieldList) > 0 {
		resp.Resource = fmt.Sprintf("cannot find fields: %s", strings.Join(notFoundFieldList, ","))
		return
	}

	value := strings.Join(values, ",")

	if listItem.Value != value {
		resp.Resource = fmt.Sprintf("value not equal, list item value is %s, data value is %s",
			listItem.Value, value)
		return
	}

	resp.Hit = true
	return
}
