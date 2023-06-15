# **Looping Construct**

*for* is Go’s only looping construct. Here are some basic types of for loops.

The basic for loop has three components separated by semicolons:

`
the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
`

```sh
func main(){
    sum := 0
    for i:1 ;i<10 ;i++{
        sum += 1
    }
    fmt.Println(sum)
}
```
The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.


The loop will stop iterating once the boolean condition evaluates to false.

> Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

The init and post statements are optional.

```sh
func main() {
	sum := 1
	for ; sum < 10; {
		sum += sum
	}
	fmt.Println(sum)
```

At that point you can drop the semicolons: `C's while is spelled for in Go.`

```sh
func main() {
	sum := 1
	for sum < 10 { //while syntax but with 'for' in go
		sum += sum
	}
	fmt.Println(sum)
}
```

To run an infinite loop:
```sh
func main() {
	for {
	}
}
```

