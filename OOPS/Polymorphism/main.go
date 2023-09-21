package main

import "github.com/prajwalbharadwajbm/OOPS/Polymorphism/polymorphism"

func main() {
	var shape1 polymorphism.Shape = polymorphism.Circle{} // we achieve polymorphism by implementing a interace method in a struct.
	var shape2 polymorphism.Shape = polymorphism.Triangle{}
	shape1.WhichShape()
	shape2.WhichShape()
}
