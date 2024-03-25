package main

import "testing"

func TestTask19(t *testing.T) {
	if reverseString("hello") != "olleh" {
		t.Fail()
	}
}