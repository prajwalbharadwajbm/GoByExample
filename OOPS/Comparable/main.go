package main

import (
	"fmt"

	"github.com/prajwalbharadwajbm/OOPS/Comparable/comparable"
)

func main() {
	p := comparable.NewPerson("Prajwal", "Bharadwaj", 23, []string{"a"})
	p1 := comparable.NewPerson("Prajwal", "Bharadwaj", 23, []string{"b"}) //When primitive properties exists we can directly compare to structs.

	// But when non primitive properties such as slices, or anything that is stored in dynamic memory. We have add a function in package
	// comparable to check equality of that non-primitive data type, such as by checking lenght etc

	// if p == p1 {
	// 	fmt.Println("Same")
	// }

	if p.Equals(p1) {
		fmt.Println("Same")
	}
}
