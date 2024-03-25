package main

import (
	"reflect"
	"testing"
)

func TestTask23(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := []int{1, 3}
	if !reflect.DeepEqual(expected, removeElement(arr, 1)) {
		t.Fail()
	}
}