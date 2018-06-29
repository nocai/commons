package mdl

import (
	"testing"
)

func TestGetKInt(t *testing.T) {
	kv := KV{Key: 1, Value: "value"}

	if kv.GetKInt() != 1 {
		t.Fail()
	}
}
