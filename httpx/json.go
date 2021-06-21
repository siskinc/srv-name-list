package httpx

type JSONResult struct {
	Code    int         `json:"code"`    // 回包code，表明是否正确，在code == 0时，表明服务正常
	Message string      `json:"message"` // 回报message，在code != 0时，展示给前端
	Data    interface{} `json:"data"`    // 数据
}

type JSONResultPaged struct {
	Code    int         `json:"code"`    // 回包code，表明是否正确，在code == 0时，表明服务正常
	Message string      `json:"message"` // 回报message，在code != 0时，展示给前端
	Data    interface{} `json:"data"`    // 数据
	Total   int         `json:"total"`   // 总数量
}
