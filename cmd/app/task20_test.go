package main

import "testing"

func TestTask29(t *testing.T) {
	if reverseWords("snow dog sun") != "sun dog snow" {
		t.Fail()
	}
}