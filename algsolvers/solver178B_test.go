// solver178B_test
package algsolvers

import (
	"testing"
)

type testpair struct {
	values []int
	result int
}

var tests = []testpair{
	{[]int{1, 2, 3}, 1},
}

func TestSolver178B(t *testing.T) {
	for _, pair := range tests {
		res := Solver178B(pair.values)
		if res != pair.result {
			t.Error("For ", pair.values, " expected ", pair.result, " got ", res)
		}
	}
}
