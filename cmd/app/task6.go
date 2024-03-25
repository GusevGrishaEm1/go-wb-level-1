package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func task6() {
	ch := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		timer := time.NewTimer(2 * time.Second)
		defer wg.Done()
		select {
		case <- timer.C:
			fmt.Print("Поток1 закончен")
		case <- ch:
			fmt.Print("Поток1 остановлен")
		}
	}()

	time.Sleep(time.Second)
	close(ch)

	wg.Wait()

	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		timer := time.NewTimer(2 * time.Second)
		defer wg.Done()
		select {
		case <- timer.C:
			fmt.Print("Поток2 закончен")
		case <- ctx.Done():
			fmt.Print("Поток2 остановлен")
		}
	}()

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
}
