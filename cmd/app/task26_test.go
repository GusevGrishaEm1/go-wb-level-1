package main

import "testing"

func TestTask26(t *testing.T) {
	if checkUnique("abCdefAaf") != false {
		t.Fail()
	}
	if checkUnique("Aab") != true {
		t.Fail()
	}
}