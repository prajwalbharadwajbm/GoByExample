package main

import "fmt"

// Sometimes we need variable number of arguments to be passed to a function,
// then we use ellipses operator(or basically three dots in a row)
func getMax(vals ...int) int {
	maxVal := -1
	// Here we see that vals argument is treated as slices
	for _, v := range vals {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func main() {
	fmt.Println(getMax(1, 2, 3, 4, 5, 66, 7))
	//We can also pass a slice as an argument to func getMax
	slice := []int{1, 2, 3, 4, 5, 66, 7}
	//But notice we have added a suffix of three dots.
	fmt.Println(getMax(slice...))
}
