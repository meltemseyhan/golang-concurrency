package main

import (
	"fmt"
	"sync"
	"time"
)

func mainn() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		msg, ok := <-ch
		fmt.Println(msg, ok)
		if !ok {
			fmt.Println("Channel is closed")
		}
		// Channel close olana kadar döner
		for msg := range ch {
			fmt.Println(msg)
		}

		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch <- i
		}
		// Buffer size 1 olduğu bunu da için yazabiliriz
		ch <- 10
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
