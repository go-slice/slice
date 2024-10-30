package slice

// Slice is a wrapper of any slice that allows performing basic operations over slices using an intuitive syntax.
//
//	var s slice.Slice[int]
//	s.Push(4, 5, 6)    // [4 5 6]
//	s.Unshift(1, 2, 3) // [1 2 3 4 5 6]
//	s.Pop()            // [1 2 3 4 5]
type Slice[T any] []T

// FromRaw creates a new [Slice].
func FromRaw[T any](in []T) Slice[T] {
	return in
}

// Raw returns the underlying slice.
func (s *Slice[T]) Raw() []T {
	return *s
}

// Empty returns false when the len of the given slice equals 0.
func (s *Slice[T]) Empty() bool {
	return len(*s) == 0
}

// Shift returns the first element and removes it from the given slice.
func (s *Slice[T]) Shift() (_ T, ok bool) {
	var r T

	if s.Empty() {
		return r, false
	}

	start := *s
	defer clear(start[0:1])

	r, *s = (*s)[0], (*s)[1:]

	return r, true
}

// Unshift prepends the given input to the given slice.
func (s *Slice[T]) Unshift(v ...T) {
	if len(v) == 0 {
		return
	}

	if *s == nil {
		*s = make(Slice[T], len(v))
		copy(*s, v)

		return
	}

	*s = append(*s, v...)

	copy(
		(*s)[len(v):],
		(*s)[:len(*s)-len(v)],
	)
	copy(
		(*s)[:len(v)],
		v,
	)
}

// Pop returns the last element and removes it from the given slice.
func (s *Slice[T]) Pop() (_ T, ok bool) {
	var r T

	if s.Empty() {
		return r, false
	}

	start := *s
	defer clear(start[len(start)-1:])

	r, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]

	return r, true
}

// Push appends the given input to the given slice.
func (s *Slice[T]) Push(v ...T) {
	if *s == nil {
		*s = make([]T, 0, len(v))
	}

	*s = append(*s, v...)
}

// DeleteOne deletes a single element from the given slice.
//
//	s.DeleteOne(index) // it's an equivalent of s.Delete(index, 1)
func (s *Slice[T]) DeleteOne(index int) (ok bool) {
	return s.Delete(index, 1)
}

// Delete deletes a vector of the given length under the given index from the given slice.
//
//	s := slice.FromRaw([]int{1, 2, 3, 4, 5})
//	s.Delete(1, 3)
//	fmt.Println(s) [1 5]
func (s *Slice[T]) Delete(index int, length int) (ok bool) {
	if index < 0 || *s == nil || index+length > len(*s) {
		return false
	}

	start := *s
	defer clear(start[len(start)-length:])

	*s = append((*s)[:index], (*s)[index+length:]...)

	return true
}

// Insert inserts the given element to the existing slice under the given index.
//
//	s := slice.FromRaw([]string{"one", "four"})
//	s.Insert(1 "two", "three")
//	fmt.Println(s) // ["one", "two", "three", "four"]
func (s *Slice[T]) Insert(index int, v ...T) (ok bool) {
	if index < 0 || index > len(*s) {
		return false
	}

	*s = append(*s, v...)
	copy(
		(*s)[index+len(v):],
		(*s)[index:index+len(v)],
	)
	copy(
		(*s)[index:index+len(v)],
		v,
	)

	return true
}

// Replace replaces s[index:index+len(v)] with v.
// It does not succeed when the given slice does not have enough elements to be replaced.
//
//	s := slice.FromRaw([]string{"one", "two", "two", "two", "two"})
//	s.Replace(2, "three", "four", "five")
//	fmt.Println(s) // [one two three four five]
func (s *Slice[T]) Replace(index int, v ...T) (ok bool) {
	if index < 0 || index+len(v) > len(*s) {
		return false
	}

	copy((*s)[index:index+len(v)], v)

	return true
}

// Clone returns a new slice with the same length and copies to it all the elements from the existing slice.
func (s *Slice[T]) Clone() Slice[T] {
	if *s == nil {
		return nil
	}

	r := make(Slice[T], s.Len())
	copy(r, *s)

	return r
}

// Cap returns the capacity of the given slice.
func (s *Slice[T]) Cap() int {
	return cap(*s)
}

// Len returns the length of the given slice.
func (s *Slice[T]) Len() int {
	return len(*s)
}

// Filter filters the given slice using the provided func.
func (s *Slice[T]) Filter(keep func(index int, val T) bool) {
	if *s == nil {
		return
	}

	n := (*s)[:0]

	for i, x := range *s {
		if keep(i, x) {
			n = append(n, x)
		}
	}

	clear((*s)[len(n):])

	*s = n
}

// Get returns an element under the given index.
// It accepts negative indexes.
//
//	x := slice.FromRaw([]int{1, 2, 3, 4, 5})
//	val, ok := x.Get(-1)
//	fmt.Println(val) // 5
func (s *Slice[T]) Get(index int) (_ T, ok bool) {
	if index < 0 {
		index += len(*s)
	}

	if index < 0 || index >= len(*s) {
		var r T

		return r, false
	}

	return (*s)[index], true
}

// Reverse reverses order of the given slice.
func (s *Slice[T]) Reverse() {
	for i := len(*s)/2 - 1; i >= 0; i-- {
		j := len(*s) - 1 - i
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

// Sort sorts the given slice in ascending order as determined by the cmp function.
// It requires that cmp is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.
func (s *Slice[T]) Sort(cmp func(a T, b T) int) {
	sort(*s, cmp)
}

// Shuffle shuffles the given input.
// randIntN must generate a pseudo-random number in the half-open interval [0,n).
//
//	s := slice.FromRaw([]int{1, 2, 3, 4, 5})
//	s.Shuffle(rand.Intn)
func (s *Slice[T]) Shuffle(randIntN func(n int) int) {
	for i := len(*s) - 1; i > 0; i-- {
		j := randIntN(i + 1)
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
