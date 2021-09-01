package helper

import "testing"

func TestUuid(t *testing.T) {
	uuid := TUuid.Uuid()
	if uuid == "false" {
		t.Error("uuid build errors")
	}
}