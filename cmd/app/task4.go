package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func task4(N int) {
	ch := make(chan string)
	endCh := make(chan os.Signal, 1)
	signal.Notify(endCh, syscall.SIGINT)
	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for el := range ch {
				fmt.Print("worker", i, " message: ", el, "\n")
			}
		}()
	}

	loop:
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case <-endCh:
			fmt.Println("Вы завершили программу")
			break loop
		default:
			word := generateRandomWord()
			ch <- word
		}
	}

	close(ch)

	wg.Wait()
}