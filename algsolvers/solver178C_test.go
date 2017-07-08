// solver178C_test
package algsolvers

import (
	"testing"
)

var tests178C = []testpair{
	{[]int{1, 2, 3}, 0},
}

func TestSolver178C(t *testing.T) {
	for _, pair := range tests178C {
		res := Solver178C(pair.values)
		if res != pair.result {
			t.Error("For ", pair.values, " expected ", pair.result, " got ", res)
		}
	}
}
