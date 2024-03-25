package main

import (
	"fmt"
	"sync"
)

func task9(N int) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i <= N; i++ {
			ch1 <- i
		}
		defer close(ch1)
	}()
	go func() {
		for el := range ch1 {
			ch2 <- el * el
		}
		defer close(ch2)
	}()
	go func() {
		for el := range ch2 {
			fmt.Print(el, ", ")
		}
		defer wg.Done()
	}()
	wg.Wait()
}