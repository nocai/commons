package bizerrs

import (
	"testing"
)

func Test(t *testing.T) {
	if ErrSystem.Code() != "-1" {
		t.Fail()
	}
}
