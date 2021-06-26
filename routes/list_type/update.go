package list_type

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	listTypeService "github.com/siskinc/srv-name-list/service/list_type"
	"net/http"
)

// UpdateListType godoc
// @TAGS 名单类型
// @Summary 名单类型修改功能
// @Description 名单类型修改功能
// @Accept json
// @Produce json
// @Param id path string true "名单类型id" minlength(1)
// @Param message body listTypeService.UpdateListTypeReq true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.ListType} "正常回包, 回复更新成功的名单类型数据"
// @Router /type/{id} [patch]
func UpdateListType(c *gin.Context) {
	req := &listTypeService.UpdateListTypeReq{}
	err := c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("bind json to UpdateListTypeReq have an err: %v", err)
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}
	req.OId = c.Params.ByName("id")
	service := listTypeService.NewService()
	listType, err := service.Update(req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.JSONResult{Data: listType})
}
