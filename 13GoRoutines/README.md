# Understanding Goroutines and WaitGroups in Go

## Introduction

In this example, lets explore the concepts of Goroutines and WaitGroups in Go programming language. Goroutines are lightweight threads managed by the Go runtime, allowing concurrent execution of functions. 

WaitGroups are used to wait for a collection of Goroutines to finish their execution.

## Code Overview

The provided Go program demonstrates the use of Goroutines and WaitGroups to perform basic arithmetic operations concurrently.

### Goroutines:
- Goroutines are concurrent threads of execution in Go.
- They are created using the `go` keyword followed by the function call.

### WaitGroups:
- WaitGroups are used to wait for the completion of a collection of Goroutines.
- The `sync.WaitGroup` type is used to manage the synchronization.
- The `Add` method increments the WaitGroup counter, and `Done` decrements it.
- The `Wait` method blocks until the counter becomes zero.

### Concurrency:
- The main function launches four Goroutines concurrently to perform addition, subtraction, multiplication, and division.
- Each Goroutine increments the WaitGroup counter and defers `Done` to decrement the counter when finished.

### Defer Statement:
- The `defer` statement is used to ensure that the `Done` method is called even if the Goroutine encounters an error or panics.

### Output:
- Due to the concurrent execution, the order of output may vary, but the waitgroup from sync package ensures that all Goroutines finish before the main function exits.
