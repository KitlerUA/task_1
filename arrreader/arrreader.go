// ArrayReader
package arrreader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadArray() ([]int, int) {
	fmt.Print("Enter numbers: ")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	words := strings.Fields(line)
	var result []int
	var size int = 0
	for w := range words {
		var temp, _ = strconv.ParseInt(words[w], 10, 32)
		result = append(result, int(temp))
		size++
	}
	return result, size
}
