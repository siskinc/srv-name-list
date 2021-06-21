package list_type

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/httpx"
	"github.com/siskinc/srv-name-list/models"
	"net/http"
)

type CreateListTypeReq struct {
	Code        string   `json:"code"`        // 名单类型编码
	Fields      []string `json:"fields"`      // 这类名单的值被构建的字段
	IsValid     bool     `json:"is_valid"`    // 是否生效
	Description string   `json:"description"` // 描述
}

// CreateListType godoc
// @TAGS 名单类型
// @Summary 名单类型创建功能
// @Description 名单类型查找功能, 通过code, is_valid, 分页
// @Accept json
// @Produce json
// @Param message body CreateListTypeReq true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.ListType} "正常回包, 回复创建成功的名单类型数据"
// @Router /type [post]
func CreateListType(c *gin.Context) {
	c.JSON(http.StatusOK, httpx.JSONResult{Data: models.ListType{}})
}
