package main

import (
	"fmt"
	"time"
)

func task5(N int) {
	ch := make(chan string)
	timer := time.NewTimer(time.Duration(N) * time.Second)
	go func() {
		for {
			ch <- generateRandomWord()
			time.Sleep(2 * time.Second)
		}
	}()
	go func() {
		for {
			word := <-ch
			fmt.Print(word + "\n")
		}
	}()
	<- timer.C
	close(ch)
}