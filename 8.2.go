package main

import (
	"fmt"
	"./lib"
	"strconv"
	"strings"
)

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
	}
	return &Node{metadata, children}, input[metadataCount:]
}

func Value(node Node) int {
	sum := 0
	if len(node.children) == 0 {
		for _, value := range node.metadata { sum += value }
		return sum
	}
	for _, i := range node.metadata {
		index := i - 1
		if index >= 0 && index < len(node.children) { sum += Value(*node.children[index]) }
	}
	return sum
}


func main() {
	line := lib.GetItems("inputs/8.txt")[0]
	input := strings.Split(line, " ")
	root, _ := GetNode(input)
	fmt.Printf("%d\n", Value(*root))
}