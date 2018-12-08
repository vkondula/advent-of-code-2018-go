package main

import (
	"fmt"
	"./lib"
	"strconv"
	"strings"
)

var sum = lib.Aggregate{0,0,0}

type Node struct {
	metadata []int
	children []*Node
}

func GetNode(input []string) (*Node, []string) {
	childrenCount, _ := strconv.Atoi(input[0])
	metadataCount, _ := strconv.Atoi(input[1])
	children := make([]*Node, childrenCount)
	metadata := make([]int, metadataCount)
	input = input[2:]
	for i := range children {
		children[i], input = GetNode(input)
	}
	for i := range metadata {
		metadata[i], _ = strconv.Atoi(input[i])
		sum.Sum(metadata[i])
	}
	return &Node{metadata, children}, input[metadataCount:]
}


func main() {
	line := lib.GetItems("inputs/8.txt")[0]
	input := strings.Split(line, " ")
	GetNode(input)
	fmt.Printf("%d\n", sum.Result)
}