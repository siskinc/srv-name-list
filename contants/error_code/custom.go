package error_code

// 通用错误
const (
	CustomForbidden                 = 400001
	CustomForbiddenParameterInvalid = 400002
)

// 未知错误
const (
	CustomForbiddenNotFoundNamespace = 401001
	CustomForbiddenNotFoundListType  = 401002
	CustomForbiddenNotFoundListItem  = 401003
)

// 冲突错误
const (
	CustomForbiddenConflictNamespace = 402001
	CustomForbiddenConflictListType  = 402002
	CustomForbiddenConflictListItem  = 402003
)
