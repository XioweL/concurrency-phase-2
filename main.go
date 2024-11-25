package main

import (
	"REVIEW-P1/task-concurrency"
	"time"
)

func main() {
	go task_concurrency.PrintNumbers()
	time.Sleep(1 * time.Second)
	go task_concurrency.PrintLetters()
	time.Sleep(1 * time.Second)
}
