package demo

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

//export Echo
func Echo(input C.cstring) unsafe.Pointer {
	return bridge.WrapFunc(func() (interface{}, error) {
		str := C.GoString(input)
		fmt.Println("input:", str)
		// fmt.Println("asdasdaslkdjaksdjlajdslajlskdjkla\n\n\n\n\n")
		return map[string]interface{}{
			"data": str,
		}, nil
	})
}
