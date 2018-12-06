package main

import (
	"fmt"
	L "./lib"
	"strconv"
)

func main() {
	lines := L.GetItems("inputs/1.txt")
	sum := 0
	for _, line := range lines  {
		value, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		sum += value
	}
	fmt.Printf("%d\n", sum)
}