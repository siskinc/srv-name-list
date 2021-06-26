package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	namespaceService "github.com/siskinc/srv-name-list/service/namespace"
)

// DeleteNamespace godoc
// @TAGS 命名空间
// @Summary 命名空间删除功能
// @Description 命名空间删除功能, 通过namespace_id删除
// @Accept json
// @Produce json
// @Param id path string true "命名空间id" minlength(1)
// @Success 200 {object} httpx.JSONResult
// @Router /namespace/{id} [delete]
func DeleteNamespace(c *gin.Context) {
	req := &namespaceService.DeleteNamespaceReq{}
	err := c.ShouldBindUri(req)
	if err != nil {
		logrus.Errorf("cannot format data to namespaceService.DeleteNamespaceReq")
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}
	namespaceServiceObj := namespaceService.NewService()
	err = namespaceServiceObj.Delete(req)
	if err != nil {
		logrus.Errorf("delete namespace have an err: %v, req: %+v", err, req)
		httpx.SetRespErr(c, err)
		return
	}
	return
}