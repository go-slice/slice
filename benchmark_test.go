package slice_test

import (
	"testing"

	"github.com/go-slice/slice"
)

func BenchmarkSlice_Unshift(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s := slice.FromRaw(make([]int, 5, 10))
		s.Unshift(1, 2)
	}
}

func BenchmarkSlice_Unshift_native(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s := make([]int, 5, 10)
		s = append(append([]int(nil), 1, 2), s...) //nolint:ineffassign,staticcheck
	}
}

func BenchmarkSlice_Filter(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s := slice.FromRaw([]int{1, 2, 3, 4, 5})
		s.Filter(func(_ int, val int) bool {
			return val%2 == 0
		})
	}
}
