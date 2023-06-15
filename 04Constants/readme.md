# **Constants**
Prajwal Bharadwaj BM
14th June 2023

## Introduction
Go is statically typed language,  which means that the types of variables and expressions are known at compile time.This helps to prevent errors and makes code more readable. However, Go also has a feature called type inference, which allows the compiler to infer the type of a variable or expression from its context. So it doesn't permit operations that mix numeric types. You can't declare int to float64, or even an int to int32. 

While desiging go, they decided not mix up numeric types like they did with other programming languages such as C, Java etc

```sh
unsigned int u = 1e9;
long signed int i = -1;
... i + u ...
```

By above example written in C, How would a unseasoned programmer know lenght of the result? What type the value would be? Is it signed or unsigned?

In Go if we want to add i and u, then we must be explicit about what you want the result to be.
```sh
var u uint  
var i int 
```
### Solution:
either `uint(i)+u` or `i+int(u)` ,with both the meaning and type of the addition clearly expressed, but unlike in C you cannot write i+u. You can’t even mix int and int32, even when int is a 32-bit type.

This strictness eliminates a common cause of bugs and other failures. It is a vital property of Go. But it has a cost: it sometimes requires programmers to *decorate their code with clumsy numeric conversions to express their meaning clearly*.

What about constants? Given the declarations above, what would make it legal to write `i` `=` `0` or `u` `=` `0`? What is the *type* of `0`? It would be unreasonable to require constants to have type conversions in simple contexts such as `i` `=` `int(0)`.

Then they realized that numeric constants should work differently from how they behave in other C-like languages. After much thinking and experimentation, we came up with a design that we believe feels right almost always. This design frees the programmer from converting constants all the time yet allows them to write things like `math.Sqrt(2)` without being warned by the compiler.

In short, constants in Go just work, most of the time. Let’s see how this happens.

## String Constant
A string constant encloses some text between double quotes. (Go also has raw string literals, enclosed by backquotes ``, but for the purpose of this discussion they have all the same properties.) Here is a string constant:
```sh
"Hello, World!"
```
According to our knowledge we feel it of type string or note it as a string constant, **but its not!!!!**

It is a untyped string constant, which means it is just a constant textual type that does not have any fixed type.

Yes, it’s a string, but it’s not a Go value of type string. It remains an untyped string constant even when given a name:
```sh
const hello = "Hello, World"
```
A typed constant would look like this:
```sh
const typedHello string = "Hello, World"
```

Lets explore more by considering an example program

```sh
const typedHello string = "Hello, World"
var s string //define variable s of type string
s = typedHello 
fmt.Println(s) //Prints Hello, World

//But this does not print, as it throws a type error.

type prajwal's_string string
var m prajwal's_string
m = typedHello // Type error
fmt.Println(m)
```

The variable m has type prajwal's_string and cannot be assigned a value of a different type. It can only be assigned values of type prajwal's_string, like this:

```sh
const helloString prajwal's_string = "Hello, World"
m = helloString
fmt.Println(m)
```
or by forcing the issue with a conversion, like this:
```sh
m = prajwal's_string(typedHello)
fmt.Println(m)
```
Returning to our untyped string constant, it has the helpful property that, since it has no type, assigning it to a typed variable does not cause a type error. That is, we can write

`m = "Hello, World" `
or `m = hello`

because, unlike the typed constants typedHello and helloString, the untyped constants "Hello, World" and hello have no type. Assigning them to a variable of any type compatible with strings works without error.

These untyped string constants are strings, of course, so they can only be used where a string is allowed, but they do not have type string.


**Please refer docs for further info*