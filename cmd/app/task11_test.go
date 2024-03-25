package main

import (
	"reflect"
	"testing"
)

func TestTask11(t *testing.T) {
	set1 := []int{1, 2, 3}
	set2 := []int{5, 3, 2}
	expected := []int{3, 2}

	equal := reflect.DeepEqual(expected, task11(set1, set2))

	if !equal {
		t.Fail()
	}
}