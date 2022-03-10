package helper

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"reflect"
	"strconv"
	"unsafe"
)

// Int2Str 将整数转换为字符串
func (tc *TsConvert) Int2Str(val interface{}) string {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	// Type is not integers, return empty string
	default:
		return ""
	}
}

// Float2Str 将浮点数转换为字符串,decimal为小数位数.
func (tc *TsConvert) Float2Str(val interface{}, decimal int) string {
	switch val.(type) {
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'f', decimal, 32)
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', decimal, 64)
	// Type is not floats, return empty string
	default:
		return ""
	}
}

// Bool2Str 将布尔值转换为字符串.
func (tc *TsConvert) Bool2Str(val bool) (b string) {
	if val {
		b = "true"
		return
	}
	b = "false"
	return
}

// Bool2Int 将布尔值转换为整型.
func (tc *TsConvert) Bool2Int(val bool) int {
	if val {
		return 1
	}
	return 0
}

// Str2IntStrict 严格将字符串转换为有符号整型.
// bitSize为类型位数,strict为是否严格检查.
func (tc *TsConvert) Str2IntStrict(val string, bitSize int, strict bool) int64 {
	integer, err := strconv.ParseInt(val, 0, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return integer
}

// Str2Int 将字符串转换为int.其中"true", "TRUE", "True"为1.
func (tc *TsConvert) Str2Int(val string) (integer int) {
	if val == "true" || val == "TRUE" || val == "True" {
		integer = 1
		return
	}

	integer, _ = strconv.Atoi(val)
	return
}

// Str2Int8 将字符串转换为int8.
func (tc *TsConvert) Str2Int8(val string) int8 {
	return int8(tc.Str2IntStrict(val, 8, false))
}

// Str2Int16 将字符串转换为int16.
func (tc *TsConvert) Str2Int16(val string) int16 {
	return int16(tc.Str2IntStrict(val, 16, false))
}

// Str2Int32 将字符串转换为int32.
func (tc *TsConvert) Str2Int32(val string) int32 {
	return int32(tc.Str2IntStrict(val, 32, false))
}

// Str2Int64 将字符串转换为int64.
func (tc *TsConvert) Str2Int64(val string) int64 {
	return tc.Str2IntStrict(val, 64, false)
}

// Str2UintStrict 严格将字符串转换为无符号整型,bitSize为类型位数,strict为是否严格检查
func (tc *TsConvert) Str2UintStrict(val string, bitSize int, strict bool) uint64 {
	uInteger, err := strconv.ParseUint(val, 0, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return uInteger
}

// Str2Uint 将字符串转换为uint.
func (tc *TsConvert) Str2Uint(val string) uint {
	return uint(tc.Str2UintStrict(val, 0, false))
}

// Str2Uint8 将字符串转换为uint8.
func (tc *TsConvert) Str2Uint8(val string) uint8 {
	return uint8(tc.Str2UintStrict(val, 8, false))
}

// Str2Uint16 将字符串转换为uint16.
func (tc *TsConvert) Str2Uint16(val string) uint16 {
	return uint16(tc.Str2UintStrict(val, 16, false))
}

// Str2Uint32 将字符串转换为uint32.
func (tc *TsConvert) Str2Uint32(val string) uint32 {
	return uint32(tc.Str2UintStrict(val, 32, false))
}

// Str2Uint64 将字符串转换为uint64.
func (tc *TsConvert) Str2Uint64(val string) uint64 {
	return tc.Str2UintStrict(val, 64, false)
}

// Str2FloatStrict 严格将字符串转换为浮点型.
// bitSize为类型位数,strict为是否严格检查.
func (tc *TsConvert) Str2FloatStrict(val string, bitSize int, strict bool) float64 {
	parseFloat, err := strconv.ParseFloat(val, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return parseFloat
}

// Str2Float32 将字符串转换为float32.
func (tc *TsConvert) Str2Float32(val string) float32 {
	return float32(tc.Str2FloatStrict(val, 32, false))
}

// Str2Float64 将字符串转换为float64.其中"true", "TRUE", "True"为1.0 .
func (tc *TsConvert) Str2Float64(val string) (parseFloat float64) {
	if val == "true" || val == "TRUE" || val == "True" {
		parseFloat = 1.0
	} else {
		parseFloat = tc.Str2FloatStrict(val, 64, false)
	}

	return
}

// Str2Bool 将字符串转换为布尔值.
// 1, t, T, TRUE, true, True 等字符串为真.
// 0, f, F, FALSE, false, False 等字符串为假.
func (tc *TsConvert) Str2Bool(val string) (parseBool bool) {
	if val != "" {
		parseBool, _ = strconv.ParseBool(val)
	}
	return
}

// Str2Bytes 将字符串转换为字节切片.
// 该方法零拷贝,但不安全.它直接转换底层指针,两者指向的相同的内存,改一个另外一个也会变.
// 仅当临时需将长字符串转换且不长时间保存时可以使用.
// 转换之后若没做其他操作直接改变里面的字符,则程序会崩溃.
// 如 b:=String2bytes("xxx"); b[1]='d'; 程序将panic.
func (tc *TsConvert) Str2Bytes(val string) []byte {
	pSliceHeader := &reflect.SliceHeader{}
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&val))
	pSliceHeader.Data = strHeader.Data
	pSliceHeader.Len = strHeader.Len
	pSliceHeader.Cap = strHeader.Len
	return *(*[]byte)(unsafe.Pointer(pSliceHeader))
}

// Bytes2Str 将字节切片转换为字符串.
// 零拷贝,不安全.效率是string([]byte{})的百倍以上,且转换量越大效率优势越明显.
func (tc *TsConvert) Bytes2Str(val []byte) string {
	return *(*string)(unsafe.Pointer(&val))
}

// Dec2Bin 将十进制转换为二进制.
func (tc *TsConvert) Dec2Bin(number int64) string {
	return strconv.FormatInt(number, 2)
}

// Bin2Dec 将二进制转换为十进制.
func (tc *TsConvert) Bin2Dec(str string) (int64, error) {
	return strconv.ParseInt(str, 2, 0)
}

// Hex2Bin 将十六进制字符串转换为二进制字符串.
func (tc *TsConvert) Hex2Bin(data string) (s string, err error) {
	i, err := strconv.ParseInt(data, 16, 0)
	if err != nil {
		return
	}
	s = strconv.FormatInt(i, 2)
	return
}

// Bin2Hex 将二进制字符串转换为十六进制字符串.
func (tc *TsConvert) Bin2Hex(str string) (s string, err error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return
	}
	s = strconv.FormatInt(i, 16)
	return
}

// Dec2Hex 将十进制转换为十六进制.
func (tc *TsConvert) Dec2Hex(number int64) string {
	return strconv.FormatInt(number, 16)
}

// Hex2Dec 将十六进制转换为十进制.
func (tc *TsConvert) Hex2Dec(str string) (i int64, err error) {
	start := 0
	if len(str) > 2 && str[0:2] == "0x" {
		start = 2
	}
	// bitSize 表示结果的位宽(包括符号位),0 表示最大位宽
	i, err = strconv.ParseInt(str[start:], 16, 0)
	return
}

// Dec2Oct 将十进制转换为八进制.
func (tc *TsConvert) Dec2Oct(number int64) string {
	return strconv.FormatInt(number, 8)
}

// Oct2Dec 将八进制转换为十进制.
func (tc *TsConvert) Oct2Dec(str string) (i int64, err error) {
	start := 0
	if len(str) > 1 && str[0:1] == "0" {
		start = 1
	}

	i, err = strconv.ParseInt(str[start:], 8, 0)
	return
}

// BaseConvert 进制转换,在任意进制之间转换数字.
func (tc *TsConvert) BaseConvert(number string, fromBase, toBase int) (s string, err error) {
	i, err := strconv.ParseInt(number, fromBase, 0)
	if err != nil {
		return
	}
	s = strconv.FormatInt(i, toBase)
	return
}

// Ip2Long 将 IPV4 的字符串互联网协议转换成长整型数字.
func (tc *TsConvert) Ip2Long(ipAddress string) (l uint32) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return
	}
	l = binary.BigEndian.Uint32(ip.To4())
	return
}

// Long2Ip 将长整型转化为字符串形式带点的互联网标准格式地址(IPV4).
func (tc *TsConvert) Long2Ip(properAddress uint32) (ips string) {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ips = net.IP(ipByte).String()
	return
}

// Gettype 获取变量类型.
func (tc *TsConvert) Gettype(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// ToStr 强制将变量转换为字符串.
func (tc *TsConvert) ToStr(val interface{}) string {
	//先处理其他类型
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Ptr, reflect.Struct, reflect.Map: //指针、结构体和字典
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return ""
		}
		return string(b)
	}

	//再处理字节切片
	switch val.(type) {
	case []uint8:
		return string(val.([]uint8))
	}

	return fmt.Sprintf("%v", val)
}

// ToBool 强制将变量转换为布尔值.
func (tc *TsConvert) ToBool(val interface{}) bool {
	switch val.(type) {
	case int:
		return val.(int) > 0
	case int8:
		return val.(int8) > 0
	case int16:
		return val.(int16) > 0
	case int32:
		return val.(int32) > 0
	case int64:
		return val.(int64) > 0
	case uint:
		return val.(uint) > 0
	case uint8:
		return val.(uint8) > 0
	case uint16:
		return val.(uint16) > 0
	case uint32:
		return val.(uint32) > 0
	case uint64:
		return val.(uint64) > 0
	case float32:
		return val.(float32) > 0
	case float64:
		return val.(float64) > 0
	case []uint8:
		return tc.Str2Bool(string(val.([]uint8)))
	case string:
		return tc.Str2Bool(val.(string))
	case bool:
		return val.(bool)
	default:
		return false
	}
}

// ToInt 强制将变量转换为整型;其中true 或 "true"为1.
func (tc *TsConvert) ToInt(val interface{}) int {
	switch val.(type) {
	case int:
		return val.(int)
	case int8:
		return int(val.(int8))
	case int16:
		return int(val.(int16))
	case int32:
		return int(val.(int32))
	case int64:
		return int(val.(int64))
	case uint:
		return int(val.(uint))
	case uint8:
		return int(val.(uint8))
	case uint16:
		return int(val.(uint16))
	case uint32:
		return int(val.(uint32))
	case uint64:
		return int(val.(uint64))
	case float32:
		return int(val.(float32))
	case float64:
		return int(val.(float64))
	case []uint8:
		return tc.Str2Int(string(val.([]uint8)))
	case string:
		return tc.Str2Int(val.(string))
	case bool:
		return tc.Bool2Int(val.(bool))
	default:
		return 0
	}
}

// ToFloat 强制将变量转换为浮点型;其中true或"true"为1.0 .
func (tc *TsConvert) ToFloat(val interface{}) (toFloat float64) {
	switch val.(type) {
	case int:
		toFloat = float64(val.(int))
	case int8:
		toFloat = float64(val.(int8))
	case int16:
		toFloat = float64(val.(int16))
	case int32:
		toFloat = float64(val.(int32))
	case int64:
		toFloat = float64(val.(int64))
	case uint:
		toFloat = float64(val.(uint))
	case uint8:
		toFloat = float64(val.(uint8))
	case uint16:
		toFloat = float64(val.(uint16))
	case uint32:
		toFloat = float64(val.(uint32))
	case uint64:
		toFloat = float64(val.(uint64))
	case float32:
		toFloat = float64(val.(float32))
	case float64:
		toFloat = val.(float64)
	case []uint8:
		toFloat = tc.Str2Float64(string(val.([]uint8)))
	case string:
		toFloat = tc.Str2Float64(val.(string))
	case bool:
		if val.(bool) {
			toFloat = 1.0
		}
	}

	return
}

// Float64ToByte 64位浮点数转字节切片.
func (tc *TsConvert) Float64ToByte(val float64) (parseFloat []byte) {
	bits := math.Float64bits(val)
	parseFloat = make([]byte, 8)
	binary.LittleEndian.PutUint64(parseFloat, bits)
	return
}

// Byte2Float64 字节切片转64位浮点数.
func (tc *TsConvert) Byte2Float64(bytes []byte) (f float64) {
	bits := binary.LittleEndian.Uint64(bytes)
	f = math.Float64frombits(bits)
	return
}

// Int64ToByte 64位整型转字节切片.
func (tc *TsConvert) Int64ToByte(val int64) []byte {
	parseByte := make([]byte, 8)
	binary.BigEndian.PutUint64(parseByte, uint64(val))
	return parseByte
}

// Byte2Int64 字节切片转64位整型.
func (tc *TsConvert) Byte2Int64(val []byte) int64 {
	return int64(binary.BigEndian.Uint64(val))
}

// Byte2Hex 字节切片转16进制字符串.
func (tc *TsConvert) Byte2Hex(val []byte) string {
	return hex.EncodeToString(val)
}

// Hex2Byte 16进制字符串转字节切片.
func (tc *TsConvert) Hex2Byte(str string) (parseHex []byte) {
	parseHex, _ = hex.DecodeString(str)
	return
}

// GetPointerAddrInt 获取变量指针地址整型值.variable为变量.
func (tc *TsConvert) GetPointerAddrInt(variable interface{}) (addr int64) {
	addr, _ = tc.Hex2Dec(fmt.Sprintf("%p", &variable))
	return
}