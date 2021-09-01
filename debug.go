package helper

import (
	"fmt"
	"go/parser"
	"go/token"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// GetFuncName 获取调用方法的名称;f为目标方法;onlyFun为true时仅返回方法,不包括包名.
func (td *TsDebug) GetFuncName(f interface{}, onlyFun ...bool) (funcName string) {
	var funcObj *runtime.Func
	if f == nil {
		// Skip this function, and fetch the PC and file for its parent
		pc, _, _, _ := runtime.Caller(1)
		// Retrieve a Function object this functions parent
		funcObj = runtime.FuncForPC(pc)
	} else {
		funcObj = runtime.FuncForPC(reflect.ValueOf(f).Pointer())
	}

	funcName = funcObj.Name()
	if len(onlyFun) > 0 && onlyFun[0] == true {
		// extract just the function name (and not the module path)
		funcName = strings.TrimPrefix(filepath.Ext(funcName), ".")
	}
	return
}

// GetFuncLine 获取调用方法的行号.
func (td *TsDebug) GetFuncLine() (line int) {
	// Skip this function, and fetch the PC and file for its parent
	_, _, line, _ = runtime.Caller(1)
	return
}

// GetFuncFile 获取调用方法的文件路径.
func (td *TsDebug) GetFuncFile() (file string) {
	_, file, _, _ = runtime.Caller(1)
	return
}

// GetFuncDir 获取调用方法的文件目录
func (td *TsDebug) GetFuncDir() string {
	return filepath.Dir(td.GetFuncFile())
}

// DumpStacks 打印堆栈信息.
func (td *TsDebug) DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}

// HasMethod 检查对象是否具有某方法.
func (td *TsDebug) HasMethod(t interface{}, method string) (exist bool) {
	_, exist = reflect.TypeOf(t).MethodByName(method)
	return
}

// GetFuncPackage 获取调用方法或源文件的包名.funcFile为源文件路径.
func (td *TsDebug) GetFuncPackage(funcFile ...string) string {
	var sourceFile string
	if len(funcFile) == 0 {
		sourceFile = td.GetFuncFile()
	} else {
		sourceFile = funcFile[0]
	}

	fSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fSet, sourceFile, nil, parser.PackageClauseOnly)
	if err != nil || astFile.Name == nil {
		return ""
	}

	return astFile.Name.Name
}

// GetMethod 获取对象中的方法.
// 注意:返回的该方法中的第一个参数是接收者.
// 所以,调用该方法时,必须将接收者作为第一个参数传递.
func (td *TsDebug) GetMethod(f interface{}, method string) interface{} {
	m := td.getMethod(f, method)
	if !m.IsValid() || m.IsNil() {
		return nil
	}
	return m.Interface()
}

// CallMethod 调用对象的方法.
// 若执行成功,则结果是该方法的返回结果;
// 否则返回(nil, error)
func (td *TsDebug) CallMethod(f interface{}, method string, args ...interface{}) ([]interface{}, error) {
	m := td.GetMethod(f, method)
	if m == nil {
		return nil, fmt.Errorf("don't have method: %s", method)
	}
	_args := make([]interface{}, len(args)+1)
	_args[0] = f
	copy(_args[1:], args)
	return td.CallFunc(m, _args...)
}

// CallFunc 动态调用函数.
func (td *TsDebug) CallFunc(f interface{}, args ...interface{}) (results []interface{}, err error) {
	vf, vars, _err := td.ValidFunc(f, args...)
	if _err != nil {
		return nil, _err
	}
	ret := vf.Call(vars)
	_len := len(ret)
	results = make([]interface{}, _len)
	for i := 0; i < _len; i++ {
		results[i] = ret[i].Interface()
	}
	return
}

// getMethod 获取对象的方法.
func (td *TsDebug) getMethod(f interface{}, method string) reflect.Value {
	m, exist := reflect.TypeOf(f).MethodByName(method)
	if !exist {
		return reflect.ValueOf(nil)
	}
	return m.Func
}

// ValidFunc 检查是否函数,并且参数个数、类型是否正确.
// 返回有效的函数、有效的参数.
func (td *TsDebug) ValidFunc(f interface{}, args ...interface{}) (vf reflect.Value, vars []reflect.Value, err error) {
	vf = reflect.ValueOf(f)
	if vf.Kind() != reflect.Func {
		return reflect.ValueOf(nil), nil, fmt.Errorf("[validFunc] %v is not the function", f)
	}

	tf := vf.Type()
	_len := len(args)
	if tf.NumIn() != _len {
		return reflect.ValueOf(nil), nil, fmt.Errorf("[validFunc] %d number of the argument is incorrect", _len)
	}

	vars = make([]reflect.Value, _len)
	for i := 0; i < _len; i++ {
		typ := tf.In(i).Kind()
		if (typ != reflect.Interface) && (typ != reflect.TypeOf(args[i]).Kind()) {
			return reflect.ValueOf(nil), nil, fmt.Errorf("[validFunc] %d-td argument`s type is incorrect", i+1)
		}
		vars[i] = reflect.ValueOf(args[i])
	}
	return vf, vars, nil
}
