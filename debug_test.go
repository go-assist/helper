package helper

import (
	"strings"
	"testing"
)

func TestGetFuncName(t *testing.T) {
	res1 := TDebug.GetFuncName(nil, false)
	res2 := TDebug.GetFuncName(nil, true)
	res3 := TDebug.GetFuncName(TArr.ArrayDiff, false) // ...ArrayDiff-fm
	res4 := TDebug.GetFuncName(TArr.ArrayDiff, true)  // ArrayDiff-fm

	if !strings.Contains(res1, "TestGetFuncName") || res2 != "TestGetFuncName" || !strings.Contains(res3, "ArrayDiff") || !strings.HasPrefix(res4, "ArrayDiff") {
		t.Error("GetFuncName unit test fail")
		return
	}
}

func BenchmarkGetFuncName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TDebug.GetFuncName(nil, true)
	}
}

func TestGetFuncLine(t *testing.T) {
	res := TDebug.GetFuncLine()
	if res <= 0 {
		t.Error("GetFuncLine unit test fail")
		return
	}
}

func BenchmarkGetFuncLine(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TDebug.GetFuncLine()
	}
}

func TestGetFuncFileDir(t *testing.T) {
	res1 := TDebug.GetFuncFile()
	res2 := TDebug.GetFuncDir()
	if res1 == "" {
		t.Error("GetFuncFile unit test fail")
		return
	} else if res2 == "" {
		t.Error("GetFuncDir unit test fail")
		return
	}
}

func BenchmarkGetFuncFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TDebug.GetFuncFile()
	}
}

func BenchmarkGetFuncDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TDebug.GetFuncDir()
	}
}

func TestDumpStacks(t *testing.T) {
	TDebug.DumpStacks()
}

func BenchmarkDumpStacks(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 100; i++ {
		TDebug.DumpStacks()
	}
}

func TestHasMethod(t *testing.T) {
	var test = &TOs

	chk1 := TDebug.HasMethod(test, "IsLinux")
	chk2 := TDebug.HasMethod(test, "Hello")
	if !chk1 || chk2 {
		t.Error("HasMethod unit test fail")
		return
	}
}

func BenchmarkHasMethod(b *testing.B) {
	b.ResetTimer()
	var test = &TOs
	for i := 0; i < b.N; i++ {
		TDebug.HasMethod(test, "IsLinux")
	}
}

func TestGetFuncPackage(t *testing.T) {
	res1 := TDebug.GetFuncPackage()
	res2 := TDebug.GetFuncPackage(TDebug.GetFuncFile())
	res3 := TDebug.GetFuncPackage("test")

	if res1 != "kgo" || res1 != res2 || res3 != "" {
		t.Error("GetFuncPackage unit test fail")
		return
	}
}

func TestGetMethod(t *testing.T) {
	var test = &TOs

	fun1 := TDebug.GetMethod(test, "GoMemory")
	fun2 := TDebug.GetMethod(test, "Hello")

	if fun1 == nil || fun2 != nil {
		t.Error("GetMethod unit test fail")
		return
	}
}

func BenchmarkGetMethod(b *testing.B) {
	b.ResetTimer()
	var test = &TOs
	for i := 0; i < b.N; i++ {
		TDebug.GetMethod(test, "GoMemory")
	}
}

func TestCallMethod(t *testing.T) {
	var test = &TOs
	//无参数调用
	res1, err1 := TDebug.CallMethod(test, "GoMemory")
	if res1 == nil || err1 != nil {
		t.Error("CallMethod unit test fail")
		return
	}

	//调用不存在的方法
	res2, err2 := TDebug.CallMethod(test, "Hello")
	if res2 != nil || err2 == nil {
		t.Error("CallMethod unit test fail")
		return
	}

	//有参数调用
	var conv = &TConv
	res3, err3 := TDebug.CallMethod(conv, "BaseConvert", "123456", 10, 16)
	//结果 [1e240 <nil>]
	if len(res3) != 2 || res3[0] != "1e240" || res3[1] != nil || err3 != nil {
		t.Error("CallMethod unit test fail")
		return
	}
}

func BenchmarkCallMethod(b *testing.B) {
	b.ResetTimer()
	var test = &TOs
	for i := 0; i < b.N; i++ {
		TDebug.GetMethod(test, "GoMemory")
	}
}

func TestValidFunc(t *testing.T) {
	var err error
	var conv = &TConv
	method := TDebug.GetMethod(conv, "BaseConvert")

	//不存在的方法
	_, _, err = TDebug.ValidFunc("test", "echo")
	if err == nil {
		t.Error("ValidFunc unit test fail")
		return
	}

	//参数数量不足
	_, _, err = TDebug.ValidFunc(method, "12345")
	if err == nil {
		t.Error("ValidFunc unit test fail")
		return
	}

	//参数类型错误
	_, _, err = TDebug.ValidFunc(method, 0, "12345", "10", 16)
	if err == nil {
		t.Error("ValidFunc unit test fail")
		return
	}
}

func TestCallFunc(t *testing.T) {
	var err error
	var conv = &TConv
	method := TDebug.GetMethod(conv, "BaseConvert")

	_, err = TDebug.CallFunc(method, 0, "12345", "10", 16)
	if err != nil {
		println(err.Error())
	}
}

