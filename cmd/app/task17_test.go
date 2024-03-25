package main

import "testing"

func TestTask16(t *testing.T) {
	arr := []int{4, 5, 3, 12, 44, 9}
	if !binarySearch(arr, 44) {
		t.Fail()
	}
	if binarySearch(arr, 55) {
		t.Fail()
	}
}