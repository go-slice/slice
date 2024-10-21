package slice

func clear[T any, S ~[]T](in S) {
	var x T

	for i := 0; i < len(in); i++ {
		in[i] = x
	}
}
