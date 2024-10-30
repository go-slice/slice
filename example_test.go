package slice_test

import (
	"fmt"
	"math/rand"

	"github.com/go-slice/slice"
)

func Example() {
	var s slice.Slice[int]

	s.Push(4, 5, 6)
	s.Unshift(1, 2, 3)
	s.DeleteOne(0)
	s.Pop()

	fmt.Println(s)

	// Output: [2 3 4 5]
}

func ExampleSlice_Raw_nil() {
	s := slice.Slice[int](nil)
	fmt.Println(s.Raw() == nil)

	// Output: true
}

func ExampleSlice_Raw() {
	data := []int{1, 2, 3}
	s := slice.Slice[int](data)

	fmt.Println(s.Raw())
	// Output: [1 2 3]
}

func ExampleSlice_Empty() {
	s1 := slice.FromRaw(make([]int, 0))
	s2 := slice.Slice[int](nil)
	s3 := slice.Slice[int]([]int{1})

	s3.Pop()

	fmt.Println(s1.Empty())
	fmt.Println(s2.Empty())
	fmt.Println(s3.Empty())

	// Output:
	// true
	// true
	// true
}

func ExampleSlice_Shift() {
	s := slice.FromRaw([]int{1, 2, 3})
	fmt.Println(s.Shift())
	fmt.Println(s)

	// Output:
	// 1 true
	// [2 3]
}

func ExampleSlice_Shift_nil() {
	s := slice.Slice[int](nil)
	fmt.Println(s.Shift())

	// Output:
	// 0 false
}

func ExampleSlice_Unshift() {
	s := slice.FromRaw([]int{4, 5, 6})
	s.Unshift(2, 3)
	s.Unshift(1)
	s.Unshift() // do nothing

	fmt.Println(s)

	// Output: [1 2 3 4 5 6]
}

func ExampleSlice_Unshift_nil() {
	s := slice.FromRaw[int](nil)
	s.Unshift(1, 2, 3)

	fmt.Println(s)

	// Output: [1 2 3]
}

func ExampleSlice_Pop() {
	s := slice.FromRaw([]int{1, 2})
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	// Output:
	// 2 true
	// 1 true
	// 0 false
}

func ExampleSlice_Push() {
	s := slice.FromRaw([]int{1, 2})
	s.Push(3)
	s.Push(4, 5)
	fmt.Println(s)

	// Output:
	// [1 2 3 4 5]
}

func ExampleSlice_Push_nil() {
	s := slice.FromRaw[int](nil)
	s.Push(1, 2, 3)
	fmt.Println(s)

	// Output:
	// [1 2 3]
}

func ExampleSlice_DeleteOne() {
	s := slice.FromRaw([]string{"one", "two", "three"})
	fmt.Println(s.DeleteOne(3)) // no element under the index 3
	fmt.Println(s.DeleteOne(1))
	fmt.Println(s)

	// Output:
	// false
	// true
	// [one three]
}

func ExampleSlice_Delete() {
	s := slice.FromRaw([]string{"one", "two", "three", "four", "five", "six"})
	fmt.Println(s.Delete(0, 1)) // [two three four five six]
	fmt.Println(s.Delete(4, 1)) // [two three four five]
	fmt.Println(s.Delete(1, 2)) // [two five]
	fmt.Println(s.Delete(0, 3)) // [two five] - do nothing, invalid input
	fmt.Println(s)

	// Output:
	// true
	// true
	// true
	// false
	// [two five]
}

func ExampleSlice_Insert() {
	s := slice.FromRaw([]string{"one", "four"})
	s.Insert(1, "two", "three")
	fmt.Println(s)

	// Output:
	// [one two three four]
}

func ExampleSlice_Replace() {
	s := slice.FromRaw([]string{"one", "two", "two", "two", "two"})
	s.Replace(2, "three", "four", "five")
	fmt.Println(s)

	// Output:
	// [one two three four five]
}

func ExampleSlice_Replace_false() {
	s := slice.FromRaw([]string{"one", "two", "two", "two"})

	// cannot replace 3 elements starting from index 2
	// the given slice does not have enough elements
	fmt.Println(s.Replace(2, "three", "four", "five"))
	fmt.Println(s)

	// Output:
	// false
	// [one two two two]
}

func ExampleSlice_Insert_false() {
	s := slice.FromRaw[string](nil)
	fmt.Println(s.Insert(1, "one")) // the highest possible index to insert == len(s)

	s = slice.FromRaw([]string{"zero", "one", "two"})
	fmt.Println(s.Insert(-1, "minus one")) // index MUST be >= 0

	// Output:
	// false
	// false
}

func ExampleSlice_Clone() {
	original := make([]int, 2, 10)
	original[0] = 1
	original[1] = 2

	s := slice.FromRaw(original)

	clone := s.Clone()

	// modify the original slice
	s.DeleteOne(1)
	s.Push(5)

	fmt.Println(original)
	fmt.Println(clone) // clone remains unchanged

	// Output:
	// [1 5]
	// [1 2]
}

func ExampleSlice_Clone_nil() {
	s := slice.FromRaw[int](nil)
	fmt.Println(s.Clone() == nil)

	// Output: true
}

func ExampleSlice_Cap() {
	s1 := slice.Slice[int](nil)
	s2 := slice.FromRaw(make([]int, 0, 100))

	fmt.Println(s1.Cap())
	fmt.Println(s2.Cap())

	// Output:
	// 0
	// 100
}

func ExampleSlice_Len() {
	s1 := slice.Slice[int](nil)
	s2 := slice.FromRaw(make([]int, 100))

	fmt.Println(s1.Len())
	fmt.Println(s2.Len())

	// Output:
	// 0
	// 100
}

func ExampleSlice_Filter() {
	x := slice.FromRaw([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	x.Filter(func(_ int, val int) bool {
		return val%2 == 0
	})

	fmt.Println(x)

	// Output:
	// [2 4 6 8 10]
}

func ExampleSlice_Filter_nil() {
	x := slice.FromRaw[int](nil)
	x.Filter(func(index int, val int) bool { // do nothing
		return true
	})
	fmt.Println(x.Raw() == nil)

	// Output: true
}

func ExampleSlice_Get() {
	x := slice.FromRaw([]int{1, 2, 3, 4, 5})
	fmt.Println(x.Get(0))  // 1 true
	fmt.Println(x.Get(4))  // 5 true
	fmt.Println(x.Get(-1)) // 5 true
	fmt.Println(x.Get(-5)) // 1 true
	fmt.Println(x.Get(-6)) // 0 false
	fmt.Println(x.Get(5))  // 0 false

	// Output:
	// 1 true
	// 5 true
	// 5 true
	// 1 true
	// 0 false
	// 0 false
}

func ExampleSlice_Reverse() {
	s := slice.FromRaw([]int{5, 4, 3, 2, 1})
	s.Reverse()
	fmt.Println(s)

	// Output: [1 2 3 4 5]
}

func ExampleSlice_Sort() {
	s := slice.FromRaw([]int{3, 2, 5, 4, 1})
	s.Sort(func(a int, b int) int {
		return a - b
	})
	fmt.Println(s)

	// Output: [1 2 3 4 5]
}

func ExampleSlice_Sort_nil() {
	var s slice.Slice[int]
	s.Sort(func(a int, b int) int { // do nothing, there is nothing to sort
		return a - b
	})
	fmt.Println(s.Raw() == nil)

	// Output: true
}

func ExampleSlice_Shuffle() {
	s := slice.FromRaw([]int{1, 2, 3, 4, 5})
	s.Shuffle(rand.Intn)
	fmt.Println(s) // e.g. [3 1 5 4 2]
}
