package helper

import (
	"fmt"
	"strconv"
)

// Abs 绝对值.
func (ti *TsInt) Abs(n int64) (i int64) {
	y := n >> 63
	i = (n ^ y) - y
	return
}

// AverageInt 对整数序列求平均值.
func (ti *TsInt) AverageInt(nums ...int) (ave float64) {
	length := len(nums)
	if length == 0 {
		return
	} else if length == 1 {
		ave = float64(nums[0])
	} else {
		total := ti.SumInt(nums...)
		ave = float64(total) / float64(length)
	}
	return
}

// SumInt 整数求和.
func (ti *TsInt) SumInt(nums ...int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return
}

// MinInt 整数序列求最小值.
func (ti *TsInt) MinInt(nums ...int) (min int) {
	if len(nums) < 1 {
		panic("[MinInt]: the nums length is less than 1")
	}
	min = nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return
}

// MaxInt 整数序列求最大值.
func (ti *TsInt) MaxInt(nums ...int) (max int) {
	if len(nums) < 1 {
		panic("[MaxInt]: the nums length is less than 1")
	}
	max = nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return
}

// NumberFormat 以千位分隔符方式格式化一个数字.
// decimal为要保留的小数位数,point为小数点显示的字符,thousand为千位分隔符显示的字符.
func (ti *TsInt) NumberFormat(number float64, decimal uint8, point, thousand string) (format string) {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimal)
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousand)
	n, l1, l2 := 0, len(prefix), len(sep)
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	format = string(tmp)
	if dec > 0 {
		format += point + suffix
	}
	if neg {
		format = "-" + format
	}
	return format
}