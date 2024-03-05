package main

import (
	"fmt"
	"sync"
)

func add(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Addition: ", a+b)
}

func subtract(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Subraction: ", a-b)
}

func multiply(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Multiplication: ", a*b)
}

func divide(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Division: ", a/b)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go add(1, 2, &wg)
	go subtract(1, 2, &wg)
	go multiply(1, 2, &wg)
	go divide(1, 2, &wg)
	wg.Wait()
}
