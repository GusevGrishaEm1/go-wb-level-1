package main

import (
	"testing"
)

func TestTask8(t *testing.T) {
	t.Run("test3", func(t *testing.T) {
		if setBit(200, 2) != 204 {
			t.Fail()
		}
		if unsetBit(200, 3) != 192 {
			t.Fail()
		}
	})
}