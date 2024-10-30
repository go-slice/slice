package slice

// The function [clear] has been introduced in GO 1.21.
// See https://tip.golang.org/doc/go1.21#language.
func clear[S ~[]T, T any](in S) {
	var x T

	for i := 0; i < len(in); i++ {
		in[i] = x
	}
}
