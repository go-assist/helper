package helper

import "math"

// EqualFloat 比较两个浮点数是否相等.decimal为小数精确位数.
func (tf *TsFloat) EqualFloat(f1 float64, f2 float64, decimal ...int) (r bool) {
	var threshold float64
	var dec int
	if len(decimal) == 0 {
		dec = FloatDecimal
	} else {
		dec = decimal[0]
	}

	//比较精度
	threshold = math.Pow10(-dec)

	r = math.Abs(f1-f2) <= threshold
	return
}

// Round 对浮点数进行四舍五入.
func (tf *TsFloat) Round(value float64) float64 {
	return math.Floor(value + 0.5)
}

// Ceil 向上取整.
func (tf *TsFloat) Ceil(value float64) float64 {
	return math.Ceil(value)
}

// Floor 向下取整.
func (tf *TsFloat) Floor(value float64) float64 {
	return math.Floor(value)
}

// MaxFloat64 64位浮点数序列求最大值.
func (tf *TsFloat) MaxFloat64(nums ...float64) (maxFloat float64) {
	if len(nums) < 1 {
		panic("[MaxFloat64]: the nums length is less than 1")
	}

	maxFloat = nums[0]
	for _, v := range nums {
		maxFloat = math.Max(maxFloat, v)
	}

	return
}

// MinFloat64 64位浮点数序列求最小值.
func (tf *TsFloat) MinFloat64(nums ...float64) (minFloat float64) {
	if len(nums) < 1 {
		panic("[MinFloat64]: the nums length is less than 1")
	}
	minFloat = nums[0]
	for _, v := range nums {
		minFloat = math.Min(minFloat, v)
	}
	return
}

// SumFloat64 浮点数求和.
func (tf *TsFloat) SumFloat64(nums ...float64) (sum float64) {
	for _, v := range nums {
		sum += v
	}
	return sum
}