package types

// Response 统一的返回结构
type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Timing  int64       `json:"timing"`
}

// NewResponse 创建新的响应
func NewResponse() *Response {
	return &Response{
		Status:  "success",
		Code:    200,
		Message: "ok",
	}
}

// SetError 设置错误信息
func (r *Response) SetError(err error) {
	r.Status = "error"
	r.Code = 500
	r.Message = err.Error()
}
