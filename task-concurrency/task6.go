package main

import "fmt"

func main() {
	evenChan := make(chan int)
	oddChan := make(chan int)
	errorChan := make(chan int)
	doneChan := make(chan bool)

	// Goroutine untuk mengirim numbers ke channel
	go func() {
		for i := 1; i <= 25; i++ {
			if i > 20 {
				errorChan <- i
			} else if i%2 == 0 {
				evenChan <- i
			} else {
				oddChan <- i
			}
		}
		//close channel ketika selesai
		close(evenChan)
		close(oddChan)
		close(errorChan)
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
			case error, ok := <-errorChan:
				if ok {
					fmt.Printf("Error: number %v is greater than 20\n", error)
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
