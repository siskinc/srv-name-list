package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	namespaceService "github.com/siskinc/srv-name-list/service/namespace"
)

// UpdateNamespace godoc
// @TAGS 命名空间
// @Summary 命名空间修改功能
// @Description 命名空间修改功能
// @Accept json
// @Produce json
// @Param id path string true "命名空间id" minlength(1)
// @Param message body namespaceService.UpdateNamespaceReq true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.Namespace} "正常回包, 回复更新成功的名单类型数据"
// @Router /namespace/{id} [patch]
func UpdateNamespace(c *gin.Context) {
	req := &namespaceService.UpdateNamespaceReq{}
	err := c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to namespaceService.UpdateListItemInfo, err: %v", err)
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}
	req.Oid = c.Params.ByName("id")
	namespaceServiceObj := namespaceService.NewService()
	npObj, err := namespaceServiceObj.Update(req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetRespJSON(c, npObj, "")
}
