package main

import "fmt"

func oddOrEven(x int64) string {
	if x%2 == 0 {
		return fmt.Sprintf("%d is even", x)
	} else {
		return fmt.Sprintf("%d is odd", x)
	}
}

func isDivisible(x int64, y int64) string {
	if x%y == 0 {
		return fmt.Sprintf("%d is divisible by %d", x, y)
	}
	return fmt.Sprintf("%d is not divisible by %d", x, y)
}

func shortHandIf() string {
	if num := 9; num < 0 {
		return fmt.Sprintln(num, "is negative")
	} else if num < 10 {
		return fmt.Sprintln(num, "has 1 digit") //num can only be accessible by else and else if
	} else {
		return fmt.Sprintln(num, "has multiple digits")
	}
}

func main() {
	oddOrEvenResult := oddOrEven(7)
	isDivisibleResult := isDivisible(8, 4)
	shortHandIfResult := shortHandIf()

	fmt.Println(oddOrEvenResult)
	fmt.Println(isDivisibleResult)
	fmt.Println(shortHandIfResult)
}
