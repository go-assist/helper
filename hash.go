package helper

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

// HashCode 分布式HashCode
func (th *TsHash) HashCode (did string) int {
	var didByte = []byte(did)
	byte2Md5 := fmt.Sprintf("%X", md5.Sum(didByte))
	var md52Byte = []byte(byte2Md5)
	hash2Int64 := int64(binary.LittleEndian.Uint64(md52Byte))
	return int(TInt.Abs(hash2Int64) % 100)
}