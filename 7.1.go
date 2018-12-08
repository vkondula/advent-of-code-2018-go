package main

import (
	"fmt"
	"./lib"
	"regexp"
	"sort"
)

type Graph struct {
	dependencies map[int][]int
	steps []int
	Done []int
}

func (graph *Graph) Add (before int, after int) {
	value, _ := graph.dependencies[after]
	value = append(value, before)
	graph.dependencies[after] = value
	if !lib.Member(graph.steps, after){
		graph.steps = append(graph.steps, after)
	}
	if !lib.Member(graph.steps, before){
		graph.steps = append(graph.steps, before)
	}
}

func (graph *Graph) Check (step int) bool {
	dependencies, ok := graph.dependencies[step]
	if !ok { return true }
	for _, dependency := range dependencies {
		if !lib.Member(graph.Done, dependency) { return false }
	}
	return true
}

func (graph *Graph) Process () {
	sort.Ints(graph.steps)
	for len(graph.steps) > 0{
		changed := false
		for index, key := range graph.steps {
			if graph.Check(key) {
				graph.Done = append(graph.Done, key)
				graph.steps = append(graph.steps[:index], graph.steps[index + 1:]...)
				changed = true
				break
			}
		}
		if !changed { panic("No solution!") }
	}
}


func main() {
	lines := lib.GetItems("inputs/7.txt")
	graph := Graph{make(map[int][]int), []int{}, []int{}}
	re, _ := regexp.Compile(`Step (.) .* step (.).*`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		graph.Add(int(matches[1][0]), int(matches[2][0]))
	}
	graph.Process()
	var steps = make([]uint8, len(graph.Done))
	for i, value := range graph.Done { steps[i] = uint8(value)}
	fmt.Printf("%s\n", string(steps))
}