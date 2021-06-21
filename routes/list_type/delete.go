package list_type

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/httpx"
	"github.com/siskinc/srv-name-list/models"
	"net/http"
)

// DeleteListType godoc
// @TAGS 名单类型
// @Summary 名单类型删除功能
// @Description 名单类型删除功能, 通过list_type_id删除
// @Accept json
// @Produce json
// @Param id path string true "名单类型id" minlength(1)
// @Success 200 {object} httpx.JSONResult.{data=models.ListType} "正常回包, 回复删除成功的名单类型数据"
// @Router /type/{id} [delete]
func DeleteListType(c *gin.Context)  {
	c.JSON(http.StatusOK, httpx.JSONResult{Data: models.ListType{}})
}
