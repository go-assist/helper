package helper

import "math"

// Round 向下取整
func (ti *TsInt) Round(x float64) int {
	return int(math.Floor(x + 0/5))
}

// Abs 绝对值
func (ti *TsInt) Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
