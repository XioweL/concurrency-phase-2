package main

import (
	"fmt"
	"sync"
)

func PrintNumbersTask2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func PrintLettersTask2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'j'; i++ {
		fmt.Println(string(i))
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go PrintNumbersTask2(&wg)
	go PrintLettersTask2(&wg)
	wg.Wait()
	fmt.Println("All goroutines finished")
}
