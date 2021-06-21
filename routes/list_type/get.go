package list_type

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/httpx"
	"github.com/siskinc/srv-name-list/models"
	"net/http"
)

// QueryListType godoc
// @TAGS 名单类型
// @Summary 名单类型查找功能
// @Description 名单类型查找功能, 通过code, is_valid, 分页
// @Accept json
// @Produce json
// @Param is_valid query boolean false "是否生效"
// @Param code query string false "名单类型编码"
// @Param page_index query int false "页码" minimum(1) default(1)
// @Param page_size query int false "分页大小" minimum(10) default(10)
// @Success 200 {object} httpx.JSONResultPaged.{data=[]models.ListType} "正常回包, 回复查询成功的名单类型数据"
// @Router /type [get]
func QueryListType(c *gin.Context) {
	c.JSON(http.StatusOK, httpx.JSONResultPaged{Data: []models.ListType{}})
}
