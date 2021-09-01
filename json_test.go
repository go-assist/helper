package helper

import (
	"testing"
)

func TestParseJson(t *testing.T) {
	_, err := TJson.ParseJson(jsonExample)
	if err != nil {
		t.Errorf("parse json errors of %v \n", err.Error())
	}
	_, err = TJson.ParseJson(errJson)
	if err == nil {
		t.Errorf("parse json errors")
	}
}

func TestMapToJson(t *testing.T) {
	jsonStr := TJson.MapToJson(jsonMap)
	_, err := TJson.ParseJson(jsonStr)
	if err != nil {
		t.Errorf("parse json errors of %v \n", err.Error())
	}
}

func TestJsonToMap(t *testing.T) {
	m := TJson.JsonToMap(jsonExample)
	if !isMap(m) {
		t.Errorf("reflect valueof does not map \n")
	}
	if _, ok := m["k1"]; !ok {
		t.Errorf("map conv errors\n")
	}
}

func TestJsonToMapArray(t *testing.T) {
	m := TJson.JsonToMapArr(jsonArr)
	if !isMap(m) || len(m) != 2 {
		t.Errorf("JsonToMapArr errors\n")
	}
}

func TestStructToMap(t *testing.T) {
	example := Example{
		Examples: "test",
	}
	m := TJson.StructToMap(example)
	if !isMap(m) {
		t.Errorf("StructToMap does not map")
	}

	if _, ok := m["Examples"]; !ok {
		t.Errorf("StructToMap key does not exists\n")
	}
}

func TestMapToStruct(t *testing.T) {

	var ex Example
	ex1, err := TJson.MapToStruct(jsonStruct, ex)
	if err != nil {
		t.Errorf("MapToStruct errors \n")
	}
	if ex1.(Example).Examples != "test" {
		t.Errorf("MapToStruct values of %v, not test \n", ex1.(Example).Examples)
	}
}

func TestJsonEncode(t *testing.T) {
	bJson, err := TJson.JsonEncode(jsonExample)
	if err != nil {
		t.Errorf("JsonEncode errors: %v\n", err)
	}
	_, err = TJson.ParseJson(*TStr.ByteToString(&bJson))
	if err != nil {
		t.Errorf("JsonEncode parse json errors of %v \n", err.Error())
	}
}

func TestJsonDecode(t *testing.T) {
	var i interface{}
	err := TJson.JsonDecode([]byte(jsonExample), &i)
	if err != nil {
		t.Errorf("JsonDecode errors: %v\n", err)
	}
}

