package substring_test

import (
	"testing"

	"github.com/mishamyrt/checode/v1/pkg/substring"
)

type Trimmer func(text string) string

type Subsequenter func(delimeter string, text string) string
type Previouser func(delimeter string, text string) string
type Midster func(startDel string, endDel string, text string) string

func TrimSuite(t *testing.T, trim Trimmer) {
	res := trim(" * Test 	")
	if res != "Test" {
		t.Errorf("Wrong trim: \"%s\"", res)
		t.Fail()
	}
}

func SubsequentSuite(t *testing.T, sub Subsequenter) {
	res := sub(":", "Junk :Test")
	if res != "Test" {
		t.Errorf("Wrong subsequent: \"%s\"", res)
		t.Fail()
	}
}

func PreviousSuite(t *testing.T, prev Previouser) {
	res := prev(":", "Test: Junk")
	if res != "Test" {
		t.Errorf("Wrong previous: \"%s\"", res)
		t.Fail()
	}
}

func MidstSuite(t *testing.T, midst Midster) {
	res := midst("<", ">", "junk <Test> junk")
	if res != "Test" {
		t.Errorf("Wrong midst: \"%s\"", res)
		t.Fail()
	}
}

func TestTrim(t *testing.T) {
	TrimSuite(t, substring.Trim)
}

func TestSubsequent(t *testing.T) {
	SubsequentSuite(t, substring.GetSubsequent)
}

func TestPrevious(t *testing.T) {
	PreviousSuite(t, substring.GetPrevious)
}

func TestMidst(t *testing.T) {
	MidstSuite(t, substring.GetMidst)
}
