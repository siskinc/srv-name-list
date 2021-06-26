package list_item_hit

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	listItemService "github.com/siskinc/srv-name-list/service/list_item"
)

// ItemHitPre godoc
// @TAGS 名单项命中
// @Summary 名单项预命中
// @Description 名单项预命中, 指定某个名单项, 能够判断数据是否能够命中该名单项, 如果不能, 输出不能命中的原因
// @Accept json
// @Produce json
// @Param message body listItemService.ItemHitPreReq true "命中信息"
// @Success 200 {object} httpx.JSONResult.{data=listItemService.ItemHitPreResp} "正常回包, 回复创建成功的名单类型数据"
// @Router /item-hit/pre [post]
func ItemHitPre(c *gin.Context) {
	req := &listItemService.ItemHitPreReq{}
	err := c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to listItemService.ItemHitPreReq")
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}

	service := listItemService.NewService()
	resp, err := service.ItemHitPre(req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetRespJSON(c, resp, "")
}