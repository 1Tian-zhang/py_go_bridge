package wrapper

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

	"py_go_bridge/go/bridge/types"
)

// WrapFunc 包装Go函数以供导出
func WrapFunc(fn interface{}) unsafe.Pointer {
	start := time.Now()

	resp := types.NewResponse()

	// 执行函数
	if result, err := fn.(func() (interface{}, error))(); err != nil {
		resp.SetError(err)
	} else {
		resp.Data = result
	}

	// 设置执行时间
	resp.Timing = time.Since(start).Milliseconds()

	// 转换为JSON
	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		errResp := types.NewResponse()
		errResp.SetError(err)
		jsonBytes, _ = json.Marshal(errResp)
	}

	// 返回C字符串
	return unsafe.Pointer(C.CString(string(jsonBytes)))
}
