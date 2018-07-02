package mdl

import (
	"testing"
)

func TestGetKInt(t *testing.T) {
	kvp := Kvp{Key: 1, Value: "value"}

	if kvp.GetKInt() != 1 {
		t.Fail()
	}
}
