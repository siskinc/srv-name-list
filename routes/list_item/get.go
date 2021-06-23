package list_item

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	listItemService "github.com/siskinc/srv-name-list/service/list_item"
)


// QueryListItem godoc
// @TAGS 名单项
// @Summary 名单项查找功能
// @Description 名单项查找功能, 通过code, is_valid, 分页
// @Accept json
// @Produce json
// @Param message query listItemService.ListItemQueryInfo false "a1111111111"
// @Success 200 {object} httpx.JSONResultPaged.{data=[]models.ListItem} "正常回包, 回复查询成功的名单类型数据"
// @Router /item [get]
func QueryListItem(c *gin.Context) {
	req := &listItemService.ListItemQueryInfo{}
	err := c.Bind(req)
	if err != nil {
		logrus.Errorf("cannot format data to listItemService.ListItemQueryInfo")
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}
	listItemServiceObj := listItemService.NewListItemService()
	listItemList, total, err := listItemServiceObj.Query(req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetRespJSONPaged(c, listItemList, "", total)
}
