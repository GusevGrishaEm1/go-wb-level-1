package main

import (
	"reflect"
	"testing"
)

func TestTask15(t *testing.T) {
	arr := []int{8, 5, 3, 10, 4}
	quickSort(arr)
	expected := []int{3, 4, 5, 8, 10}
	if !reflect.DeepEqual(expected, arr) {
		t.Fail()
	}
}