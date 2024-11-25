package task_concurrency

import "fmt"

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
