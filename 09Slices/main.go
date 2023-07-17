package main

import (
	"fmt"
)

func main() {
	a_slice := []bool{true, true, false}
	fmt.Println(a_slice)

	a_struct_slice := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
	}
	fmt.Println(a_struct_slice)

	//slice default expression
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a[:])   //Prints whole array
	fmt.Println(a[0:2]) //Prints elements from 0 to 1
	fmt.Println(a[2:])  //Prints elements from 2 to length of the array
	fmt.Println(a[:3])  //Prints elements from 0 to 2

	//slice len() and cap()
	s := []int{1, 2, 3, 4, 5, 6}
	printSlice(s)

	s = s[:0] //Slicing the slice to give length to be zero
	printSlice(s)

	s = s[:4] //Extending slice to length 4
	printSlice(s)

	s = s[2:] //Dropping 2 elements from the slice
	//so lower bound drops the elements and thus taking down the capacity of the slice.
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}
