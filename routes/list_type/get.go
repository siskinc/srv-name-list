package list_type

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"github.com/siskinc/srv-name-list/internal/httpx"
	listTypeRepo "github.com/siskinc/srv-name-list/repository/list_type"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type QueryListTypeReq struct {
	Namespace   *string `form:"namespace"`
	IsValid     *bool   `form:"is_valid"`
	Code        *string `form:"code"`
	SortedField *string `form:"sorted_field"`
	PageIndex   int64   `form:"page_index"`
	PageSize    int64   `form:"page_size"`
}

// QueryListType godoc
// @TAGS 名单类型
// @Summary 名单类型查找功能
// @Description 名单类型查找功能, 通过code, is_valid, 分页
// @Accept json
// @Produce json
// @Param is_valid query boolean false "是否生效"
// @Param code query string false "名单类型编码" minlength(1)
// @Param page_index query int false "页码" minimum(1) default(1)
// @Param page_size query int false "分页大小" minimum(10) default(10)
// @Param sorted_field query string false "排序方式" minlength(1)
// @Param namespace query string false "命名空间" minlength(1)
// @Success 200 {object} httpx.JSONResultPaged.{data=[]models.ListType} "正常回包, 回复查询成功的名单类型数据"
// @Router /type [get]
func QueryListType(c *gin.Context) {
	req := &QueryListTypeReq{}
	err := c.Bind(req)
	if err != nil {
		httpx.SetRespErr(c, errorx.NewError(error_code.CustomForbiddenParameterInvalid, err))
		return
	}
	repo := listTypeRepo.NewRepoListTypeMgo()
	filter := bson.D{}
	if req.IsValid != nil {
		filter = append(filter, bson.E{Key: "is_valid", Value: *req.IsValid})
	}
	if req.Code != nil {
		regexFilter := bson.E{
			Key:   "code",
			Value: bson.D{{"$regex", primitive.Regex{Pattern: fmt.Sprintf("%s", *req.Code), Options: "i"}}},
		}
		filter = append(filter, regexFilter)
	}
	if req.Namespace != nil {
		filter = append(filter, bson.E{Key: "namespace", Value: *req.Namespace})
	}
	sortedField := "-_id"
	if req.SortedField != nil && *req.SortedField != "" {
		sortedField = *req.SortedField
	}
	result, total, err := repo.Query(filter, req.PageIndex, req.PageSize, sortedField)
	if err != nil {
		logrus.Errorf("query list type have an err: %v, req: %+v, sortedField: %s, filter: %+v",
			err, req, sortedField, filter)
		httpx.SetRespErr(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.JSONResultPaged{Data: result, Total: total})
}
