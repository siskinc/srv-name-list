package list_type

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/httpx"
	listTypeRepo "github.com/siskinc/srv-name-list/repository/list_type"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteListType godoc
// @TAGS 名单类型
// @Summary 名单类型删除功能
// @Description 名单类型删除功能, 通过list_type_id删除
// @Accept json
// @Produce json
// @Param id path string true "名单类型id" minlength(1)
// @Success 200 {object} httpx.JSONResult "正常回包, 回复删除成功的名单类型数据"
// @Router /type/{id} [delete]
func DeleteListType(c *gin.Context) {
	listTypeId := c.Param("id")
	listTypeOid, err := primitive.ObjectIDFromHex(listTypeId)
	if err != nil {
		logrus.Errorf("list type id: %s, cannot convert to object id", listTypeId)
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("%s不是一个合法的id", listTypeId)))
		return
	}
	repo := listTypeRepo.NewRepoListTypeMgo(listTypeRepo.NewCollection())
	err = repo.Delete(listTypeOid)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	httpx.SetRespJSON(c, nil, fmt.Sprintf("删除id为%s的数据成功", listTypeId))
}
