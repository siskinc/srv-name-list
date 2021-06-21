package list_item

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/internal/httpx"
)

// DeleteListItem godoc
// @TAGS 名单项
// @Summary 名单项删除功能
// @Description 名单项删除功能, 通过list_item_id删除
// @Accept json
// @Produce json
// @Param id path string true "名单项id" minlength(1)
// @Success 200 {object} httpx.JSONResult
// @Router /item/{id} [delete]
func DeleteListItem(c *gin.Context) {
	httpx.SetRespJSON(c, nil, "")
}