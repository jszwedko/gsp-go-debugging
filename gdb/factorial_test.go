package math2_test

import (
	"testing"

	"."
)

type testpair struct {
	value     int64
	factorial int64
}

var tests = []testpair{
	{5, 120},
	{0, 1},
}

func TestFactorial(t *testing.T) {
	for _, pair := range tests {
		factorial := math2.Factorial(pair.value)
		if factorial != pair.factorial {
			t.Error(
				"for", pair.value,
				"expected", pair.factorial,
				"got", factorial,
			)
		}
	}
}
