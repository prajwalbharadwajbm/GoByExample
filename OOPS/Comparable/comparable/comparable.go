package comparable

type Person struct {
	firstName string
	lastName  string
	age       int
	slice     []string
}

func NewPerson(FirstName, LastName string, Age int, Slice []string) Person {
	return Person{
		firstName: FirstName,
		lastName:  LastName,
		age:       Age,
		slice:     Slice,
	}
}
func (p *Person) Equals(p1 Person) bool {
	return len(p.slice) == len(p1.slice)
}
