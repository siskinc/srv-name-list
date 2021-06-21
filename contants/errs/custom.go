package errs

import (
	"fmt"
	"github.com/goools/tools/errorx"
	"github.com/siskinc/srv-name-list/contants/error_code"
)

var (
	CustomForbiddenParameterInvalidError = errorx.NewError(error_code.CustomForbiddenParameterInvalid, fmt.Errorf("参数错误"))
)
