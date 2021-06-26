package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	namespaceService "github.com/siskinc/srv-name-list/service/namespace"
)

// QueryNamespace godoc
// @TAGS 命名空间
// @Summary 命名空间查找功能
// @Description 命名空间查找功能
// @Accept json
// @Produce json
// @Param message query namespaceService.QueryNamespaceReq false "a"
// @Success 200 {object} httpx.JSONResultPaged.{data=[]models.Namespace} "正常回包, 回复查询成功的名单类型数据"
// @Router /namespace [get]
func QueryNamespace(c *gin.Context) {
	req := &namespaceService.QueryNamespaceReq{}
	err := c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to namespaceService.QueryNamespaceReq")
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}
	namespaceServiceObj := namespaceService.NewService()
	namespaces, total, err := namespaceServiceObj.Query(req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetRespJSONPaged(c, namespaces, "", total)
}