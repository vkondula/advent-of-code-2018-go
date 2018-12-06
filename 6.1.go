package main

import (
	"fmt"
	"./lib"
	"regexp"
	"strconv"
	"math"
)


const SIZE = 400


type Field struct {
	Distance int
	Value int
	X int
	Y int
}

func (field *Field) SetDistance(x, y, value int){
	distance := lib.Abs(field.X - x) + lib.Abs(field.Y - y)
	if distance < field.Distance {
		field.Distance = distance
		field.Value = value
	} else if distance == field.Distance {
		field.Distance = distance
		field.Value = -1
	}
}


func initFields() [SIZE][SIZE]Field {
	var fields [SIZE][SIZE]Field
	for j := 0 ; j < SIZE ; j++ {
		for k := 0 ; k < SIZE ; k++ {
			fields[j][k].X = j
			fields[j][k].Y = k
			fields[j][k].Value = -1
			fields[j][k].Distance = math.MaxInt32
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

	forbidden := make(map[int]bool)
	for i := 0 ; i < SIZE ; i++ {
		forbidden[fields[i][0].Value] = true
		forbidden[fields[0][i].Value] = true
		forbidden[fields[i][SIZE-1].Value] = true
		forbidden[fields[SIZE-1][i].Value] = true
	}

	var counter = make([]int,len(lines))
	for j := 0 ; j < SIZE ; j++ {
		for k := 0 ; k < SIZE ; k++ {
			value := fields[j][k].Value
			_, ok := forbidden[value]
			if ok || value == -1 {
				continue
			}
			counter[value]++
		}
	}

	_, max := lib.Max(counter)

	fmt.Printf("%d\n", max)
}