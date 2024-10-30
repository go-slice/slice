package slice_test

import (
	"math/rand"
	"reflect"
	"runtime"
	"testing"

	"github.com/go-slice/slice"
)

func Test(t *testing.T) {
	data := make([]int, 0, 11)

	s := slice.FromRaw(data)             // []
	s.Push(2, 3, 3)                      // [2 3 3]
	s.DeleteOne(2)                       // [2 3]
	s.Unshift(1)                         // [1 2 3]
	s.Unshift(0)                         // [0 1 2 3]
	s.Shift()                            // [1 2 3]
	s.Push(4, 4)                         // [1 2 3 4 4]
	s.Pop()                              // [1 2 3 4]
	s.Push(20)                           // [1 2 3 4 20]
	s.Filter(func(_ int, val int) bool { // [1 2 3 4]
		return val != 20
	})
	s.Push(5, 6, 7, 8, 9, 10) // [1 2 3 4 5 6 7 8 9 10]

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, result := range [][]int{data[1:11], s.Raw()} {
		if !reflect.DeepEqual(expected, result) {
			t.Error("expected != result", result)
		}
	}
}

func TestUnderlyingSlice(t *testing.T) {
	requireEqual := func(expected, given any) {
		if !reflect.DeepEqual(expected, given) {
			t.Helper()
			t.Errorf("expected != given: %v != %v", expected, given)
			t.FailNow()
		}
	}

	data := make([]int, 0, 100)
	s := slice.FromRaw(data[:0])

	s.Push(2, 3)
	requireEqual([]int{2, 3}, s.Raw())
	requireEqual([]int{2, 3}, data[:2])

	s.Unshift(1)
	requireEqual([]int{1, 2, 3}, s.Raw())
	requireEqual([]int{1, 2, 3}, data[:3])

	s.Shift()
	requireEqual([]int{2, 3}, s.Raw())
	requireEqual([]int{0, 2, 3}, data[:3])

	s.Pop()
	requireEqual([]int{2}, s.Raw())
	requireEqual([]int{0, 2, 0}, data[:3])

	s.Unshift(1)
	s.Push(3, 4, 5)
	requireEqual([]int{1, 2, 3, 4, 5}, s.Raw())
	requireEqual([]int{0, 1, 2, 3, 4, 5}, data[:6])

	s.Delete(3, 2)
	requireEqual([]int{1, 2, 3}, s.Raw())
	requireEqual([]int{0, 1, 2, 3, 0, 0}, data[:6])

	s.Push(4, 5)
	requireEqual([]int{1, 2, 3, 4, 5}, s.Raw())
	requireEqual([]int{0, 1, 2, 3, 4, 5}, data[:6])

	s.Filter(func(_ int, val int) bool {
		return val > 3
	})
	requireEqual([]int{4, 5}, s.Raw())
	requireEqual([]int{0, 4, 5, 0, 0, 0}, data[:6])

	s.Insert(0, 1, 2, 3)
	requireEqual([]int{1, 2, 3, 4, 5}, s.Raw())
	requireEqual([]int{0, 1, 2, 3, 4, 5, 0, 0, 0}, data[:9])

	data[:2][1] = 99
	requireEqual([]int{99, 2, 3, 4, 5}, s.Raw())
	requireEqual([]int{0, 99, 2, 3, 4, 5, 0, 0, 0}, data[:9])

	s.Shift()
	s.Unshift(1)
	s.Pop()
	s.Pop()
	s.Push(6, 7)
	requireEqual([]int{1, 2, 3, 6, 7}, s.Raw())
	requireEqual([]int{0, 0, 1, 2, 3, 6, 7, 0}, data[:8])

	s.Replace(3, 4, 5)
	requireEqual([]int{1, 2, 3, 4, 5}, s.Raw())
	requireEqual([]int{0, 0, 1, 2, 3, 4, 5, 0}, data[:8])

}

func heapAlloc() uint64 {
	runtime.GC()

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	return m.HeapAlloc
}

func TestSlice_Clone(t *testing.T) {
	s := slice.FromRaw(make([]int64, 0, 1024*1024))
	before := heapAlloc()
	s = s.Clone()
	after := heapAlloc()

	var expected float64 = 1024 * 1024 * 8 * .9

	if float64(before-after) < expected {
		t.Error("Memory usage has not decreased")
	}

	_ = s
}

func TestSlice_Shuffle(t *testing.T) {
	in := slice.FromRaw([]int{1, 2, 3, 4, 5})
	notExpected := []int{1, 2, 3, 4, 5}

	differs := false
	for i := 0; i < 10; i++ {
		in.Shuffle(rand.Intn)
		if !reflect.DeepEqual(notExpected, in.Raw()) {
			differs = true

			break
		}
	}

	if !differs {
		t.Error("Shuffle does not work")
	}
}
