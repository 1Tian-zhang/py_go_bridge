package main

import (
	_ "py_go_bridge_example2/internal/demo"  // 导入以触发init注册
	_ "py_go_bridge_example2/internal/excel" // 导入以触发init注册
)

func main() {}
