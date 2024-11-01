package main

import (
	"fmt"
	"sync"
)

func MergeChannels(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(ch <-chan int) {
		for val := range ch {
			out <- val
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 4; i <= 6; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 7; i <= 9; i++ {
			ch3 <- i
		}
	}()

	mergedChannel := MergeChannels(ch1, ch2, ch3)

	for val := range mergedChannel {
		fmt.Println(val)
	}
}
