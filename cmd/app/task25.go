package main

import "time"

func sleep(ms int) {
	<-time.After(time.Duration(ms) * time.Millisecond)
}