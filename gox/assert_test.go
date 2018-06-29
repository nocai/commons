package gox

import (
	"testing"
)

func TestAssert(t *testing.T) {
	msg := "F"
	defer func() {
		if r := recover(); r != msg {
			t.Fail()
		}
	}()
	Assert(false, "F")
}
