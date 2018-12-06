package main

import (
	"fmt"
	"./lib"
	"strconv"
)

func main() {
	lines := lib.GetItems("inputs/1.txt")
	sum := 0
	var frequencies []int
	for ; ;  {
		for _, line := range lines  {
			value, err := strconv.Atoi(line)
			if err != nil {
				continue
			}
			sum += value
			if lib.Member(frequencies, sum){
				fmt.Printf("%d\n", sum)
				return
			}
			frequencies = append(frequencies, sum)
		}
	}
}