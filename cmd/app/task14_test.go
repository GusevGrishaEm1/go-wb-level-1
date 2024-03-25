package main

import "testing"

func TestTask14(t *testing.T) {
	tests := []interface{} {"str", 45, true, make(chan int)}
	for _, el := range tests {
		if task14(el) == "unknown" {
			t.Fail()
		}
	}
}