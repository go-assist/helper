package helper

import "time"

const (
	Int64Max  = int64(^uint(0) >> 1)
	Uint64Max = ^uint64(0)
	// PatternFloat 正则模式-浮点数
	PatternFloat = `^(-?\d+)(\.\d+)?`
	// PatternDNSName 正则模式DNS名称
	PatternDNSName = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	// PatternMultibyte 正则模式-多字节字符
	PatternMultibyte = "[^\x00-\x7F]"
	// PatternFullwidth 正则模式-全角字符
	PatternFullwidth = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	// PatternHalfWidth 正则模式-半角字符
	PatternHalfWidth = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	// PatternAlphaLower 正则模式-小写英文
	PatternAlphaLower = `^[a-z]+$`
	// PatternAlphaUpper 正则模式-大写英文
	PatternAlphaUpper = `^[A-Z]+$`
	// PatternChineseAll 正则模式-全中文
	PatternChineseAll = "^[\u4e00-\u9fa5]+$"
	// PatternChineseName 正则模式-中文名称
	PatternChineseName = "^[\u4e00-\u9fa5][.•·\u4e00-\u9fa5]{0,30}[\u4e00-\u9fa5]$"
	// PatternEmail 正则模式-邮箱
	PatternEmail = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	// PatternMobileCN 正则模式-大陆手机号
	PatternMobileCN = `^1[3-9]\d{9}$`
	// PatternTelFix 正则模式-固定电话
	PatternTelFix = `^(010|02\d{1}|0[3-9]\d{2})-\d{7,9}(-\d+)?$`
	// PatternTel4800 正则模式-400或800
	PatternTel4800 = `^[48]00\d?(-?\d{3,4}){2}$`
	// PatternTelephone 正则模式-座机号(固定电话或400或800)
	PatternTelephone = `(` + PatternTelFix + `)|(` + PatternTel4800 + `)`
	// PatternPhone 正则模式-电话(手机或固话)
	PatternPhone = `(` + PatternMobileCN + `)|(` + PatternTelFix + `)`
	// PatternCreditNo 正则模式-身份证号码,18位或15位
	PatternCreditNo = `(^[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)`
	// PatternDatetime 正则模式-日期时间
	PatternDatetime = `^[0-9]{4}(|\-[0-9]{2}(|\-[0-9]{2}(|\s+[0-9]{2}(|:[0-9]{2}(|:[0-9]{2})))))$`
	// PatternAlphaNumeric 正则模式-字母和数字
	PatternAlphaNumeric = `^[a-zA-Z0-9]+$`
	// PatternRGBColor 正则模式-RGB颜色
	PatternRGBColor = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	// PatternWhitespaceAll 正则模式-全空白字符
	PatternWhitespaceAll = "^[[:space:]]+$"
	// PatternWhitespaceHas 正则模式-带空白字符
	PatternWhitespaceHas = ".*[[:space:]]"
	// PatternBase64 正则模式-base64字符串
	PatternBase64 = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	// PatternBase64Image 正则模式-base64编码图片
	PatternBase64Image = `^data:\s*(image|img)\/(\w+);base64`
	// PatternMd5 正则模式-MD5
	PatternMd5 = `^(?i)([0-9a-h]{32})$`
	// PatternSha1 正则模式-SHA1
	PatternSha1 = `^(?i)([0-9a-h]{40})$`
	// PatternSha256 正则模式-SHA256
	PatternSha256 = `^(?i)([0-9a-h]{64})$`
	// PatternSha512 正则模式-SHA512
	PatternSha512 = `^(?i)([0-9a-h]{128})$`
	// CaseNone 忽略大小写
	CaseNone LetterCase = 0
	// CaseLower 检查小写
	CaseLower LetterCase = 1
	// CaseUpper 检查大写
	CaseUpper LetterCase = 2

	// RandStringAlpha 随机字符串类型,字母
	RandStringAlpha RandomString = 0
	// RandStringNumeric 随机字符串类型,数值
	RandStringNumeric RandomString = 1
	// RandStringAluminum 随机字符串类型,字母+数值
	RandStringAluminum RandomString = 2
	// RandStringSpecial 随机字符串类型,字母+数值+特殊字符
	RandStringSpecial RandomString = 3
	// RandStringChinese 随机字符串类型,仅中文
	RandStringChinese RandomString = 4

	// CheckConnectTimeout 检查连接超时的时间
	CheckConnectTimeout = time.Second * 5

	// DynamicKeyLen AuthCode 动态密钥长度,须<32
	DynamicKeyLen = 8

	// FloatDecimal 默认浮点数精确小数位数
	FloatDecimal = 10
)
