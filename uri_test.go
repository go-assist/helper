package helper

import (
	"testing"
)

func TestParseUriQueryToMap(t *testing.T) {

	m := TUri.ParseUriQueryToMap(exampleUriStr)
	if _, ok := m["Av"]; !ok {
		t.Errorf("ParseUriQueryToMap unit test fail")
	}
}

func TestGetQueryParams(t *testing.T) {
	m := make(UriValues)
	m, mError := TUri.ParseUriQuery(m, exampleUriStr)
	if mError != nil {
		t.Errorf("parseQuery Errors:%v \n", mError)
	}
	av := TUri.GetQueryParams(m, "av")
	if av == "" {
		t.Errorf("the values of %v is not %v \n", "5.3.5", "")
	}
}

func TestGetDomain(t *testing.T) {
	actual := ""
	for _, test := range exampleUriTests {
		actual = TUri.GetDomain(test.param, test.isMain)
		if actual != test.expected {
			t.Errorf("Expected GetDomain(%q) to be %v, got is %v \n", test.param, test.expected, actual)
		}
	}
	TUri.GetDomain("123456")
}
