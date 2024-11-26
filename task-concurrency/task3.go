package main

import (
	"fmt"
	"time"
)

func produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)
	go produce(ch)
	go consume(ch)
	time.Sleep(1 * time.Second)

}
