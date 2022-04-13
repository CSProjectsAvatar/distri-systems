package utils

import "testing"

func TestCheckErrNotNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic, but didn't get one")
		}
	}()
	CheckErr(fakeErr("hello"))
}

func TestCheckErrNil(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Expected no panic, but got one")
		}
	}()
	CheckErr(nil)
}

type fakeErr string

func (e fakeErr) Error() string {
	return string(e)
}
