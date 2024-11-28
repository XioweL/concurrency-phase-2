package main

import "fmt"

func main() {
	evenChan := make(chan int)
	oddChan := make(chan int)
	doneChan := make(chan bool)

	// Goroutine untuk mengirim numbers ke channel
	go func() {
		for i := 1; i <= 20; i++ {
			if i%2 == 0 {
				evenChan <- i
			} else {
				oddChan <- i
			}
		}
		//close channel ketika selesai
		close(evenChan)
		close(oddChan)
		doneChan <- true
	}()

	// Reading dari channels menggunakan select
	go func() {
		for {
			select {
			case even, ok := <-evenChan:
				if ok {
					fmt.Printf("Received an even number\t: %v\n", even)
				}
			case odd, ok := <-oddChan:
				if ok {
					fmt.Printf("Received an odd number\t: %v\n", odd)
				}
			case <-doneChan:
				fmt.Println("Done")
				return
			}
		}
	}()
	// Menunggu goroutine selesai.
	<-doneChan
}
