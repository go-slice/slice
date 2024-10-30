[![Go Reference](https://pkg.go.dev/badge/github.com/go-slice/slice.svg)](https://pkg.go.dev/github.com/go-slice/slice)
[![Tests](https://github.com/go-slice/slice/actions/workflows/tests.yml/badge.svg)](https://github.com/go-slice/slice/actions/workflows/tests.yml)
[![Coverage Status](https://coveralls.io/repos/github/go-slice/slice/badge.svg?branch=main)](https://coveralls.io/github/go-slice/slice?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-slice/slice)](https://goreportcard.com/report/github.com/go-slice/slice)

# slice

A simple, generic, and easy-to-use Go library that provides a rich set of utility functions for managing and manipulating slices. It wraps native slices and offers an intuitive API for performing common operations such as adding, removing, filtering, and modifying elements.

## Features

- **Simple and Intuitive Syntax**: Easy-to-use methods for operations like push, pop, and insertion.
- **Flexible Types**: Works with any type of slice thanks to Go’s generics.
- **Advanced Operations**: Support for advanced functionality like filtering, sorting, shuffling, and reversing.
- **Safe Indexing**: Support for negative indexes and index-bounds checks.

## Performance Characteristics

- **No Extra Allocations**: This library is designed to avoid unnecessary memory allocations. Operations like shifting, popping, deleting, or inserting elements reuse the existing slice capacity wherever possible. This helps optimize memory usage and keeps performance consistent, especially in tight loops or high-frequency operations.
- **Memory Zeroing**: When elements are deleted, shifted, or popped, the library automatically "zeroes out" the removed elements by setting them to the zero value of their type. This prevents potential memory leaks and ensures that the garbage collector can efficiently clean up unused elements.


## Installation

```bash
go get github.com/go-slice/slice
```

## Usage Examples

### Creating a Slice

There are multiple ways to create a slice, either with or without predefined elements:

#### Using an Empty Slice

To create an empty slice of integers:

```go
package main

import (
    "fmt"

    "github.com/go-slice/slice"
)

func main() {
    // Create an empty slice
    var s slice.Slice[int]

    // Add some elements
    s.Push(4, 5, 6)
    
    fmt.Println(s)  // Output: [4 5 6]
}
```

#### Initializing with Elements

You can also initialize the slice with elements using the following syntax:

```go
// Create a slice with predefined elements
s := slice.Slice[int]{4, 5, 6}
fmt.Println(s)  // Output: [4 5 6]
```

#### Creating with Predefined Capacity

To create a slice with a predefined capacity (which can help avoid redundant allocations when you know the size in advance):

```go
// Create a slice with a capacity of 100 but no initial elements
s := slice.FromRaw(make([]int, 0, 100))

// Now you can add elements without triggering new allocations until the slice exceeds the capacity
s.Push(1, 2, 3)
fmt.Println(s)  // Output: [1 2 3]
```


### Basic Operations

#### Push and Pop
Append or remove elements from the end of the slice.

```go
var s slice.Slice[int]

// Push 3 elements to the slice
s.Push(10, 20, 30)
fmt.Println(s)  // Output: [10 20 30]

// Pop the last element
val, _ := s.Pop()
fmt.Println(val)  // Output: 30
fmt.Println(s)    // Output: [10 20]
```

#### Unshift and Shift
Add or remove elements from the beginning of the slice.

```go
var s slice.Slice[int]

// Push elements to the slice
s.Push(3, 4, 5)

// Unshift adds elements to the beginning of the slice
s.Unshift(1, 2)
fmt.Println(s)  // Output: [1 2 3 4 5]

// Shift removes the first element and returns it
val, _ := s.Shift()
fmt.Println(val)  // Output: 1
fmt.Println(s)    // Output: [2 3 4 5]
```

### Deleting Elements

#### Delete a Single Element

You can delete a single element from any position in the slice. For example, deleting the element at index `2`:

```go
var s slice.Slice[int]

// Push elements to the slice
s.Push(1, 2, 3, 4, 5)

// Delete 1 element starting at index 2
s.Delete(2, 1)
fmt.Println(s)  // Output: [1 2 4 5]
```

#### Delete the Last Element

You can use a negative index to delete the last element. This can be done using `s.Delete(-1, 1)` or `s.DeleteOne(-1)`, as both are equivalent when removing the last element:

```go
var s slice.Slice[int]

// Push elements to the slice
s.Push(1, 2, 3, 4, 5)

// Delete 1 element starting from the last index (-1)
s.Delete(-1, 1)
fmt.Println(s)  // Output: [1 2 3 4]

// Alternatively, delete the last element with DeleteOne
s.DeleteOne(-1)
```

### Modifying the Slice

#### Insert
Insert elements at a specific index.

```go
var s slice.Slice[string]

// Push initial elements to the slice
s.Push("one", "three")

// Insert "two" at index 1
s.Insert(1, "two")
fmt.Println(s)  // Output: [one two three]
```

#### Replace
Replace a section of the slice with new elements.

```go
var s slice.Slice[string]

// Push elements to the slice
s.Push("a", "b", "b", "b", "e")

// Replace the slice from index 2 with "c" and "d"
s.Replace(2, "c", "d")
fmt.Println(s)  // Output: [a b c d e]
```

### Advanced Operations

#### Filter
Remove elements based on a custom condition.

```go
var s slice.Slice[int]

// Push elements to the slice
s.Push(1, 2, 3, 4, 5, 6)

// Keep only even numbers using the Filter method
s.Filter(func(_ int, v int) bool {
    return v%2 == 0
})
fmt.Println(s)  // Output: [2 4 6]
```

#### Reverse
Reverse the order of elements in the slice.

```go
var s slice.Slice[int]

// Push elements to the slice
s.Push(1, 2, 3, 4, 5)

// Reverse the order of the slice
s.Reverse()
fmt.Println(s)  // Output: [5 4 3 2 1]
```

#### Sort
Sort the slice using a custom comparison function.

```go
var s slice.Slice[int]

// Push elements to the slice
s.Push(5, 3, 1, 4, 2)

// Sort the slice in ascending order
s.Sort(func(a, b int) int {
    return a - b
})
fmt.Println(s)  // Output: [1 2 3 4 5]
```

#### Shuffle
Randomly shuffle the elements in the slice.

```go
import "math/rand"

var s slice.Slice[int]

// Push elements to the slice
s.Push(1, 2, 3, 4, 5)

// Shuffle the slice using a random integer generator
s.Shuffle(rand.Intn)
fmt.Println(s)  // Output: [3 5 1 4 2] (shuffled randomly)
```

### Safe Indexing

#### Get with Negative Indexes
Get an element using a negative index, counting from the end.

```go
var s slice.Slice[int]

// Push elements to the slice
s.Push(1, 2, 3, 4, 5)

// Get the last element using a negative index (-1)
val, _ := s.Get(-1)
fmt.Println(val)  // Output: 5
```

### Other Utility Methods

- `Clone()`: Create a copy of the slice.
- `Len()`: Get the length of the slice.
- `Cap()`: Get the capacity of the slice.
- `Empty()`: Check if the slice is empty.

## License

This project is licensed under the MIT License.

---

Feel free to use, modify, and contribute!
