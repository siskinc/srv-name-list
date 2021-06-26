package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	namespaceService "github.com/siskinc/srv-name-list/service/namespace"
)

// CreateNamespace godoc
// @TAGS 命名空间
// @Summary 命名空间创建功能
// @Description 命名空间创建功能
// @Accept json
// @Produce json
// @Param message body namespaceService.CreateNamespaceReq true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.Namespace} "正常回包, 回复创建成功的名单类型数据"
// @Router /namespace [post]
func CreateNamespace(c *gin.Context) {
	req := &namespaceService.CreateNamespaceReq{}
	var err error
	err = c.ShouldBind(req)
	if err != nil {
		logrus.Errorf("cannot format data to namespaceService.CreateNamespaceReq, err: %v", err)
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}
	namespaceServiceObj := namespaceService.NewService()
	npObj, err := namespaceServiceObj.Create(req)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetRespJSON(c, npObj, "")
}

