package math2

func Factorial(n int64) int64 {
	f := int64(1)

	for i := n; i > 1; i-- {
		f = f + i
	}

	return f
}
