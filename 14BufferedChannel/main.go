package main

import (
	"fmt"
	"sync"
)

func add(a, b int, wg *sync.WaitGroup, result chan<- string) {
	defer wg.Done()
	result <- fmt.Sprintln("Addition:", a+b)
}

func subtract(a, b int, wg *sync.WaitGroup, result chan<- string) {
	defer wg.Done()
	result <- fmt.Sprintln("Subtraction:", a-b)
}

func multiply(a, b int, wg *sync.WaitGroup, result chan<- string) {
	defer wg.Done()
	result <- fmt.Sprintln("Multiplication:", a*b)
}

func divide(a, b int, wg *sync.WaitGroup, result chan<- string) {
	defer wg.Done()
	result <- fmt.Sprintln("Division:", a/b)
}
func main() {
	var wg sync.WaitGroup
	channel := make(chan string, 4)
	wg.Add(4)
	go add(1, 2, &wg, channel)
	go subtract(1, 2, &wg, channel)
	go multiply(1, 2, &wg, channel)
	go divide(1, 2, &wg, channel)
	wg.Wait()

	go func() {
		wg.Wait()
		close(channel)
	}()
	for result := range channel {
		fmt.Println(result)
	}
}
