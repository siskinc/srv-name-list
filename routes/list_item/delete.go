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

// DeleteListItem godoc
// @TAGS 名单项
// @Summary 名单项删除功能
// @Description 名单项删除功能, 通过list_item_id删除
// @Accept json
// @Produce json
// @Param id path string true "名单项id" minlength(1)
// @Success 200 {object} httpx.JSONResult
// @Router /item/{id} [delete]
func DeleteListItem(c *gin.Context) {
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
	listItemServiceObj := listItemService.NewListItemService()
	err = listItemServiceObj.Delete(listItemOid)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetRespJSON(c, nil, "")
}
