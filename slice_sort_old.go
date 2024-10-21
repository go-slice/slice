//go:build !go1.21
// +build !go1.21

package slice

import (
	pkgSort "sort"
)

func sort[T any](s []T, cmp func(a T, b T) int) {
	pkgSort.SliceStable(s, func(i, j int) bool {
		return cmp(s[i], s[j]) < 0
	})
}
