package main

import "testing"

func TestTask1(t *testing.T) {
	act := &Action{}
	t.Run("test1", func(t *testing.T) {
		if act.do() != "do" {
			t.Fail()
		}
	})
}