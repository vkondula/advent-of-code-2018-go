package main

import (
	"fmt"
	"./lib"
	"regexp"
	"sort"
)

var WORKERS = 5
var TIME = 61

type Event struct {
	end int
	letter int
}


type EventManager struct {
	workers []*Event
	graph Graph
	time int
}

func (manager *EventManager) Assign (worker int, letter int) {
	end := manager.time + letter - 'A' + TIME
	event := Event{end, letter}
	manager.workers[worker] = &event
}

func (manager *EventManager) Clear () {
	for index, worker := range manager.workers {
		if worker != nil && worker.end == manager.time {
			manager.graph.Done = append(manager.graph.Done, worker.letter)
			manager.workers[index] = nil
		}
	}
}

func (manager *EventManager) GetWorker () (int, bool) {
	for index, worker := range manager.workers {
		if worker == nil {
			return index, true
		}
	}
	return 0, false
}

func (manager *EventManager) Working () bool {
	for _, worker := range manager.workers {
		if worker != nil {
			return true
		}
	}
	return false
}

func (manager *EventManager) Process () int {
	sort.Ints(manager.graph.steps)
	for len(manager.graph.steps) > 0 || manager.Working(){
		manager.time++
		manager.Clear()
		for index, key := range manager.graph.steps {
			if manager.graph.Check(key) {
				worker, ok := manager.GetWorker()
				if !ok { break }
				manager.Assign(worker, key)
				manager.graph.steps = append(manager.graph.steps[:index], manager.graph.steps[index + 1:]...)
			}
		}
	}

	return manager.time - 1
}


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


func main() {
	lines := lib.GetItems("inputs/7.txt")
	graph := Graph{make(map[int][]int), []int{}, []int{}}
	re, _ := regexp.Compile(`Step (.) .* step (.).*`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		graph.Add(int(matches[1][0]), int(matches[2][0]))
	}
	manager := EventManager{make([]*Event, WORKERS), graph, -1}
	time := manager.Process()
	fmt.Printf("%d\n", time)
}