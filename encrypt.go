package helper

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"hash"
	"strconv"
	"strings"
	"time"
)

// Base64Encode 使用 MIME base64 对数据进行编码.
func (te *TsEncrypt) Base64Encode(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

// Base64Decode 对使用 MIME base64 编码的数据进行解码.
func (te *TsEncrypt) Base64Decode(str string) (b []byte, err error) {
	switch len(str) % 4 {
	case 2:
		str += "=="
	case 3:
		str += "="
	}

	decodeString, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return
	}
	b = decodeString
	return
}

// Base64UrlEncode Base64UrlSafeEncode url安全的Base64Encode,没有'/'和'+'及结尾的'=' .
func (te *TsEncrypt) Base64UrlEncode(source []byte) (safeUrl string) {
	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
	byteArr := base64.StdEncoding.EncodeToString(source)
	safeUrl = strings.Replace(byteArr, "/", "_", -1)
	safeUrl = strings.Replace(safeUrl, "+", "-", -1)
	safeUrl = strings.Replace(safeUrl, "=", "", -1)
	return
}

// Base64UrlDecode url安全的Base64Decode.
func (te *TsEncrypt) Base64UrlDecode(data string) (b []byte, err error) {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)
	b, err = base64.URLEncoding.DecodeString(data)
	return
}

// AuthCode 授权码编码或解码;encode为true时编码,为false解码;expiry为有效期,秒;返回结果为加密/解密的字符串和有效期时间戳.
func (te *TsEncrypt) AuthCode(str, key string, encode bool, expiry int64) (auth string, expire int64) {
	// DYNAMIC_KEY_LEN 动态密钥长度,相同的明文会生成不同密文就是依靠动态密钥
	// 加入随机密钥,可以令密文无任何规律,即便是原文和密钥完全相同,加密结果也会每次不同,增大破解难度。
	// 取值越大,密文变动规律越大,密文变化 = 16 的 DYNAMIC_KEY_LEN 次方
	// 当此值为 0 时,则不产生随机密钥

	if str == "" {
		return
	}

	if !encode && len(str) < DynamicKeyLen {
		return
	}

	// 密钥
	keyByte := TStr.Md5Hex([]byte(key), 32)

	// 密钥a会参与加解密
	keyA := TStr.Md5Hex(keyByte[:16], 32)

	// 密钥b会用来做数据完整性验证
	keyB := TStr.Md5Hex(keyByte[16:], 32)

	// 密钥c用于变化生成的密文
	var keyC []byte
	if encode == false {
		keyC = []byte(str[:DynamicKeyLen])
	} else {
		cLen := 32 - DynamicKeyLen
		now, _ := time.Now().MarshalBinary()
		timeBytes := TStr.Md5Hex(now, 32)
		keyC = timeBytes[cLen:]
	}

	// 参与运算的密钥
	keyD := TStr.Md5Hex(append(keyA, keyC...), 32)
	encryptionKey := append(keyA, keyD...)
	keyLength := len(encryptionKey)
	// 明文,前10位用来保存时间戳,解密时验证数据有效性,10到26位用来保存keyB(密钥b),解密时会通过这个密钥验证数据完整性
	// 如果是解码的话,会从第 DYNAMIC_KEY_LEN 位开始,因为密文前 DYNAMIC_KEY_LEN 位保存 动态密钥,以保证解密正确
	if encode == false {
		strByte, err := te.Base64UrlDecode(str[DynamicKeyLen:])
		if err != nil {
			return
		}
		str = string(strByte)
	} else {
		if expiry != 0 {
			expiry = expiry + time.Now().Unix()
		}
		expMd5 := TStr.Md5Hex(append([]byte(str), keyB...), 16)
		str = fmt.Sprintf("%010d%s%s", expiry, expMd5, str)
	}
	stringLength := len(str)
	resData := make([]byte, 0, stringLength)
	var rndKey, box [256]int
	// 产生密钥簿
	j := 0
	a := 0
	i := 0
	for i = 0; i < 256; i++ {
		rndKey[i] = int(encryptionKey[i%keyLength])
		box[i] = i
	}
	// 用固定的算法,打乱密钥簿,增加随机性,好像很复杂,实际上并不会增加密文的强度
	for i = 0; i < 256; i++ {
		j = (j + box[i] + rndKey[i]) % 256
		box[i], box[j] = box[j], box[i]
	}
	// 核心加解密部分
	a = 0
	j = 0
	for i = 0; i < stringLength; i++ {
		a = (a + 1) % 256
		j = (j + box[a]) % 256
		box[a], box[j] = box[j], box[a]
		// 从密钥簿得出密钥进行异或,再转成字符
		resData = append(resData, byte(int(str[i])^box[(box[a]+box[j])%256]))
	}
	result := string(resData)
	if encode == false { //解密
		// substr($result, 0, 10) == 0 验证数据有效性
		// substr($result, 0, 10) - time() > 0 验证数据有效性
		// substr($result, 10, 16) == substr(md5(substr($result, 26).$keyB), 0, 16) 验证数据完整性
		// 验证数据有效性,请看未加密明文的格式
		if len(result) <= 26 {
			return
		}

		expTime, _ := strconv.ParseInt(result[:10], 10, 0)
		if (expTime == 0 || expTime-time.Now().Unix() > 0) && result[10:26] == string(TStr.Md5Hex(append(resData[26:], keyB...), 16)) {
			auth = result[26:]
			expire = expTime
			return
		} else {
			expire = expTime
			return
		}
	} else { //加密
		// 把动态密钥保存在密文里,这也是为什么同样的明文,生产不同密文后能解密的原因
		auth = string(keyC) + te.Base64UrlEncode(resData)
		expire = expiry
		return
	}
}

// PasswordHash 创建密码的散列值;costs为算法的cost,范围4~31,默认10,注意值越大越耗时.
func (te *TsEncrypt) PasswordHash(password []byte, costs ...int) (bytes []byte, err error) {
	var cost int
	if len(costs) == 0 {
		cost = 10
	} else {
		cost = costs[0]
		if cost < 4 {
			cost = 4
		} else if cost > 31 {
			cost = 15
		}
	}

	bytes, err = bcrypt.GenerateFromPassword(password, cost)
	return
}

// PasswordVerify 验证密码是否和散列值匹配.
func (te *TsEncrypt) PasswordVerify(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

// EasyEncrypt 简单加密.
// data为要加密的原字符串,key为密钥.
func (te *TsEncrypt) EasyEncrypt(data, key string) (encr string) {
	dataLen := len(data)
	if dataLen == 0 {
		return
	}

	keyByte := TStr.Md5Hex([]byte(key), 32)
	keyLen := len(keyByte)

	var i, x, c int
	var str []byte
	for i = 0; i < dataLen; i++ {
		if x == keyLen {
			x = 0
		}

		c = (int(data[i]) + int(keyByte[x])) % 256
		str = append(str, byte(c))

		x++
	}

	encr = string(keyByte[:DynamicKeyLen]) + te.Base64UrlEncode(str)
	return
}

// EasyDecrypt 简单解密.
// val为待解密的字符串,key为密钥.
func (te *TsEncrypt) EasyDecrypt(val, key string) (decr string) {
	if len(val) <= DynamicKeyLen {
		return
	}

	data, err := te.Base64UrlDecode(val[DynamicKeyLen:])
	if err != nil {
		return
	}

	keyByte := TStr.Md5Hex([]byte(key), 32)
	if val[:DynamicKeyLen] != string(keyByte[:DynamicKeyLen]) {
		return
	}

	dataLen := len(data)
	keyLen := len(keyByte)

	var i, x, c int
	var str []byte
	for i = 0; i < dataLen; i++ {
		if x == keyLen {
			x = 0
		}
		if data[i] < keyByte[x] {
			c = int(data[i]) + 256 - int(keyByte[x])
		} else {
			c = int(data[i]) - int(keyByte[x])
		}
		str = append(str, byte(c))
		x++
	}
	decr = string(str)
	return
}

// HmacShaX HmacSHA-x加密,x为1/256/512.
func (te *TsEncrypt) HmacShaX(data, secret []byte, x uint16) (hr string) {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	var h hash.Hash
	switch x {
	case 1:
		h = hmac.New(sha1.New, secret)
		break
	case 256:
		h = hmac.New(sha256.New, secret)
		break
	case 512:
		h = hmac.New(sha512.New, secret)
		break
	default:
		panic("[HmacShaX] x must be in [1, 256, 512]")
	}

	// Write Data to it
	h.Write(data)

	// Get result and encode as hexadecimal string
	hr = hex.EncodeToString(h.Sum(nil))
	return
}

