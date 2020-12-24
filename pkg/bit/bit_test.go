package bit_test

import (
	"testing"

	"github.com/mishamyrt/checode/v1/pkg/bit"
)

const (
	FirstFlag  = 1
	SecondFlag = 2
)

func TestBitmap(t *testing.T) {
	var bm bit.Map
	bm |= SecondFlag
	if bm.IsSet(FirstFlag) {
		t.Fail()
	}
	if !bm.IsSet(SecondFlag) {
		t.Fail()
	}
}
