[![Go Reference](https://pkg.go.dev/badge/github.com/go-slice/slice.svg)](https://pkg.go.dev/github.com/go-slice/slice)
[![Tests](https://github.com/go-slice/slice/actions/workflows/tests.yml/badge.svg)](https://github.com/go-slice/slice/actions/workflows/tests.yml)
[![Coverage Status](https://coveralls.io/repos/github/go-slice/slice/badge.svg?branch=main)](https://coveralls.io/github/go-slice/slice?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-slice/slice)](https://goreportcard.com/report/github.com/go-slice/slice)

# slice

A simple, generic, and easy-to-use Go library that provides a rich set of utility functions for managing and manipulating slices.
It wraps native slices and offers an intuitive API for performing common operations such as adding, removing, filtering, and modifying elements.

## Features

- **Simple and Intuitive Syntax**: Easy-to-use methods for operations like push, pop, and insertion.
- **Flexible Types**: Works with any type of slice thanks to Go’s generics.
- **Advanced Operations**: Support for advanced functionality like filtering, sorting, shuffling, and reversing.
- **Safe Indexing**: Support for negative indexes and index-bounds checks.

## Installation

```bash
go get github.com/go-slice/slice
```

## Usage Examples

### Creating a Slice

To create a slice, either start with an empty slice or use the `FromRaw` function to wrap an existing slice:

```go
package main

import (
    "fmt"
	
    "github.com/go-slice/slice"
)

func main() {
    // Create an empty slice of integers
    var s slice.Slice[int]
    
    // Create a slice from an existing raw slice
    s2 := slice.FromRaw([]int{1, 2, 3, 4})
    
    fmt.Println(s2)  // Output: [1 2 3 4]
}
```

You can also initialize a slice using the following syntax:

```go
s := slice.Slice[int]{}
```

### Basic Operations

#### Push and Pop
Append or remove elements from the end of the slice.

```go
var s slice.Slice[int]
s.Push(10, 20, 30)
fmt.Println(s)  // Output: [10 20 30]

val, _ := s.Pop()
fmt.Println(val)  // Output: 30
fmt.Println(s)    // Output: [10 20]
```

#### Unshift and Shift
Add or remove elements from the beginning of the slice.

```go
var s slice.Slice[int]
s.Push(3, 4, 5)
s.Unshift(1, 2)
fmt.Println(s)  // Output: [1 2 3 4 5]

val, _ := s.Shift()
fmt.Println(val)  // Output: 1
fmt.Println(s)    // Output: [2 3 4 5]
```

### Deleting Elements

#### Delete a Single Element

You can delete a single element from any position in the slice. For example, deleting the element at index `2`:

```go
var s slice.Slice[int]
s.Push(1, 2, 3, 4, 5)
s.Delete(2, 1)
fmt.Println(s)  // Output: [1 2 4 5]
```

#### Delete the Last Element

You can use a negative index to delete the last element. This can be done using `s.Delete(-1, 1)` or `s.DeleteOne(-1)`, as both are equivalent when removing the last element:

```go
var s slice.Slice[int]
s.Push(1, 2, 3, 4, 5)
s.Delete(-1, 1) // Alternatively: s.DeleteOne(-1)
fmt.Println(s)  // Output: [1 2 3 4]
```

### Modifying the Slice

#### Insert
Insert elements at a specific index.

```go
var s slice.Slice[string]
s.Push("one", "three")
s.Insert(1, "two")
fmt.Println(s)  // Output: ["one", "two", "three"]
```

#### Replace
Replace a section of the slice with new elements.

```go
var s slice.Slice[string]
s.Push("a", "b", "b", "b", "e")
s.Replace(2, "c", "d")
fmt.Println(s)  // Output: ["a", "b", "c", "d", "e"]
```

### Advanced Operations

#### Filter
Remove elements based on a custom condition.

```go
var s slice.Slice[int]
s.Push(1, 2, 3, 4, 5, 6)
s.Filter(func(_ int, v int) bool {
    return v%2 == 0  // Keep only even numbers
})
fmt.Println(s)  // Output: [2 4 6]
```

#### Reverse
Reverse the order of elements in the slice.

```go
var s slice.Slice[int]
s.Push(1, 2, 3, 4, 5)
s.Reverse()
fmt.Println(s)  // Output: [5 4 3 2 1]
```

#### Sort
Sort the slice using a custom comparison function.

```go
var s slice.Slice[int]
s.Push(5, 3, 1, 4, 2)
s.Sort(func(a, b int) int {
    return a - b  // Sort in ascending order
})
fmt.Println(s)  // Output: [1 2 3 4 5]
```

#### Shuffle
Randomly shuffle the elements in the slice.

```go
import "math/rand"

var s slice.Slice[int]
s.Push(1, 2, 3, 4, 5)
s.Shuffle(rand.Intn)
fmt.Println(s)  // Output: [3 5 1 4 2] (shuffled randomly)
```

### Safe Indexing

#### Get with Negative Indexes
Get an element using a negative index, counting from the end.

```go
var s slice.Slice[int]
s.Push(1, 2, 3, 4, 5)
val, _ := s.Get(-1)
fmt.Println(val)  // Output: 5
```

### Other Utility Methods

- `Clone()`: Create a copy of the slice.
- `Len()`: Get the length of the slice.
- `Cap()`: Get the capacity of the slice.

## License

This project is licensed under the MIT License.

---

Feel free to use, modify, and contribute!
