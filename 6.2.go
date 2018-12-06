package main

import (
	"fmt"
	"./lib"
	"regexp"
	"strconv"
)


const SIZE = 400
const MAX_DISTANCE = 10000


type Field struct {
	Distance int
	Value int
	X int
	Y int
}

func (field *Field) SetDistance(x, y, value int){
	field.Distance += lib.Abs(field.X - x) + lib.Abs(field.Y - y)
}


func initFields() [SIZE][SIZE]Field {
	var fields [SIZE][SIZE]Field
	for j := 0 ; j < SIZE ; j++ {
		for k := 0 ; k < SIZE ; k++ {
			fields[j][k].X = j
			fields[j][k].Y = k
			fields[j][k].Value = -1
			fields[j][k].Distance = 0
		}
	}
	return fields
}


func main() {
	lines := lib.GetItems("inputs/6.txt")
	re, _ := regexp.Compile(`(-?\d+), (-?\d+)`)
	fields := initFields()
	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		for j := 0 ; j < SIZE ; j++ {
			for k := 0 ; k < SIZE ; k++ {
				fields[j][k].SetDistance(x, y, i)
			}
		}
	}

	counter := 0
	for j := 0 ; j < SIZE ; j++ {
		for k := 0 ; k < SIZE ; k++ {
			if fields[j][k].Distance < MAX_DISTANCE {
				counter++
			}
		}
	}

	fmt.Printf("%d\n", counter)
}