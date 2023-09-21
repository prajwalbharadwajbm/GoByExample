package main

import (
	"fmt"

	"github.com/prajwalbharadwajbm/OOPS/Encapsulation/structs"
)

func main() {
	p := structs.Person{} //declared Person object. But cannot define directly as we can't access the data members which are private.

	//so we use setter function
	err := p.SetFirstName("Pr")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//and getter function
	firstName, err := p.FirstName()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(firstName)
}
