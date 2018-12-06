package main

import (
	"./lib"
	"fmt"
)

func main() {
	lines := lib.GetItems("inputs/2.txt")
	for i, line := range lines  {
		for j := i + 1; j < len(lines) ; j++ {
			// for each letter in those 2 words
			failedOnce := false
			failed := false
			failedOn := 0
			for k := range line {
				if line[k] != lines[j][k] {
					if !failedOnce {
						failedOnce = true
						failedOn = k
					} else {
						failed = true
						break
					}
				}
			}
			if !failed{
				fmt.Printf("%s%s\n", line[0:failedOn], line[failedOn+1:])
				return
			}
		}
	}
}