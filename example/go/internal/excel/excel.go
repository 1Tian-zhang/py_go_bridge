package excel

/*
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

typedef const char* c_string;
typedef int32_t c_int;
*/
import "C"
import (
	"fmt"
	"py_go_bridge/bridge"
	"unsafe"
)

//export ExportExcel
func ExportExcel(channelCode *C.char, host *C.char, port C.int32_t, user *C.char, password *C.char, database *C.char, workerCount C.int32_t) unsafe.Pointer {
	return bridge.WrapFunc(func() (interface{}, error) {
		// 转换参数
		config := DBConfig{
			Host:     C.GoString(channelCode),
			Port:     int(port),
			User:     C.GoString(user),
			Password: C.GoString(password),
			Database: C.GoString(database),
		}

		// 直接调用ExportItemsTask
		result := ExportItemsTask(C.GoString(channelCode), config, int(workerCount))

		// 处理错误
		if result["status"] == "error" {
			return nil, fmt.Errorf(result["message"].(string))
		}

		return result, nil
	})
}
