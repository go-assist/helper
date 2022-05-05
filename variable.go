package helper

import (
	"net"
	"regexp"
	"time"
)

var (
	TArr             TsArr
	TStr             TsStr // transfer
	TInt             TsInt
	TFloat           TsFloat
	TFile            TsFile
	TUri             TsUri
	THash            TsHash
	TCallFunc        TsCallFunc
	TDebug           TsDebug
	TOs              TsOs
	TConv            TsConvert
	TTime            TsTime
	TUuid            TsUuid
	TCorn            TsCorn
	TValidate        TsValidate
	TJson            TsJson
	TEncrypt         TsEncrypt
	TPrivyCiders     []*net.IPNet
	NowTime          = time.Now()
	RegDNSName       = regexp.MustCompile(PatternDNSName)
	RegMultiByte     = regexp.MustCompile(PatternMultibyte)
	RegFloat         = regexp.MustCompile(PatternFloat)
	RegFullWidth     = regexp.MustCompile(PatternFullwidth)
	RegHalfWidth     = regexp.MustCompile(PatternHalfWidth)
	RegAlphaLower    = regexp.MustCompile(PatternAlphaLower)
	RegAlphaUpper    = regexp.MustCompile(PatternAlphaUpper)
	RegChineseAll    = regexp.MustCompile(PatternChineseAll)
	RegChineseName   = regexp.MustCompile(PatternChineseName)
	RegEmail         = regexp.MustCompile(PatternEmail)
	RegMobileCN      = regexp.MustCompile(PatternMobileCN)
	RegTelephone     = regexp.MustCompile(PatternTelephone)
	RegPhone         = regexp.MustCompile(PatternPhone)
	RegCreditNo      = regexp.MustCompile(PatternCreditNo)
	RegDatetime      = regexp.MustCompile(PatternDatetime)
	RegAlphaNumeric  = regexp.MustCompile(PatternAlphaNumeric)
	RegRGBColor      = regexp.MustCompile(PatternRGBColor)
	RegWhitespaceAll = regexp.MustCompile(PatternWhitespaceAll)
	RegWhitespaceHas = regexp.MustCompile(PatternWhitespaceHas)
	RegBase64        = regexp.MustCompile(PatternBase64)
	RegBase64Image   = regexp.MustCompile(PatternBase64Image)
	RegMd5           = regexp.MustCompile(PatternMd5)
	RegSha1          = regexp.MustCompile(PatternSha1)
	RegSha256        = regexp.MustCompile(PatternSha256)
	RegSha512        = regexp.MustCompile(PatternSha512)
	// CreditArea 身份证区域
	CreditArea = map[string]string{
		"11": "北京",
		"12": "天津",
		"13": "河北",
		"14": "山西",
		"15": "内蒙古",
		"21": "辽宁",
		"22": "吉林",
		"23": "黑龙江",
		"31": "上海",
		"32": "江苏",
		"33": "浙江",
		"34": "安徽",
		"35": "福建",
		"36": "江西",
		"37": "山东",
		"41": "河南",
		"42": "湖北",
		"43": "湖南",
		"44": "广东",
		"45": "广西",
		"46": "海南",
		"50": "重庆",
		"51": "四川",
		"52": "贵州",
		"53": "云南",
		"54": "西藏",
		"61": "陕西",
		"62": "甘肃",
		"63": "青海",
		"64": "宁夏",
		"65": "新疆",
		"71": "台湾",
		"81": "香港",
		"82": "澳门",
		"91": "国外"}
)
