package list_type

import (
	"github.com/gin-gonic/gin"
	"github.com/siskinc/srv-name-list/contants/errs"
	"github.com/siskinc/srv-name-list/internal/httpx"
	"github.com/siskinc/srv-name-list/models"
	listTypeRepo "github.com/siskinc/srv-name-list/repository/list_type"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type CreateListTypeReq struct {
	Namespace   string   `json:"namespace"`   // 命名空间
	Code        string   `json:"code"`        // 名单类型编码
	Fields      []string `json:"fields"`      // 这类名单的值被构建的字段
	IsValid     bool     `json:"is_valid"`    // 是否生效
	Description string   `json:"description"` // 描述
}

// CreateListType godoc
// @TAGS 名单类型
// @Summary 名单类型创建功能
// @Description 名单类型创建功能
// @Accept json
// @Produce json
// @Param message body CreateListTypeReq true "名单属性"
// @Success 200 {object} httpx.JSONResult.{data=models.ListType} "正常回包, 回复创建成功的名单类型数据"
// @Router /type [post]
func CreateListType(c *gin.Context) {
	req := &CreateListTypeReq{}
	err := c.Bind(req)
	if err != nil {
		httpx.SetRespErr(c, errs.CustomForbiddenParameterInvalidError)
		return
	}
	listType := &models.ListType{
		Id:          primitive.NewObjectID(),
		Namespace:   req.Namespace,
		Code:        req.Code,
		Fields:      req.Fields,
		IsValid:     req.IsValid,
		Description: req.Description,
	}
	repo := listTypeRepo.NewRepoListTypeMgo(listTypeRepo.NewCollection())
	err = repo.Create(listType)
	if err != nil {
		httpx.SetRespErr(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.JSONResult{Data: listType})
}
