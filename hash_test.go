package helper

import (
	"testing"
)

func TestHashCode(t *testing.T) {
	want1 := 52
	got1 := "abc123tre"
	h1 := THash.HashCode(got1,2)
	if want1 != h1 {
		t.Errorf("The HashCode values of %v is not %v\n", want1, h1)
	}

	want2 := 2
	got2 := ""
	h2 := THash.HashCode(got2,3)
	if want2 != h2 {
		t.Errorf("The HashCode values of %v is not %v\n", want2, h2)
	}

}
