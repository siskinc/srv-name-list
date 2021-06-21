package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
	"net/http"
)

func SetRespErr(c *gin.Context, err error) {
	var newErr errorx.IError
	var ok bool
	if newErr, ok = err.(errorx.IError); !ok {
		newErr = errorx.NewError(error_code.ServerError, err)
	}
	c.JSON(http.StatusOK, MakeResultWithError(newErr))
}

func SetRespJSON(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, JSONResult{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

func SetRespJSONPaged(c *gin.Context, data interface{}, message string, total int64) {
	c.JSON(http.StatusOK, JSONResultPaged{
		Code:    0,
		Message: message,
		Data:    data,
		Total:   total,
	})
}
