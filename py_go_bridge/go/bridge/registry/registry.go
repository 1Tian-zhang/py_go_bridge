package registry

import (
	"fmt"
	"reflect"
	"sync"
)

// FuncRegistry 函数注册表
type FuncRegistry struct {
	mu    sync.RWMutex
	funcs map[string]interface{}
}

var defaultRegistry = &FuncRegistry{
	funcs: make(map[string]interface{}),
}

// Register 注册函数
func Register(name string, fn interface{}) {
	defaultRegistry.mu.Lock()
	defer defaultRegistry.mu.Unlock()

	// 验证函数签名
	if err := validateFunc(fn); err != nil {
		panic(err)
	}

	defaultRegistry.funcs[name] = fn
}

// GetFunc 获取已注册的函数
func GetFunc(name string) interface{} {
	defaultRegistry.mu.RLock()
	defer defaultRegistry.mu.RUnlock()

	return defaultRegistry.funcs[name]
}

// GetAllFuncs 获取所有注册的函数
func GetAllFuncs() map[string]interface{} {
	defaultRegistry.mu.RLock()
	defer defaultRegistry.mu.RUnlock()

	funcs := make(map[string]interface{})
	for k, v := range defaultRegistry.funcs {
		funcs[k] = v
	}
	return funcs
}

// validateFunc 验证函数签名
func validateFunc(fn interface{}) error {
	// 获取函数类型
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		return fmt.Errorf("expected function, got %v", fnType.Kind())
	}

	// 检查返回值
	if fnType.NumOut() != 2 {
		return fmt.Errorf("function must return (interface{}, error)")
	}

	// 检查第二个返回值是否为error
	if !fnType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		return fmt.Errorf("second return value must be error")
	}

	// 检查第一个返回值是否为interface{}
	if fnType.Out(0).Kind() != reflect.Interface {
		return fmt.Errorf("first return value must be interface{}")
	}

	return nil
}
