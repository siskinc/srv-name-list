package list_type

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	listTypeService "github.com/siskinc/srv-name-list/service/list_type"
	"net/http"
)

type QueryListTypeReq struct {
	Namespace   *string `form:"namespace"`    // 命名空间
	IsValid     *bool   `form:"is_valid"`     // 是否生效
	Code        *string `form:"code"`         // code
	SortedField *string `form:"sorted_field"` // 排序字段, 格式如: -_id, _id, code, -code; 带"-"前缀表明倒序排列
	PageIndex   int64   `form:"page_index"`   // 页码
	PageSize    int64   `form:"page_size"`    // 每一页数量
}

// QueryListType godoc
// @TAGS 名单类型
// @Summary 名单类型查找功能
// @Description 名单类型查找功能, 通过code, is_valid, 分页
// @Accept json
// @Produce json
// @Param message query listTypeService.QueryListTypeReq false "是否生效"
// @Success 200 {object} httpx.JSONResultPaged.{data=[]models.ListType} "正常回包, 回复查询成功的名单类型数据"
// @Router /type [get]
func QueryListType(c *gin.Context) {
	req := &listTypeService.QueryListTypeReq{}
	err := c.ShouldBind(req)
	if err != nil {
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}
	service := listTypeService.NewService()
	result, total, err := service.QueryListType(req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.JSONResultPaged{Data: result, Total: total})
}
