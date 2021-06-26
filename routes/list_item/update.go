package list_item

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	listItemService "github.com/siskinc/srv-name-list/service/list_item"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateListItemReq struct {
	IsValid bool                   `json:"is_valid"` // 是否生效
	Extra   map[string]interface{} `json:"extra"`    // 可自定义的结构, 不管控
}

// UpdateListItem godoc
// @TAGS 名单项
// @Summary 名单项修改功能
// @Description 名单项修改功能, 通过code, is_valid, 分页
// @Accept json
// @Produce json
// @Param id path string true "名单项id" minlength(1)
// @Param message body listItemService.UpdateListItemInfo true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.ListItem} "正常回包, 回复更新成功的名单类型数据"
// @Router /item/{id} [patch]
func UpdateListItem(c *gin.Context) {
	listItemId := c.Param("id")
	listItemOid, err := primitive.ObjectIDFromHex(listItemId)
	if err != nil {
		logrus.Errorf("list item id: %s, cannot convert to object id", listItemId)
		httpx.SetRespErr(
			c,
			errorx.NewError(
				error_code.CustomForbiddenParameterInvalid,
				fmt.Errorf("%s不是一个合法的id", listItemId),
			),
		)
		return
	}
	req := &listItemService.UpdateListItemInfo{}
	err = c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to listItemService.UpdateListItemInfo")
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}
	listItemServiceObj := listItemService.NewService()
	listItem, err := listItemServiceObj.Update(listItemOid, req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetRespJSON(c, listItem, "")
}
