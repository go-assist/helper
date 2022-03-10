package helper

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

// getTrimMask 去除mask字符
func getTrimMask(characterMask []string) string {
	var mask string
	if len(characterMask) == 0 {
		mask = " \t\n\r\v\f　"
	} else {
		mask = strings.Join(characterMask, "")
	}
	return mask
}

// isInt 变量是否整型数值
func isInt(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		_, err := strconv.Atoi(str)
		return err == nil
	}

	return false
}

// isNumeric 变量是否数值(不包含复数和科学计数法)
func isNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		_, err := strconv.ParseFloat(str, 64)
		return err == nil
	}

	return false
}

// isFloat 变量是否浮点数值
func isFloat(val interface{}) bool {
	switch val.(type) {
	case float32, float64:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}

		if ok := RegFloat.MatchString(str); ok {
			return true
		}
	}

	return false
}

// isMap 检查变量是否map
func isMap(varMap interface{}) bool {
	return reflect.ValueOf(varMap).Kind() == reflect.Map
}

// isStruct 检查变量是否是struct
func isStruct(varStruct interface{}) bool {
	return reflect.TypeOf(varStruct).Kind() == reflect.Struct
}

// creditChecksum 计算身份证校验码,其中id为身份证号码
func creditChecksum(id string) byte {
	//∑(ai×Wi)(mod 11)
	// 加权因子
	factor := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 校验位对应值
	code := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

	length := len(id)
	sum := 0
	for i, char := range id[:length-1] {
		num, _ := strconv.Atoi(string(char))
		sum += num * factor[i]
	}

	return code[sum%11]
}

// reflectPtr 获取反射的指向.
func reflectPtr(r reflect.Value) reflect.Value {
	// 如果是指针,则获取其所指向的元素
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	return r
}

func isArrayOrSliceHelper(data interface{}, chkType uint8) int {
	if chkType != 1 && chkType != 2 && chkType != 3 {
		panic(fmt.Sprintf("[isArrayOrSliceHelper]chkType value muset in (1, 2, 3), but it`s %d", chkType))
	}

	var res = -1
	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Array:
		if chkType == 1 || chkType == 3 {
			res = val.Len()
		}
	case reflect.Slice:
		if chkType == 2 || chkType == 3 {
			res = val.Len()
		}
	}

	return res
}

// getProcessExeByPid 根据PID获取进程的执行路径.
func getProcessExeByPid(pid int) string {
	exe := fmt.Sprintf("/proc/%d/exe", pid)
	path, _ := os.Readlink(exe)
	return path
}

// getPidByInode 根据套接字的inode获取PID.须root权限.
func getPidByInode(inode string, procDirs []string) (pid int) {
	if len(procDirs) == 0 {
		procDirs, _ = filepath.Glob("/proc/[0-9]*/fd/[0-9]*")
	}
	re := regexp.MustCompile(inode)
	for _, item := range procDirs {
		path, _ := os.Readlink(item)
		out := re.FindString(path)
		if len(out) != 0 {
			pid, _ = strconv.Atoi(strings.Split(item, "/")[2])
			break
		}
	}
	return pid
}

// getEndian 获取系统字节序类型,小端返回binary.LittleEndian,大端返回binary.BigEndian .
func getEndian() binary.ByteOrder {
	var nativeEndian binary.ByteOrder = binary.LittleEndian
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		nativeEndian = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		nativeEndian = binary.BigEndian
	}

	return nativeEndian
}

// isLittleEndian 系统字节序类型是否小端存储.
func isLittleEndian() bool {
	var i int32 = 0x01020304
	// 将int32类型的指针转换为byte类型的指针
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	// 取得pb位置对应的值
	b := *pb
	// 由于b是byte类型的,最多保存8位,那么只能取得开始的8位
	// 小端: 04 (03 02 01)
	// 大端: 01 (02 03 04)
	return b == 0x04
}

// hex2Decimal 十六进制转十进制
func hex2Decimal(str string) (decimal int) {
	i, err := strconv.ParseUint(str, 16, 32)
	if err != nil {
		return
	}
	decimal = int(i)
	return
}