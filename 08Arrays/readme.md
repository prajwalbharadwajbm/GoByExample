# **Arrays**
Prajwal Bharadwaj BM
7th July 2023

In Go, an array is just a numbered sequence of elements of a specfic length. Typically we find slices much more common than an array, in go.Due to their fixed length array are not much popular like Slice in Go language. 

Arrays are useful in **some** special scenarios. 

In an array, you are allowed to store zero or more than zero elements in it.

The type `[n]T` is an array of `n` values of type `T`.

The expression
```go
var a [10]int
```
declares a variable `a` as an array of ten `integer`

An array's length is part of its type, **so arrays cannot be resized**. Does feel limiting, but go offers a very convininet way of working with arrays

Example for **"array's length is part of its type"**:
```go
var a [10]int
fmt.Println(reflect.TypeOf(a))
var b [20]int
fmt.Println(reflect.TypeOf(b))
```
Output:
```bash
[10]int
[20]int
```
---

In golang, arrays can be created in 2 ways:
- Using var keyword : Given `type`, `size` and `variable name`
  ```go 
  var array_name [length]Type
  ```
- Using walrus operator: 
  ```go
  array_name := [length]Type
  ```
---
In Golang arrays are mutable, so we can set value for a given index i.e. We can set a value at an index using the `array[index] = value` syntax, and get a value with `array[index]`.
  ```go
  var array_name[index] = element/value
  ```

The builtin `len` returns the length of an array.
  ```go
  fmt.Println("length of array a is:", len[a])
  ```

Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.
 
In go, Array types are one-dimensional, but you can compose types to build multi-dimensional data structures. Its not like java having true multidimensional array. Its more of an array inside an array.
