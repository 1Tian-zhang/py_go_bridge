package excel

/*
#include <stdlib.h>
#include <stdint.h>

typedef char* cstring;
typedef int32_t cint;
*/
import "C"
import (
	"fmt"
	"py_go_bridge/bridge"
	"unsafe"
)

//export ExportExcel
func ExportExcel(channelCode C.cstring, host C.cstring, port C.cint, user C.cstring, password C.cstring, database C.cstring, workerCount C.cint) unsafe.Pointer {
	return bridge.WrapFunc(func() (interface{}, error) {
		// 转换参数
		config := export.DBConfig{
			Host:     C.GoString(host),
			Port:     int(port),
			User:     C.GoString(user),
			Password: C.GoString(password),
			Database: C.GoString(database),
		}

		// 调用原始函数
		result := export.ExportItemsTask(C.GoString(channelCode), config, int(workerCount))

		// 这里可以处理错误,如果ExportItemsTask返回的result中包含错误信息
		if result["status"] == "error" {
			return nil, fmt.Errorf(result["message"].(string))
		}

		return result, nil
	})
}
