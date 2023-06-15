package main

import (
	"fmt"
	"math"
	"reflect"
)

const str string = "Constant"

func main() {
	fmt.Println(reflect.TypeOf(str))
	const n = 500000000
	const d = 3e20 / n

	fmt.Println(d) // prints a float value in exp form, to get it in int form, you can convert it to int64(not int32, don't even ask why!!).

	fmt.Println("Before explicit conversion")
	fmt.Println(reflect.TypeOf(d)) // Has type which is kinda contrasting to what https://gobyexample.com/constants says!!!

	a := int64(d)

	fmt.Println("After explicit conversion")
	fmt.Println(reflect.TypeOf(a)) // it gets converted to int64

	fmt.Println(math.Sin(n)) // make sure to give correct type that a method or func expects!!
}
