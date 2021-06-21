package list_type

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/httpx"
	"github.com/siskinc/srv-name-list/models"
	"net/http"
)

type UpdateListTypeReq struct {
	IsValid     bool     `json:"is_valid"`    // 是否生效
	Description string   `json:"description"` // 描述
}

// UpdateListType godoc
// @TAGS 名单类型
// @Summary 名单类型创建功能
// @Description 名单类型查找功能, 通过code, is_valid, 分页
// @Accept json
// @Produce json
// @Param id path string true "名单类型id" minlength(1)
// @Param message body UpdateListTypeReq true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.ListType} "正常回包, 回复更新成功的名单类型数据"
// @Router /type/{id} [patch]
func UpdateListType(c *gin.Context)  {
	c.JSON(http.StatusOK, httpx.JSONResult{Data: models.ListType{}})
}
