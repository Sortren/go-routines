package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	wordsList := os.Args[1:]
	words := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		for _, word := range wordsList {
			words <- word
			time.Sleep(time.Second)
		}

		close(words)
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case word, ok := <-words:
				if !ok {
					return
				}
				fmt.Println(word)
			}
		}
	}()

	wg.Wait()
}
