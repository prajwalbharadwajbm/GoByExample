package main

import "fmt"

func main() {
	// Most basic single condition 'for' initialization
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	// A classic init/condition/statement
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	//A loop that runs until broken.
	//for without a condition will loop repeatedly until
	//you break out of the loop or return from the
	//enclosing function.
	for {
		fmt.Println("infinite without break")
		break
	}

	//continue keyword can be used to skip for a particular condition
	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
