package helper

import (
	"testing"
)

func TestDynamicCallFunction(t *testing.T) {
	fun, err := TCallFunc.DynamicCallFunction("Example")
	if err != nil {
		t.Errorf("The DynamicCallFunction errors of %v \n", err)
	}
	funRes := false
	funRes = fun[0].IsNil()
	if !funRes {
		t.Errorf("The DynamicCallFunction function errors of %v \n", fun[0])
	}
}
