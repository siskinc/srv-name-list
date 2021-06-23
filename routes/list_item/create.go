package list_item

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	listItemService "github.com/siskinc/srv-name-list/service/list_item"
)


// CreateListItem godoc
// @TAGS 名单项
// @Summary 名单项创建功能
// @Description 名单项创建功能
// @Accept json
// @Produce json
// @Param message body listItemService.ListItemCreateInfo true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.ListItem} "正常回包, 回复创建成功的名单类型数据"
// @Router /item [post]
func CreateListItem(c *gin.Context) {
	req := &listItemService.ListItemCreateInfo{}
	var err error
	err = c.Bind(req)
	if err != nil {
		logrus.Errorf("cannot format data to listItemService.ListItemCreateInfo")
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}
	listItemServiceObj := listItemService.NewListItemService()
	listItem, err := listItemServiceObj.Create(req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetRespJSON(c, listItem, "")
}
