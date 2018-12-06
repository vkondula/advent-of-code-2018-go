package main

import (
	"./lib"
	"regexp"
	"strconv"
	"fmt"
)

type Field struct{
	taken       bool
	overlapping bool
}

func (field *Field) aquire() {
	if field.taken{
		field.overlapping = true
	}
	field.taken = true
}

func main() {
	lines := lib.GetItems("inputs/3.txt")
	const size = 1000
	var matrix [size][size]Field
	re, _ := regexp.Compile(`#\d+ @ (\d+),(\d+): (\d+)x(\d+)`)
	for _, line := range lines{
		matched := re.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matched[1])
		y, _ := strconv.Atoi(matched[2])
		x_delta, _ := strconv.Atoi(matched[3])
		y_delta, _ := strconv.Atoi(matched[4])
		for x_i := x; x_i < (x + x_delta); x_i++ {
			for y_i := y; y_i < (y + y_delta); y_i++ {
				matrix[x_i][y_i].aquire()
			}
		}
	}
	counter := 0
	for _, line := range matrix {
		for _, row := range line {
			if row.overlapping {
				counter++
			}
		}
	}
	fmt.Printf("%d\n", counter)
}