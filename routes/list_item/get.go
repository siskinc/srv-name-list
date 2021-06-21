package list_item

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/internal/httpx"
	"github.com/siskinc/srv-name-list/models"
)

type QueryListItemReq struct {
	Namespace   *string `form:"namespace"`
	IsValid     *bool   `form:"is_valid"`
	Code        *string `form:"code"`
	SortedField *string `form:"sorted_field"`
	PageIndex   int64   `form:"page_index"`
	PageSize    int64   `form:"page_size"`
}

// QueryListItem godoc
// @TAGS 名单项
// @Summary 名单项查找功能
// @Description 名单项查找功能, 通过code, is_valid, 分页
// @Accept json
// @Produce json
// @Param is_valid query boolean false "是否生效"
// @Param code query string false "名单类型编码" minlength(1)
// @Param page_index query int false "页码" minimum(1) default(1)
// @Param page_size query int false "分页大小" minimum(10) default(10)
// @Param sorted_field query string false "排序方式" minlength(1)
// @Param namespace query string false "命名空间" minlength(1)
// @Success 200 {object} httpx.JSONResultPaged.{data=[]models.ListItem} "正常回包, 回复查询成功的名单类型数据"
// @Router /item [get]
func QueryListItem(c *gin.Context) {
	httpx.SetRespJSON(c, []models.ListItem{}, "")
}
