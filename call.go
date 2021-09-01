package helper

import (
	"reflect"
)

type Routers struct {
}

// ControllerMapsType 方法集合
type ControllerMapsType map[string]reflect.Value

// DynamicCallFunction 动态调用方法
func (tc *TsCallFunc) DynamicCallFunction(funcName string, args ... interface{}) (result []reflect.Value, err error) {
	var router Routers
	ControllerMap := make(ControllerMapsType, 0)
	rf := reflect.ValueOf(&router)
	rft := rf.Type()
	funcNum := rf.NumMethod()
	for i := 0; i < funcNum; i ++ {
		mName := rft.Method(i).Name
		ControllerMap[mName] = rf.Method(i)
	}
	parameter := make([]reflect.Value, len(args))
	for k, arg := range args {
		parameter[k] = reflect.ValueOf(arg)
	}
	result = ControllerMap[funcName].Call(parameter)
	return
}

// Example 示例
func (ts *Routers) Example(args ...interface{}) error {
	return nil
}