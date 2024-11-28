package main

import (
	"fmt"
	"time"
)

func producetask4(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumetask4(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 5)
	go producetask4(ch)
	go consumetask4(ch)
	time.Sleep(1 * time.Second)

	fmt.Println(cap(ch))
	fmt.Println(len(ch))
	/* Jika tidak menggunakan buffer, saat data dikirim ke channel
	dia akan otomatis menunggu sampai ada yang mengambil. Jika menggunakan
	buffer data langsung otomatis masuk ke buffer tanpa harus menunggu
	ada yang menerima data, kecuali data nya sudah habis, baru diminta menunggu.
	*/
}
