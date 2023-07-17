# **Slices**
Prajwal Bharadwaj BM
17th July 2023

Slices are an important data type in go, giving more powerful interface to seqences in golang than array.

The type `[]T` is a slice with elements of type `T`.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:
```go 
a[low : high]
```
This selects a half-open range which `includes` the `first element`, but `excludes` the `last one`.

```go
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s) //[3,5,7]
}
```

A slice does not store any data, it just describes a section of an underlying array.

Changing an element of slice, does change the corresponding element of its underlying array. 

Other slices that share the same underlying array will also have those changes.

```go
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) //[John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) //[John Paul][Paul George]

	b[0] = "XXX"
	fmt.Println(a, b) //[John XXX][XXX George]
	fmt.Println(names) //[John XXX George Ringo]
}
```
---
## Slice Literals
A slice literal is like an array literal without the length

An array literal would look like
```go
[3]bool{true, true, false}
```

A String Literal would look like
```go
[]bool{true, true, false}
```

## Slice Defaults
In slice, we can omit high and low bounds value for defaults.

Default is zero for `low bound` and length of the slice for `high bound`.
```go
//Example:
//For an array
var a [10]int

//Below are equivalent slice expressions:
a[0:10]
a[:10]
a[0:]
a[:]
```
## Slice length and capacity
A slice has both a `length` and `capacity`.

Length of a slice : Number of elements a slice contains.(What it stores at the moment)

Capacity of a slice : Number of elements in its underlying array.(How much it can store)

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

```go
//slice len() and cap()
	s := []int{1, 2, 3, 4, 5, 6}
	printSlice(s)

	s = s[:0] //Slicing the slice to give length to be zero
	printSlice(s)

	s = s[:4] //Extending slice to length 4
	printSlice(s)

	s = s[2:] //Dropping 2 elements from the slice
	//so lower bound drops the elements and thus taking down the capacity of the slice.
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}
```
Output:
```
len = 6 cap = 6 [1 2 3 4 5 6]
len = 0 cap = 6 []
len = 4 cap = 6 [1 2 3 4]
len = 2 cap = 4 [3 4]
```

## Nil slices

The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.

```go
var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
```
Output
```
[] 0 0
nil!
```