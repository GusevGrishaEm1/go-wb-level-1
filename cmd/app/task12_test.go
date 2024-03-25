package main

import (
	"reflect"
	"testing"
)

func TestTask12(t *testing.T) {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	expected := []string{"cat", "dog", "tree"}
	
	equal := reflect.DeepEqual(expected, task12(arr))

	if !equal {
		t.Fail()
	}
}