package polymorphism

import "fmt"

// In go we use interface to declare a method signature under it.
// And any struct making use of that method will implement it, thus attaining polymorphism.
type Shape interface {
	WhichShape()
}

type Circle struct{}

func (c Circle) WhichShape() {
	fmt.Println("I am a circle")
}

type Triangle struct{}

func (t Triangle) WhichShape() {
	fmt.Println("I am a triangle")
}
