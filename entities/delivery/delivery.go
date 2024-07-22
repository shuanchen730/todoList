package delivery

// HttpResponse 單位內部使用的通用response格式
type HttpResponse struct {
	Result    int         `json:"result"`
	Retrieve  interface{} `json:"ret,omitempty"`
	ErrorCode string      `json:"code,omitempty"`
	Message   string      `json:"msg,omitempty"`
}

type ErrorDetail struct {
	RedirectCode    string      `json:"code"`
	RedirectMessage string      `json:"message"`
	Details         interface{} `json:"detail"`
}