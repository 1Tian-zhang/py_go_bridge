package bridge

/*
#include <stdlib.h>
#include <stdint.h>

typedef char* cstring;
typedef int32_t cint;
*/
import "C"
import (
	"encoding/json"
	"time"
	"unsafe"
)

// Response 统一的返回结构
type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Timing  float64     `json:"timing"`
}

// WrapFunc 包装Go函数
func WrapFunc(fn func() (interface{}, error)) unsafe.Pointer {
	start := time.Now()

	resp := &Response{
		Status:  "success",
		Code:    200,
		Message: "ok",
	}

	// 执行函数
	if result, err := fn(); err != nil {
		resp.Status = "error"
		resp.Code = 500
		resp.Message = err.Error()
	} else {
		resp.Data = result
	}

	// 设置执行时间(秒)
	resp.Timing = time.Since(start).Seconds()

	// 转换为JSON
	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		errResp := &Response{
			Status:  "error",
			Code:    500,
			Message: err.Error(),
		}
		jsonBytes, _ = json.Marshal(errResp)
	}

	// 返回C字符串
	return unsafe.Pointer(C.CString(string(jsonBytes)))
}

//export FreeString
func FreeString(str *C.char) {
	C.free(unsafe.Pointer(str))
}
