package main

import (
	"fmt"
	"sync"
)

func task7() {
	var m sync.Mutex
	counter := 1
	mp := make(map[int]string)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			m.Lock()
			mp[counter] = generateRandomWord()
			counter++
			m.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			m.Lock()
			mp[counter] = generateRandomWord()
			counter++
			m.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			m.Lock()
			mp[counter] = generateRandomWord()
			counter++
			m.Unlock()
		}
	}()
	wg.Wait()
	for key, val := range mp {
		fmt.Print(key, "-", val, "\n")
	}
}