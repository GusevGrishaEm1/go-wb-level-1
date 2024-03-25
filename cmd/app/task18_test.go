package main

import (
	"sync"
	"testing"
)

func TestTask17(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(200)
	c1 := &CounterV1{}
	for i := 0; i < 100; i++ {
		go func() {
			c1.Inc()
			defer wg.Done()
		}()
	}
	c2 := &CounterV2{}
	for i := 0; i < 100; i++ {
		go func() {
			c2.Inc()
			defer wg.Done()
		}()
	}
	wg.Wait()
	if c1.val != 100 {
		t.Fail()
	}
	if c2.val.Load() != 100 {
		t.Fail()
	}
}