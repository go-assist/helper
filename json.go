package helper

import (
	"encoding/json"
	"errors"
	jsonIter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
	"log"
	"reflect"
)

// ParseJson Json校验.
func (tj *TsJson) ParseJson(json string) (result gjson.Result, err error) {
	if !gjson.Valid(json) {
		err = errors.New("invalid json")
		return
	}
	result = gjson.Parse(json)
	return
}

// MapToJson map转为json字符串.
func (tj *TsJson) MapToJson(m map[string]interface{}) string {
	m2Json , _ := json.Marshal(m)
	map2String := string(m2Json)
	return map2String
}

// JsonToMap json 转map.
func (tj *TsJson) JsonToMap(jsonStr string) map[string]interface{} {
	var convert map[string]interface{}
	if jsonStr == "" {
		return convert
	}
	err := json.Unmarshal([]byte(jsonStr), &convert)
	if err != nil {
		log.Println(err)
	}
	return convert
}


// JsonToMapArr json转map数组.
func (tj *TsJson) JsonToMapArr(jsonStr string) []map[string]interface{} {
	var convert []map[string]interface{}
	if jsonStr == "" {
		return convert
	}
	err := json.Unmarshal([]byte(jsonStr), &convert)
	if err != nil {
		log.Println(err)
	}
	return convert
}

// StructToMap 结构体转map.
func (tj *TsJson) StructToMap(obj interface{}) map[string]interface{} {
	convert := make(map[string]interface{})
	if isStruct(obj) {
		typeOf := reflect.TypeOf(obj)
		valueOf := reflect.ValueOf(obj)
		for i := 0; i < typeOf.NumField(); i ++ {
			convert[typeOf.Field(i).Name] = valueOf.Field(i).Interface()
		}
	}
	return convert
}

// MapToStruct map转struct
func (tj *TsJson) MapToStruct(obj interface{}, outStruct interface{}) (interface{}, error) {
	err := mapstructure.Decode(obj, &outStruct)
	return outStruct, err
}

// JsonEncode 对变量进行 JSON 编码.
func (tj *TsJson) JsonEncode(val interface{}) ([]byte, error) {
	var jsons = jsonIter.ConfigCompatibleWithStandardLibrary
	return jsons.Marshal(val)
}

// JsonDecode 对 JSON 格式的字符串进行解码,注意val使用指针.
func (tj *TsJson) JsonDecode(data []byte, val interface{}) error {
	var jsons = jsonIter.ConfigCompatibleWithStandardLibrary
	return jsons.Unmarshal(data, val)
}