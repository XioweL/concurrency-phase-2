package main

import (
	"fmt"
	"time"
)

func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
func PrintLetters() {
	for i := 'a'; i <= 'j'; i++ {
		fmt.Println(string(i))
	}
}

func main() {
	go PrintNumbers()
	go PrintLetters()
	time.Sleep(2 * time.Second)
}
