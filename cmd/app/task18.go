package main

import (
	"sync"
	"sync/atomic"
)

type CounterV1 struct {
    mu    sync.Mutex
    val int
}

type CounterV2 struct {
    val atomic.Int64
}

func (c *CounterV1) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.val++
}

func (c *CounterV2) Inc() {
    c.val.Add(1)
}