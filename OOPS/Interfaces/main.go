package main

import (
	"fmt"
	"reflect"
)

type printer interface {
	productInfo()
}

type Laptop struct {
	name  string
	price int
}
type Mobile struct {
	name  string
	price int
}

func (l Laptop) productInfo() {
	fmt.Println("Name : ", l.name)
	fmt.Println("Price : ", l.price)
}

func (m Mobile) productInfo() {
	fmt.Println("Name : ", m.name)
	fmt.Println("Price : ", m.price)
}

func (m Mobile) discount() { //*unused*, just to show interfaces allows to have other functions that what it enforces.
	fmt.Println("Discount : ", m.price-1)
}

func main() {
	var printInterface printer = Laptop{"Macbook pro", 2000}

	val, ok := printInterface.(Laptop) // Type Asserstion to get concrete type and data from an interface variable.
	if ok {
		fmt.Println("Laptop exists, and here are its concrete value accessed by interface : ", val)
		fmt.Println(reflect.TypeOf(val)) //main.Laptop
	} else {
		fmt.Println("printInterface does not hold Laptop")
	}

	iphone := Mobile{"Iphone", 1000}
	macbook := Laptop{"Macbook pro", 2000}

	//Also we use interfaces to enforce fucntions(but we dont explicitly define them like in other language, but if a struct implements functions of an interface then it automatically implements the interface implicitly )
	//To Demo this we will create a slice of interface
	info := []printer{iphone, macbook} //Lets comment down the function productinfo() for Mobile struct, which would then not adhere to interface and thus throw an error
	//Uncomment the productinfo() code binded with mobile struct to run this code.

	// info[0].productInfo()
	info[1].productInfo()

	// l.productInfo()
	// fmt.Println(reflect.TypeOf(l)) //main.Laptop
	// printInterface.productInfo()
	// fmt.Println(reflect.TypeOf(printInterface)) //main.Laptop
}
