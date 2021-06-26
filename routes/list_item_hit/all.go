package list_item_hit

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	listItemService "github.com/siskinc/srv-name-list/service/list_item"
)

// ItemHitAll godoc
// @TAGS 名单项命中
// @Summary 名单项命中
// @Description 名单项命中, 返回命中的所有名单项
// @Accept json
// @Produce json
// @Param message body listItemService.ItemHitAllReq true "预命中信息"
// @Success 200 {object} httpx.JSONResult.{data=listItemService.ItemHitAllResp} "正常回包"
// @Router /item-hit/all [post]
func ItemHitAll(c *gin.Context) {
	req := &listItemService.ItemHitAllReq{}
	err := c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to listItemService.ItemHitPreAll")
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}
	service := listItemService.NewService()
	resp, err := service.ItemHitAll(req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetRespJSON(c, resp, "")
}
