//go:build go1.21
// +build go1.21

package slice

import (
	"slices"
)

func sort[T any](s []T, cmp func(a T, b T) int) {
	slices.SortStableFunc(s, cmp)
}
