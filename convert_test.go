package helper

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestInt2Str(t *testing.T) {
	tim := TConv.Int2Str(TTime.Time())
	if fmt.Sprint(reflect.TypeOf(tim)) != "string" {
		t.Error("Int2Str unit test fail")
		return
	}

	//非整型的转为空
	res := TConv.Int2Str(1.23)
	if res != "" {
		t.Error("Int2Str unit test fail")
		return
	}
}

func BenchmarkInt2Str(b *testing.B) {
	b.ResetTimer()
	tim := TTime.Time()
	for i := 0; i < b.N; i++ {
		TConv.Int2Str(tim)
	}
}

func TestIntFloat2Str(t *testing.T) {
	fl := float32(1234.567890)
	f2 := 1234.567890
	res1 := TConv.Float2Str(fl, 4)
	res2 := TConv.Float2Str(f2, 8)
	if fmt.Sprint(reflect.TypeOf(res1)) != fmt.Sprint(reflect.TypeOf(res2)) {
		t.Error("Int2Str unit test fail")
		return
	}

	//非浮点的转为空
	res := TConv.Float2Str(123, 2)
	if res != "" {
		t.Error("Float2Str unit test fail")
		return
	}
}

func Benchmark32Float2Str(b *testing.B) {
	b.ResetTimer()
	fl := float32(1234.567890)
	for i := 0; i < b.N; i++ {
		TConv.Float2Str(fl, 4)
	}
}

func Benchmark64Float2Str(b *testing.B) {
	b.ResetTimer()
	f2 := 1234.567890
	for i := 0; i < b.N; i++ {
		TConv.Float2Str(f2, 8)
	}
}

func TestBool2Str(t *testing.T) {
	res1 := TConv.Bool2Str(true)
	res2 := TConv.Bool2Str(false)
	if res1 != "true" {
		t.Error("Bool2Str unit test fail")
		return
	} else if res2 != "false" {
		t.Error("Bool2Str unit test fail")
		return
	}
}

func BenchmarkBool2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Bool2Str(true)
	}
}

func TestBool2Int(t *testing.T) {
	res1 := TConv.Bool2Int(true)
	res2 := TConv.Bool2Int(false)
	if res1 != 1 || res2 != 0 {
		t.Error("Bool2Int unit test fail")
		return
	}
}

func BenchmarkBool2Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Bool2Int(true)
	}
}

func TestStr2IntStrict(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	res := TConv.Str2IntStrict("abc123", 8, true)
	if fmt.Sprint(reflect.TypeOf(res)) != "int8" {
		t.Error("Str2IntStrict unit test fail")
		return
	}
}

func TestStr2Int(t *testing.T) {
	res := TConv.Str2Int("123")
	if fmt.Sprint(reflect.TypeOf(res)) != "int" {
		t.Error("Str2Int unit test fail")
		return
	}

	var tests = []struct {
		param    string
		expected int
	}{
		{"", 0},
		{"123", 123},
		{"123.45", 0},
		{"True", 1},
		{"false", 0},
	}

	for _, test := range tests {
		actual := TConv.Str2Int(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToBool(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkStr2Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Int("-123")
	}
}

func TestStr2Int8(t *testing.T) {
	tim := TConv.Int2Str(TTime.MicroTime())
	res := TConv.Str2Int8(tim)
	if res > 127 {
		t.Error("Str2Int8 unit test fail")
		return
	}
}

func BenchmarkStr2Int8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Int8("128")
	}
}

func TestStr2Int16(t *testing.T) {
	tim := TConv.Int2Str(TTime.MicroTime())
	res := TConv.Str2Int16(tim)
	if res > 32767 {
		t.Error("Str2Int16 unit test fail")
		return
	}
}

func BenchmarkStr2Int16(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Int16("32768")
	}
}

func TestStr2Int32(t *testing.T) {
	tim := TConv.Int2Str(TTime.MicroTime())
	res := TConv.Str2Int32(tim)
	if res > 2147483647 {
		t.Error("Str2Int32 unit test fail")
		return
	}
}

func BenchmarkStr2Int32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Int32("2147483647")
	}
}

func TestStr2Int64(t *testing.T) {
	tim := TConv.Int2Str(TTime.MicroTime())
	res := TConv.Str2Int64(tim)
	if res > Int64Max {
		t.Error("Str2Int64 unit test fail")
		return
	}
}

func BenchmarkStr2Int64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Int64("9223372036854775808")
	}
}

func TestStr2UintStrict(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()
	res := TConv.Str2UintStrict("abc123", 8, true)
	if fmt.Sprint(reflect.TypeOf(res)) != "uint8" {
		t.Error("Str2UintStrict unit test fail")
		return
	}
}

func TestStr2Uint(t *testing.T) {
	res := TConv.Str2Uint("-123")
	if fmt.Sprint(reflect.TypeOf(res)) != "uint" {
		t.Error("Str2Uint unit test fail")
		return
	}
}

func BenchmarkStr2Uint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Uint("123")
	}
}

func TestStr2Uint8(t *testing.T) {
	tim := TConv.Int2Str(TTime.MicroTime())
	res := TConv.Str2Uint8(tim)
	if res > 255 {
		t.Error("Str2Uint8 unit test fail")
		return
	}
}

func BenchmarkStr2Uint8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Uint8("256")
	}
}

func TestStr2Uint16(t *testing.T) {
	tim := TConv.Int2Str(TTime.MicroTime())
	res := TConv.Str2Uint16(tim)
	if res > 65535 {
		t.Error("Str2Uint16 unit test fail")
		return
	}
}

func BenchmarkStr2Uint16(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Uint16("65536")
	}
}

func TestStr2Uint32(t *testing.T) {
	tim := TConv.Int2Str(TTime.MicroTime())
	res := TConv.Str2Uint32(tim)
	if res > 4294967295 {
		t.Error("Str2Uint32 unit test fail")
		return
	}
}

func BenchmarkStr2Uint32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Uint32("4294967296")
	}
}

func TestStr2Uint64(t *testing.T) {
	tim := TConv.Int2Str(TTime.MicroTime())
	res := TConv.Str2Uint64(tim)
	if res > Uint64Max {
		t.Error("Str2Uint64 unit test fail")
		return
	}
}

func BenchmarkStr2Uint64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Uint64("9223372036854775808")
	}
}

func TestStr2FloatStrict(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	res := TConv.Str2FloatStrict("abc123", 32, true)
	if fmt.Sprint(reflect.TypeOf(res)) != "float32" {
		t.Error("Str2FloatStrict unit test fail")
		return
	}
}

func TestStr2Float32(t *testing.T) {
	res := TConv.Str2Float32("123.456")
	if fmt.Sprint(reflect.TypeOf(res)) != "float32" {
		t.Error("Str2Float32 unit test fail")
		return
	}
}

func BenchmarkStr2Float32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Float32("123.456")
	}
}

func TestStr2Float64(t *testing.T) {
	res := TConv.Str2Float64("123.456")
	if fmt.Sprint(reflect.TypeOf(res)) != "float64" {
		t.Error("Str2Float64 unit test fail")
		return
	}
}

func BenchmarkStr2Float64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Float64("123.456")
	}
}

func TestStr2Bool(t *testing.T) {
	res1 := TConv.Str2Bool("true")
	res2 := TConv.Str2Bool("True")
	res3 := TConv.Str2Bool("TRUE")
	res4 := TConv.Str2Bool("Hello")

	if !res1 || !res2 || !res3 {
		t.Error("Str2Bool unit test fail")
		return
	} else if res4 {
		t.Error("Str2Bool unit test fail")
		return
	}
}

func BenchmarkStr2Bool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Str2Bool("123.456")
	}
}

func TestStr2Bytes(t *testing.T) {
	str := `hello world!`
	res := TConv.Str2Bytes(str)
	if fmt.Sprint(reflect.TypeOf(res)) != "[]uint8" {
		t.Error("Str2Bytes unit test fail")
		return
	}
}

func BenchmarkStr2Bytes(b *testing.B) {
	b.ResetTimer()
	str := `hello world!
// Convert different types to byte slice using types and functions in unsafe and reflect package. 
// It has higher performance, but notice that it may be not safe when garbage collection happens.
// Use it when you need to temporary convert a long string to a byte slice and won't keep it for long time.
`
	for i := 0; i < b.N; i++ {
		TConv.Str2Bytes(str)
	}
}

func TestBytes2Str(t *testing.T) {
	sli := []byte("hello world!")
	res := TConv.Bytes2Str(sli)
	if fmt.Sprint(reflect.TypeOf(res)) != "string" {
		t.Error("Bytes2Str unit test fail")
		return
	}
}

func BenchmarkBytes2Str(b *testing.B) {
	b.ResetTimer()
	sli := []byte(`hello world!
// Convert different types to byte slice using types and functions in unsafe and reflect package. 
// It has higher performance, but notice that it may be not safe when garbage collection happens.
// Use it when you need to temporary convert a long string to a byte slice and won't keep it for long time.
`)
	for i := 0; i < b.N; i++ {
		TConv.Bytes2Str(sli)
	}
}

func TestDec2Bin(t *testing.T) {
	var num int64 = 8
	res := TConv.Dec2Bin(num)
	if res != "1000" {
		t.Error("Dec2Bin unit test fail")
		return
	}
}

func BenchmarkDec2Bin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Dec2Bin(10)
	}
}

func TestBin2Dec(t *testing.T) {
	res, err := TConv.Bin2Dec("1000")
	if err != nil || res != 8 {
		t.Error("Bin2Dec unit test fail")
		return
	}
	_, _ = TConv.Bin2Dec("hello")
}

func BenchmarkBin2Dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TConv.Bin2Dec("1000")
	}
}

func TestHex2Bin(t *testing.T) {
	_, err := TConv.Hex2Bin("123abaft")
	if err != nil {
		t.Error("Hex2Bin unit test fail")
		return
	}
	_, _ = TConv.Hex2Bin("hello")
}

func BenchmarkHex2Bin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TConv.Hex2Bin("123abaft")
	}
}

func TestBin2Hex(t *testing.T) {
	_, err := TConv.Bin2Hex("1001000111010101111111111")
	if err != nil {
		t.Error("Bin2Hex unit test fail")
		return
	}
	_, _ = TConv.Bin2Hex("hello")
}

func BenchmarkBin2Hex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TConv.Bin2Hex("1001000111010101111111111")
	}
}

func TestDec2Hex(t *testing.T) {
	res := TConv.Dec2Hex(1234567890)
	if res != "499602d2" {
		t.Error("Dec2Hex unit test fail")
		return
	}
}

func BenchmarkDec2Hex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Dec2Hex(1234567890)
	}
}

func TestHex2Dec(t *testing.T) {
	res1, err := TConv.Hex2Dec("123abf")
	res2, _ := TConv.Hex2Dec("0x123abf")
	if err != nil {
		t.Error("Hex2Dec unit test fail")
		return
	} else if res1 != res2 {
		t.Error("Hex2Dec unit test fail")
		return
	}
}

func BenchmarkHex2Dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TConv.Hex2Dec("123abf")
	}
}

func TestDec2Oct(t *testing.T) {
	res := TConv.Dec2Oct(123456789)
	if res != "726746425" {
		t.Error("Dec2Oct unit test fail")
		return
	}
}

func BenchmarkDec2Oct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Dec2Oct(123456789)
	}
}

func TestOct2Dec(t *testing.T) {
	res1, err := TConv.Oct2Dec("726746425")
	res2, _ := TConv.Oct2Dec("0726746425")
	if err != nil {
		t.Error("Oct2Dec unit test fail")
		return
	} else if res1 != res2 {
		t.Error("Oct2Dec unit test fail")
		return
	}
}

func BenchmarkOct2Dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TConv.Oct2Dec("726746425")
	}
}

func TestBaseConvert(t *testing.T) {
	_, err := TConv.BaseConvert("726746425", 10, 16)
	if err != nil {
		t.Error("BaseConvert unit test fail")
		return
	}
	_, _ = TConv.BaseConvert("hello", 10, 16)
}

func BenchmarkBaseConvert(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TConv.BaseConvert("726746425", 10, 16)
	}
}

func TestIp2Long(t *testing.T) {
	res := TConv.Ip2Long("127.0.0.1")
	if res == 0 {
		t.Error("Ip2Long unit test fail")
		return
	}
	TConv.Ip2Long("1")
}

func BenchmarkIp2Long(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Ip2Long("127.0.0.1")
	}
}

func TestLong2Ip(t *testing.T) {
	res := TConv.Long2Ip(2130706433)
	if res != "127.0.0.1" {
		t.Error("Long2Ip unit test fail")
		return
	}
}

func BenchmarkLong2Ip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Long2Ip(2130706433)
	}
}

func TestGettype(t *testing.T) {
	res1 := TConv.Gettype(1)
	res2 := TConv.Gettype("hello")
	res3 := TConv.Gettype(false)
	if res1 != "int" || res2 != "string" || res3 != "bool" {
		t.Error("Gettype unit test fail")
		return
	}
}

func BenchmarkGettype(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.Gettype("hello")
	}
}

func TestToStr(t *testing.T) {
	var fn CallBack
	mp := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	fnPtr := &fn

	var tests = []struct {
		param    interface{}
		expected string
	}{
		{-1, "-1"},
		{int8(0), "0"},
		{int16(1), "1"},
		{int32(2), "2"},
		{Int64Max, "9223372036854775807"},
		{uint(0), "0"},
		{uint8(0), "0"},
		{uint16(0), "0"},
		{uint32(0), "0"},
		{Uint64Max, "18446744073709551615"},
		{float32(math.Pi), "3.1415927"},
		{math.Pi, "3.141592653589793"},
		{[]byte{}, ""},
		{"1", "1"},
		{true, "true"},
		{false, "false"},
		{fn, "<nil>"},
		{nil, ""},
		{fnPtr, ""},
		{mp, `{"a":"aa","b":"bb"}`},
	}

	for _, test := range tests {
		actual := TConv.ToStr(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToStr(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkToStr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.ToStr(Uint64Max)
	}
}

func TestToInt(t *testing.T) {
	var fn CallBack
	var tests = []struct {
		param    interface{}
		expected int
	}{
		{-1, -1},
		{int8(0), 0},
		{int16(1), 1},
		{int32(2), 2},
		{int64(3), 3},
		{uint(0), 0},
		{uint8(0), 0},
		{uint16(0), 0},
		{uint32(0), 0},
		{uint64(0), 0},
		{float32(0), 0},
		{float64(0), 0},
		{[]byte{}, 0},
		{"1", 1},
		{"2.1", 0},
		{"TRUE", 1},
		{true, 1},
		{false, 0},
		{fn, 0},
	}

	for _, test := range tests {
		actual := TConv.ToInt(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToInt(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkToInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.ToInt("123")
	}
}

func TestToFloat(t *testing.T) {
	var fn CallBack
	var tests = []struct {
		param    interface{}
		expected float64
	}{
		{-1, -1.0},
		{int8(0), 0.0},
		{int16(1), 1.0},
		{int32(2), 2.0},
		{int64(3), 3.0},
		{uint(0), 0.0},
		{uint8(0), 0.0},
		{uint16(0), 0.0},
		{uint32(0), 0.0},
		{uint64(0), 0.0},
		{float32(0), 0.0},
		{float64(0), 0.0},
		{[]byte{}, 0.0},
		{"1", 1.0},
		{"2.1", 2.1},
		{"TRUE", 1.0},
		{true, 1.0},
		{false, 0},
		{fn, 0},
	}

	for _, test := range tests {
		actual := TConv.ToFloat(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToFloat(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkToFloat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.ToFloat("123")
	}
}

func TestFloat64ToByte(t *testing.T) {
	num := 12345.6
	res := TConv.Float64ToByte(num)
	if len(res) == 0 {
		t.Error("Float64ToByte unit test fail")
		return
	}
}

func BenchmarkFloat64ToByte(b *testing.B) {
	b.ResetTimer()
	num := 12345.6
	for i := 0; i < b.N; i++ {
		TConv.Float64ToByte(num)
	}
}

func TestByte2Float64(t *testing.T) {
	bs := []byte{205, 204, 204, 204, 204, 28, 200, 64}
	res := TConv.Byte2Float64(bs)
	if res != 12345.6 {
		t.Error("Byte2Float64 unit test fail")
		return
	}
}

func BenchmarkByte2Float64(b *testing.B) {
	b.ResetTimer()
	bs := []byte{205, 204, 204, 204, 204, 28, 200, 64}
	for i := 0; i < b.N; i++ {
		TConv.Byte2Float64(bs)
	}
}

func TestInt64ToByte(t *testing.T) {
	var num int64 = 12345
	res := TConv.Int64ToByte(num)
	if len(res) == 0 {
		t.Error("Int64ToByte unit test fail")
		return
	}
}

func BenchmarkInt64ToByte(b *testing.B) {
	b.ResetTimer()
	var num int64 = 12345
	for i := 0; i < b.N; i++ {
		TConv.Int64ToByte(num)
	}
}

func TestByte2Int64(t *testing.T) {
	bs := []byte{0, 0, 0, 0, 0, 0, 48, 57}
	res := TConv.Byte2Int64(bs)
	if res != 12345 {
		t.Error("Byte2Float64 unit test fail")
		return
	}
}

func BenchmarkByte2Int64(b *testing.B) {
	b.ResetTimer()
	bs := []byte{0, 0, 0, 0, 0, 0, 48, 57}
	for i := 0; i < b.N; i++ {
		TConv.Byte2Int64(bs)
	}
}

func TestByte2Hex(t *testing.T) {
	bs := []byte("hello")
	res := TConv.Byte2Hex(bs)
	if res != "68656c6c6f" {
		t.Error("Byte2Hex unit test fail")
		return
	}
}

func BenchmarkByte2Hex(b *testing.B) {
	b.ResetTimer()
	bs := []byte("hello")
	for i := 0; i < b.N; i++ {
		TConv.Byte2Hex(bs)
	}
}

func TestHex2Byte(t *testing.T) {
	str := "68656c6c6f"
	res := TConv.Hex2Byte(str)
	if string(res) != "hello" {
		t.Error("Hex2Byte unit test fail")
		return
	}
}

func BenchmarkHex2Byte(b *testing.B) {
	b.ResetTimer()
	str := "68656c6c6f"
	for i := 0; i < b.N; i++ {
		TConv.Hex2Byte(str)
	}
}

func TestGetPointerAddrInt(t *testing.T) {
	v1 := 1
	v2 := []byte("hello")

	res1 := TConv.GetPointerAddrInt(v1)
	res2 := TConv.GetPointerAddrInt(v2)
	if res1 <= 0 || res2 <= 0 {
		t.Error("GetPointerAddrInt unit test fail")
		return
	}
}

func BenchmarkGetPointerAddrInt(b *testing.B) {
	b.ResetTimer()
	v := []byte("hello")
	for i := 0; i < b.N; i++ {
		TConv.GetPointerAddrInt(v)
	}
}

func TestToBool(t *testing.T) {
	//并行测试
	t.Parallel()

	var fn CallBack

	var tests = []struct {
		param    interface{}
		expected bool
	}{
		{-1, false},
		{int8(0), false},
		{int16(1), true},
		{int32(2), true},
		{int64(3), true},
		{uint(0), false},
		{uint8(0), false},
		{uint16(0), false},
		{uint32(0), false},
		{uint64(0), false},
		{float32(0), false},
		{float64(0), false},
		{[]byte{}, false},
		{"1", true},
		{"2.1", false},
		{"TRUE", true},
		{false, false},
		{fn, false},
	}

	for _, test := range tests {
		actual := TConv.ToBool(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToBool(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkToBool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TConv.ToBool(1)
	}
}