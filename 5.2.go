package main

import (
	"./lib"
	"fmt"
	"strings"
	"math"
)


func alpha() (string, string) {
	lower := make([]byte, 26)
	upper := make([]byte, 26)
	for i := range lower {
		lower[i] = 'a' + byte(i)
		upper[i] = 'A' + byte(i)
	}
	return string(lower), string(upper)
}

func reduce(line string) string {
	for {
		changed := false
		size := len(line) - 1
		for i := 0 ; i < size ; i++ {
			left := line[i : i + 1]
			right := line[i + 1 : i + 2]
			leftLower := strings.ToLower(left)
			rightLower := strings.ToLower(right)
			if left != right && leftLower == rightLower {
				line = line[:i] + line[i + 2:]
				changed = true
				break
			}
		}
		if !changed {
			return line
		}
	}
}


func main() {
	line := lib.GetItems("inputs/5.txt")[0]
	line = reduce(line)
	lower, upper := alpha()

	aggregate := lib.Aggregate{math.MaxInt32, 0, 0}

	for i := range lower {
		woLower := strings.Replace(line, lower[i : i + 1], "", -1)
		woUpper := strings.Replace(woLower, upper[i : i + 1], "", -1)
		aggregate.Min(len(reduce(woUpper)), i)
	}
	fmt.Printf("%d\n", aggregate.Result)
}