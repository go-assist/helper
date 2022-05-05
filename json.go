package helper

import (
	"bytes"
	"errors"
	jsonIter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

// ParseJson Json校验.
func (tj *TsJson) ParseJson(jsonStr string) (r GJsonResult, err error) {
	if !TValidate.IsJSONGJson(jsonStr) || !TValidate.IsJSON(jsonStr) {
		err = errors.New("invalid json")
		return
	}
	r = parseJson(jsonStr)
	return
}

// MapToJson map转为json字符串.
func (tj *TsJson) MapToJson(m map[string]interface{}) (j string) {
	var jsons = jsonIter.ConfigCompatibleWithStandardLibrary
	m2Json, _ := jsons.Marshal(m)
	j = string(m2Json)
	return
}

// JsonToMap json 转map.
func (tj *TsJson) JsonToMap(jsonStr string) (convert map[string]interface{}) {
	if jsonStr == "" || !TValidate.IsJSON(jsonStr) {
		return convert
	}
	var jsons = jsonIter.ConfigCompatibleWithStandardLibrary
	err := jsons.Unmarshal([]byte(jsonStr), &convert)
	if err != nil {
		return
	}
	return
}

// JsonToMapArr json转map数组.
func (tj *TsJson) JsonToMapArr(jsonStr string) (convert []map[string]interface{}) {
	if jsonStr == "" || !TValidate.IsJSON(jsonStr) {
		return convert
	}
	var jsons = jsonIter.ConfigCompatibleWithStandardLibrary
	err := jsons.Unmarshal([]byte(jsonStr), &convert)
	if err != nil {
		return
	}
	return
}

// StructToMap 结构体转map.
func (tj *TsJson) StructToMap(obj interface{}) map[string]interface{} {
	convert := make(map[string]interface{})
	if isStruct(obj) {
		typeOf := reflect.TypeOf(obj)
		valueOf := reflect.ValueOf(obj)
		for i := 0; i < typeOf.NumField(); i++ {
			convert[typeOf.Field(i).Name] = valueOf.Field(i).Interface()
		}
	}
	return convert
}

// MapToStruct map转struct.
func (tj *TsJson) MapToStruct(obj interface{}, outStruct interface{}) (interface{}, error) {
	err := mapstructure.Decode(obj, &outStruct)
	return outStruct, err
}

// JsonEncode 对变量进行 JSON 编码并去除转移字符.
func (tj *TsJson) JsonEncode(val interface{}) (b []byte, err error) {
	var jsons = jsonIter.ConfigCompatibleWithStandardLibrary
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := jsons.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err = jsonEncoder.Encode(val)
	if err != nil {
		return
	}
	b = []byte(bf.String())
	return
}

// JsonDecode 对 JSON 格式的字符串进行解码,注意val使用指针.
func (tj *TsJson) JsonDecode(data []byte, val interface{}) error {
	var jsons = jsonIter.ConfigCompatibleWithStandardLibrary
	return jsons.Unmarshal(data, val)
}
