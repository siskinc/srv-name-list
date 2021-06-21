package list_item

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/internal/httpx"
	"github.com/siskinc/srv-name-list/models"
)

type UpdateListItemReq struct {
	IsValid bool                   `json:"is_valid"` // 是否生效
	Extra   map[string]interface{} `json:"extra"`    // 可自定义的结构, 不管控
}

// UpdateListItem godoc
// @TAGS 名单项
// @Summary 名单项修改功能
// @Description 名单项修改功能, 通过code, is_valid, 分页
// @Accept json
// @Produce json
// @Param id path string true "名单项id" minlength(1)
// @Param message body UpdateListItemReq true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.ListItem} "正常回包, 回复更新成功的名单类型数据"
// @Router /item/{id} [patch]
func UpdateListItem(c *gin.Context) {
	httpx.SetRespJSON(c, models.ListItem{}, "")
}
