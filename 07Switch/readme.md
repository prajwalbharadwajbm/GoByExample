# **Switch Statement**
Prajwal Bharadwaj BM
20th June 2023


A `switch` statement is a shorter way to write a sequence of `if - else` statements. It runs the first case whose value is equal to the `condition expression`.

Go only runs the selected case, not all the cases that follow, like other programming languages which does run other cases.

In Java, the switch statement operates differently compared to Go. when a specific case is matched in a switch statement, the code execution continues from that point and executes all the subsequent cases until it reaches a break statement or the end of the switch block.

Let's consider an example in Java to illustrate this behavior:

```java
public class SwitchExample {
    public static void main(String[] args) {
        int number = 2;

        switch (number) {
            case 1:
                System.out.println("Case 1");
            case 2:
                System.out.println("Case 2");
            case 3:
                System.out.println("Case 3");
                break;
            case 4:
                System.out.println("Case 4");
                break;
            default:
                System.out.println("Default case");
                break;
        }
    }
}
```
Output in Java:
```sh
Case 2
Case 3
```
In this Java example, since number is 2, the code execution enters the switch block and matches the case 2. After executing the code block for case 2, the execution continues to the next case, which is 3, and executes its code block as well. It only breaks out of the switch block when encountering the break statement explicitly or reaching the end of the switch block.

Now let's consider the equivalent example in Go:

```go
package main

import "fmt"

func main() {
    number := 2

    switch number {
    case 1:
        fmt.Println("Case 1")
    case 2:
        fmt.Println("Case 2")
    case 3:
        fmt.Println("Case 3")
    case 4:
        fmt.Println("Case 4")
    default:
        fmt.Println("Default case")
    }
}
```
Output in Go:
```sh
Case 2
```

In Go, when the case 2 is matched, it executes the code block associated with that case and then immediately exits the switch statement. It doesn't continue to execute the subsequent cases as it does in Java. Therefore, in Go, only the code block for the matched case is executed, resulting in a single output line.

### **Another important difference is that Go's switch cases **need not be constants**, and the values involved need not be integers.*

```go
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

(For example,
```go
switch i {
case 0:
case f():
}
```
does not call f if i==0.)


---
Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else chains.

```go
func main() {
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```
---

