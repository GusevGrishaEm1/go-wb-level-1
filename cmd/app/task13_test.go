package main

import (
	"testing"
)

func TestTask13(t *testing.T) {
	a := 2
	b := 3
	a, b = swap(a, b)
	if b != 2 {
		t.Fail()
	}
	if a != 3 {
		t.Fail()
	}
}