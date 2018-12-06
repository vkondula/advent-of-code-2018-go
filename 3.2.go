package main

import (
	"./lib"
	"regexp"
	"strconv"
	"fmt"
)

const SIZE = 1000

type Field struct{
	taken       bool
	overlapping bool
	value 		int
}

func (field *Field) aquire(value int) (wasTaken bool) {
	field.value = value
	if field.taken{
		field.overlapping = true
		field.value = 0
	}
	wasTaken = field.taken
	field.taken = true
	return
}

func parse(re *regexp.Regexp, line string) (value int, x int, y int, x_delta int, y_delta int){
	matched := re.FindStringSubmatch(line)
	value, _ = strconv.Atoi(matched[1])
	x, _ = strconv.Atoi(matched[2])
	y, _ = strconv.Atoi(matched[3])
	x_delta, _ = strconv.Atoi(matched[4])
	y_delta, _ = strconv.Atoi(matched[5])
	return
}

func overlapping(matrix [SIZE][SIZE]Field, re *regexp.Regexp, line string) int {
	value, x, y, x_delta, y_delta := parse(re, line)
	for x_i := x; x_i < (x + x_delta); x_i++ {
		for y_i := y; y_i < (y + y_delta); y_i++ {
			if value != matrix[x_i][y_i].value {
				return 0
			}
		}
	}
	return value
}

func main() {
	lines := lib.GetItems("inputs/3.txt")
	var matrix [SIZE][SIZE]Field
	re, _ := regexp.Compile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	for _, line := range lines{
		value, x, y, x_delta, y_delta := parse(re, line)
		for x_i := x; x_i < (x + x_delta); x_i++ {
			for y_i := y; y_i < (y + y_delta); y_i++ {
				matrix[x_i][y_i].aquire(value)
			}
		}
	}
	for _, line := range lines{
		value := overlapping(matrix, re, line)
		if value == 0 {
			continue
		}
		fmt.Printf("%d\n", value)
		return
	}
}