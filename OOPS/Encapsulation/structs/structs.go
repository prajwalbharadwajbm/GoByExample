package structs

import "errors"

type Person struct {
	firstName string // private data members starts with lower case in golang
	lastName  string // These are accessable inside the same package, but for other package to access it we need getter and setter methods
	age       int
}

// Another way to create a person object is by constructor. Here in golang, a constructor would be a method/reciver function of struct.
func NewPerson(FirstName, LastName string, Age int) Person {
	// Better to add validation of each field in constructor to have better code.
	return Person{
		firstName: FirstName,
		lastName:  LastName,
		age:       Age,
	}
}

// Always use a pointer based reciever function when trying to alter an object's property. As you need address to that object for altering the value.
func (p *Person) SetFirstName(firstName string) error { // setter functions must always start with set(as part of gopher convention).
	if len(firstName) < 3 {
		return errors.New("INVALID FIRSTNAME")
	}
	p.firstName = firstName
	return nil
}

func (p Person) FirstName() (string, error) { // getter functions should start with capital letter as per gopher convention.
	if len(p.firstName) < 3 {
		return "", errors.New("INVALID FIRSTNAME inside Firstname")
	}
	return p.firstName, nil
}
