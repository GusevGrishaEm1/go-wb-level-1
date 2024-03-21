package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	task4(5)	
}

func task4(workersCount int) {
	ch := make(chan string)
	endCh := make(chan os.Signal, 1)
	signal.Notify(endCh, syscall.SIGINT)
	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
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
			fmt.Println(word)
			ch <- word
		}
	}

	close(ch)

	wg.Wait()
}



func generateRandomWord() string {
    letters := "abcdefghijklmnopqrstuvwxyz"
    word := make([]byte, 5)
    for i := range word {
        word[i] = letters[rand.Intn(len(letters))]
    }
    return string(word)
}
