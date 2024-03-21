package main

import "testing"

func TestTask3(t *testing.T) {
	t.Run("test3", func(t *testing.T) {
		if task3() != 220 {
			t.Fail()
		}
	})
}