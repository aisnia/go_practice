package strev

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if Reverse("123") != "321" {
		t.Log("123 string reverse must be 321")
		t.Fail()
	}
	if Reverse("") != "" {
		t.Log(" '' string reverse must be '' ")
		t.Fail()
	}
}
