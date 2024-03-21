package main

import (
	"fmt"
	"sync"
)

func task2() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(arr))

	for _, el := range arr {
		go func (el int) {
			defer wg.Done()
			ch <- el * el
		} (el)
	}

	go func() {
		wg.Wait()
		defer close(ch)
	}()

	for el := range ch {
		fmt.Print(el, " ")
	}
}