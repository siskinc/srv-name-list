package list_type

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	listTypeService "github.com/siskinc/srv-name-list/service/list_type"
)

// DeleteListType godoc
// @TAGS 名单类型
// @Summary 名单类型删除功能
// @Description 名单类型删除功能, 通过list_type_id删除
// @Accept json
// @Produce json
// @Param id path string true "名单类型id" minlength(1)
// @Success 200 {object} httpx.JSONResult "正常回包, 回复删除成功的名单类型数据"
// @Router /type/{id} [delete]
func DeleteListType(c *gin.Context) {
	req := &listTypeService.DeleteListTypeReq{}
	err := c.ShouldBindUri(req)
	if err != nil {
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("params is invalid")))
		return
	}

	service := listTypeService.NewService()
	err = service.Delete(req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}

	httpx.SetRespJSON(c, nil, "")
}
