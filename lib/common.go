package lib

import (
	"os"
	"bufio"
)

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func GetItems(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic("Couldn't open given file!")
	}
	var lines []string
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Member(sequence []int, key int) bool {
	for _, item := range sequence {
		if item == key {
			return true
		}
	}
	return false
}

func Max(vector []int) (maxIndex int, max int) {
	for key, value := range vector {
		if value > max {
			max = value
			maxIndex = key
		}
	}
	return
}


type Aggregate struct {
	Result int
	Count int
	Key int
}
func (aggregate *Aggregate) Min(value int, key int) {
	aggregate.Count++
	if value < aggregate.Result {
		aggregate.Result = value
		aggregate.Key = key
	}
}
func (aggregate *Aggregate) Max(value int, key int) {
	aggregate.Count++
	if value > aggregate.Result {
		aggregate.Result = value
		aggregate.Key = key
	}
}
func (aggregate *Aggregate) Sum(value int) {
	aggregate.Count++
	aggregate.Result += value
}