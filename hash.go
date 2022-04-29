package helper

import (
	"crypto/md5"
	"fmt"
)

// HashCode 分布式HashCode.
func (th *TsHash) HashCode (did string, step int) (score int) {
	didByte := []byte(did)
	md5Str := fmt.Sprintf("%X", md5.Sum(didByte))
	// 倒数第 1 3 5 7 位, 可随机
	// hexadecimalStr := md5Str[step:step+1] + md5Str[29:30] + md5Str[27:28] + md5Str[25:26] + md5Str[17:18] + md5Str[10:11]
	hexadecimalStr := md5Str[step:step+1] + md5Str[29:30] + md5Str[27:28] + md5Str[25:26] + md5Str[17:18]
	// 1a2b--> 转10进制-> 移位-> mod
	dec := hex2Decimal(hexadecimalStr)
	i := dec >> 2 // 移位
	score = int(TInt.Abs(int64(i)) % 100)
	return
}