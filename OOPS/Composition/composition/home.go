package composition

type room struct {
	noOfRoom int
	roomName string // instead of just "name" be specific
}

// Imagine having name property in both room and door structs, and we write a getter function for each as follows

// func (r *room) Name() string{
// 	return r.name
// }

//	func (d *door) Name() string{
//		return d.name
//	}
//
// This would create an ambuguity for main function while calling <Home'sObject>.name(). Thus always write function specific to struct under the main struct Home.
// Example
func (h *Home) DoorName() string {
	return h.doorName
}
func (h *Home) RoomName() string {
	return h.roomName
}

// getter function for room
func (r *room) NoOfRoom() int {
	return r.noOfRoom
}

type door struct {
	doorDimension int
	doorName      string // instead of just "name" be specific
}

// getter function for door
func (d *door) DoorDimension() int {
	return d.doorDimension
}

// This is called a composition. Struct inside a struct. In golang we dont have concept of inheritance
// and thus can't acheive inheritance in go. So we use composition instead.
type Home struct {
	HomeName string
	room
	door
}

// Lets create a constructor to understand how we embed structs inside structs to achieve composition
func NewHome(homeName string, noOfRoom int, roomName string, doorDimesion int, doorName string) Home {
	return Home{
		HomeName: homeName,
		room: room{
			noOfRoom: noOfRoom,
			roomName: roomName,
		},
		door: door{
			doorDimension: doorDimesion,
			doorName:      doorName,
		},
	}
}
