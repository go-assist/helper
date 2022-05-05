package helper

import "testing"

func TestEqualFloat(t *testing.T) {
	var f1, f2 float64

	f1 = 1.2345678
	f2 = 1.2345679

	chk1 := TFloat.EqualFloat(f1, f2)
	chk2 := TFloat.EqualFloat(f1, f2, 6)

	if chk1 || !chk2 {
		t.Error("EqualFloat unit test fail")
		return
	}
}

func TestRound(t *testing.T) {
	got := float64(3)
	want1 := TFloat.Round(3.14)
	if got != want1 {
		t.Errorf("Round unit test, The values of %v is not %v\n", want1, got)
	}

	want2 := TFloat.Round(3.55)
	if got != want2 {
		t.Errorf("Round unit test, The values of %v is not %v\n", want2, got)
	}
}

func TestCeil(t *testing.T) {
	num1 := 0.3
	num2 := 1.6
	res1 := TFloat.Ceil(num1)
	res2 := TFloat.Ceil(num2)
	if int(res1) != 1 || int(res2) != 2 {
		t.Error("Ceil unit test fail")
		return
	}
}

func TestFloor(t *testing.T) {
	num1 := 0.3
	num2 := 0.6
	res1 := TFloat.Floor(num1)
	res2 := TFloat.Floor(num2)
	if int(res1) != 0 || int(res2) != 0 {
		t.Error("Floor unit test fail")
		return
	}
}

func TestMaxFloat64(t *testing.T) {
	nums := []float64{-4, 0, 3, 9}
	res := TFloat.MaxFloat64(nums...)
	if int(res) != 9 {
		t.Error("MaxFloat64 unit test fail")
		return
	}

	res = TFloat.MaxFloat64(-1)
	if int(res) != -1 {
		t.Error("MaxFloat64 unit test fail")
		return
	}
}

func TestMaxMinFloat64(t *testing.T) {
	nums := []float64{-4, 0, 3, 9}
	res := TFloat.MinFloat64(nums...)
	if int(res) != -4 {
		t.Error("MinFloat64 unit test fail")
		return
	}

	res = TFloat.MinFloat64(-1)
	if int(res) != -1 {
		t.Error("MinFloat64 unit test fail")
		return
	}
}

func TestIsNegative(t *testing.T) {
	for _, test := range exampleFloatIsNegative {
		actual := TValidate.IsNegative(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNegative(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsPositive(t *testing.T) {
	for _, test := range exampleFloatIsPositive {
		actual := TValidate.IsPositive(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsPositive(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsNonNegative(t *testing.T) {
	for _, test := range exampleFloatIsNonNegative {
		actual := TValidate.IsNonNegative(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNonNegative(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsNonPositive(t *testing.T) {
	for _, test := range exampleFloatIsNonPositive {
		actual := TValidate.IsNonPositive(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNonPositive(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsWhole(t *testing.T) {
	for _, test := range exampleFloatIsWhole {
		actual := TValidate.IsWhole(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsWhole(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestInRangeFloat32(t *testing.T) {
	for _, test := range exampleFloatTests {
		actual := TValidate.IsRangeFloat32(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeFloat32(%v, %v, %v) to be %v, got %v", test.param, test.left, test.right, test.expected, actual)
		}
	}
}

func TestInRangeFloat64(t *testing.T) {
	for _, test := range exampleFloatTests {
		actual := TValidate.IsRangeFloat64(float64(test.param), float64(test.left), float64(test.right))
		if actual != test.expected {
			t.Errorf("Expected InRangeFloat64(%v, %v, %v) to be %v, got %v", test.param, test.left, test.right, test.expected, actual)
		}
	}
}

func TestAverageFloat64(t *testing.T) {
	var res1, res2, res3 float64
	res1 = TValidate.AverageFloat64()
	res2 = TValidate.AverageFloat64(1)
	res3 = TValidate.AverageFloat64(1, 2, 3, 4, 5, 6)

	if res1 != 0 || int(res2) != 1 || TInt.NumberFormat(res3, 2, ".", "") != "3.50" {
		t.Error("AverageFloat64 unit test fail")
		return
	}
}

func TestSumFloat64(t *testing.T) {
	sum := TFloat.SumFloat64(0.0, 1.1, -2.2, 3.30, 5.55)
	if TInt.NumberFormat(sum, 2, ".", "") != "7.75" {
		t.Error("SumFloat64 unit test fail")
		return
	}
}
