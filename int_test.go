package helper

import "testing"

func TestRound(t *testing.T) {
	got := 3
	want1 := TInt.Round(3.14)
	if got != want1 {
		t.Errorf("The Round values of %v is not %v\n", want1, got)
	}

	want2 := TInt.Round(3.55)
	if got != want2 {
		t.Errorf("The Round values of %v is not %v\n", want2, got)
	}
}

func TestAbs(t *testing.T) {
	got1 := -1
	want1 := int64(1)
	abs1 := TInt.Abs(int64(got1))
	if abs1 != want1 {
		t.Errorf("The Abs values of %v is not %v\n", abs1, want1)
	}

	got2 := 0
	want2 := int64(0)
	abs2 := TInt.Abs(int64(got2))
	if abs2 != want2 {
		t.Errorf("The Abs values of %v is not %v\n", abs2, want2)
	}

	got3 := 1
	want3 := int64(1)
	abs3 := TInt.Abs(int64(got3))
	if abs3 != want3 {
		t.Errorf("The Abs values of %v is not %v\n", abs3, want3)
	}
}
