package main

import "fmt"

func main() {
	fmt.Println("Go" + "-" + "Lang")

	fmt.Println("2+2 =", 2+2)
	fmt.Println("Println (6.28318530718/2.0 =)", 6.28318530718/2.0)
	fmt.Printf("Printf (6.28318530718/2.0 = %.2f)\n", 6.28318530718/2.0)

	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
