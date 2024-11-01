package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumberGenerator(ch chan int) {
	for {
		num := rand.Intn(100)
		ch <- num
	}
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	ch := make(chan int)

	go RandomNumberGenerator(ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
