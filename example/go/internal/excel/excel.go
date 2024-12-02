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
	"unsafe"

	"py_go_bridge/bridge"
)

//export ExportExcel
func ExportExcel(channelCode C.cstring, host C.cstring, port C.cint, user C.cstring, password C.cstring, database C.cstring, workerCount C.cint) unsafe.Pointer {
	return bridge.WrapFunc(func() (interface{}, error) {
		// 转换参数
		config := DBConfig{
			Host:     C.GoString(host),
			Port:     int(port),
			User:     C.GoString(user),
			Password: C.GoString(password),
			Database: C.GoString(database),
		}

		fmt.Println(config)
		// 调用导出函数
		result, err := ExportItemsTask(C.GoString(channelCode), config, int(workerCount))
		if err != nil {
			return nil, fmt.Errorf("export failed: %v", err)
		}

		return result, nil
	})
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func ExportItemsTask(channelCode string, config DBConfig, workerCount int) (map[string]interface{}, error) {
	// 这里简化实现，返回模拟数据
	return map[string]interface{}{
		"status":  "success",
		"message": "Export completed",
		"files":   []string{"file1.xlsx", "file2.xlsx"},
	}, nil
}
