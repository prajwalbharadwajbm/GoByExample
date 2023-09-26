package main

import "fmt"

// So sometimes, we have cases where we may have to wait for a subprocess(a go routine) to complete before we call a function.
// We can do that by sleeping the main process, until the go routine gets completed. But how long can we sleep, how do we have to calculate those?
// Thus we have deffered Function Calls.
// Call can be deferred until the surrounding function completes
// Typically used for cleanup activities, as it basically excecutes at the end of the scope.
func main() {
	defer fmt.Println("bye")

	fmt.Println("Hello")

	deferCall()
}

// One thing to keep in mind is that the arguments in a deferred call are not evaluated in a deferred way,
// but it is evaluated immediately but only the call is deferred.
// This wouldn't mean anything when sending a fixed argument, which can't change or doesn't need evaluation.
// But if i send an argument, which is to be evaluated. then we have to note that argument is evaluated right
// there where the deferred statment is encountered, but not later when the call actually happens.
func deferCall() {
	i := 1
	//The args get evaluated here, so it prints 2.
	defer fmt.Println(i + 1)
	// hence even though the value of i would be 3, its printed 2.
	i++
	fmt.Println("Hello inside defer")
}
