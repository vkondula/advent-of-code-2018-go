package main

import (
	"./lib"
	"fmt"
	"strings"
)


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
	fmt.Printf("%d\n", len(line))
}