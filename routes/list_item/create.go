package list_item

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/internal/httpx"
	"github.com/siskinc/srv-name-list/models"
)

type CreateListItemReq struct {
	Namespace   string   `json:"namespace"`   // 命名空间
	Code        string   `json:"code"`        // 名单类型编码
	Fields      []string `json:"fields"`      // 这类名单的值被构建的字段
	IsValid     bool     `json:"is_valid"`    // 是否生效
	Description string   `json:"description"` // 描述
}

// CreateListItem godoc
// @TAGS 名单项
// @Summary 名单项创建功能
// @Description 名单项创建功能
// @Accept json
// @Produce json
// @Param message body CreateListItemReq true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.ListItem} "正常回包, 回复创建成功的名单类型数据"
// @Router /item [post]
func CreateListItem(c *gin.Context) {
	httpx.SetRespJSON(c, models.ListItem{}, "")
}
