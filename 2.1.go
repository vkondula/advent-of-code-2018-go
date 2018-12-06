package main

import (
	"./lib"
	"fmt"
)

func getMap(word string) map[int32]int{
	var data map[int32]int
	data = make(map[int32]int)
	for _, letter := range word{
		i, ok := data[letter]
		if !ok {
			i = 0
		}
		data[letter] = i + 1
	}
	return data
}


func main() {
	lines := lib.GetItems("inputs/2.txt")
	twos := 0
	threes := 0
	for _, line := range lines  {
		wordData := getMap(line)
		twoFound := false
		threeFound := false
		for _, value := range wordData{
			if !twoFound && value == 2 {
				twoFound = true
				twos++
			}
			if !threeFound && value == 3 {
				threeFound = true
				threes++
			}
		}
	}
	fmt.Printf("%d\n", twos * threes)
}