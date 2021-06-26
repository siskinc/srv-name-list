package list_type

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/contants/errs"
	"github.com/siskinc/srv-name-list/internal/httpx"
	listTypeService "github.com/siskinc/srv-name-list/service/list_type"
	"net/http"
)



// CreateListType godoc
// @TAGS 名单类型
// @Summary 名单类型创建功能
// @Description 名单类型创建功能
// @Accept json
// @Produce json
// @Param message body listTypeService.CreateListTypeReq true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.ListType} "正常回包, 回复创建成功的名单类型数据"
// @Router /type [post]
func CreateListType(c *gin.Context) {
	req := &listTypeService.CreateListTypeReq{}
	err := c.ShouldBind(req)
	if err != nil {
		httpx.SetRespErr(c, errs.CustomForbiddenParameterInvalidError)
		return
	}
	listTypeServiceObj := listTypeService.NewService()
	listType, err := listTypeServiceObj.Create(req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.JSONResult{Data: listType})
}
