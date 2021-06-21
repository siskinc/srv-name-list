package list_type

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	listTypeRepo "github.com/siskinc/srv-name-list/repository/list_type"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type UpdateListTypeReq struct {
	IsValid     bool   `json:"is_valid"`    // 是否生效
	Description string `json:"description"` // 描述
}

// UpdateListType godoc
// @TAGS 名单类型
// @Summary 名单类型修改功能
// @Description 名单类型修改功能, 通过code, is_valid, 分页
// @Accept json
// @Produce json
// @Param id path string true "名单类型id" minlength(1)
// @Param message body UpdateListTypeReq true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.ListType} "正常回包, 回复更新成功的名单类型数据"
// @Router /type/{id} [patch]
func UpdateListType(c *gin.Context) {
	listTypeId := c.Param("id")
	listTypeOid, err := primitive.ObjectIDFromHex(listTypeId)
	if err != nil {
		logrus.Errorf("list type id: %s, cannot convert to object id", listTypeId)
		httpx.SetRespErr(
			c,
			errorx.NewError(
				error_code.CustomForbiddenParameterInvalid,
				fmt.Errorf("%s不是一个合法的id", listTypeId),
			),
		)
		return
	}
	req := &UpdateListTypeReq{}
	err = c.Bind(req)
	if err != nil {
		logrus.Errorf("bind json to UpdateListTypeReq have an err: %v", err)
		httpx.SetRespErr(
			c,
			errorx.NewError(error_code.CustomForbiddenParameterInvalid, err),
		)
	}
	repo := listTypeRepo.NewRepoListTypeMgo()
	resp, err := repo.Update(listTypeOid, req.IsValid, req.Description)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.JSONResult{Data: resp})
}
