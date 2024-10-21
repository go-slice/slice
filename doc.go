// Package slice provides a toolset to manipulate slices.
//
//	var s slice.Slice[int]
//	s.Push(4, 5, 6)    // [4 5 6]
//	s.Unshift(1, 2, 3) // [1 2 3 4 5 6]
//	s.Pop()            // [1 2 3 4 5]
package slice
