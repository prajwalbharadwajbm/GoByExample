package main

import (
	"fmt"

	"github.com/prajwalbharadwajbm/OOPS/Composition/composition"
)

func main() {
	c := composition.NewHome("Yagnavalkya", 1, "room1", 10, "door1")
	fmt.Println(c.DoorDimension()) //w/ich kinda similar to achieving inheritance rather than c.door.DoorDimension(),
	// which must have done if created a data member for each struct.

	fmt.Println(c.DoorName())
	fmt.Println(c.RoomName())
}
