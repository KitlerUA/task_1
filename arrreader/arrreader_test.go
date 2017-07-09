// arrreader_test
package arrreader

import (
	"testing"
)

var arrayExample []int
var nExample = 0

func TestReadArray(t *testing.T) {
	arr, n := ReadArray()
	if arr != nil {
		t.Error("Fail reading array\n Expect", arrayExample, " got ", arr)
	}
	if n != nExample {
		t.Error("Wrong array size\n Expect", nExample, " got ", n)
	}
}
