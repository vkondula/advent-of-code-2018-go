package main

import (
	"fmt"
	"./lib"
	"strconv"
	"regexp"
)

type Node struct {
	value int
	left *Node
	right *Node
}

func (node *Node) Insert(value int) *Node {
	right := node.right
	nextRight := node.right.right
	newNode := &Node{value, right, nextRight}
	right.right = newNode
	nextRight.left = newNode
	return newNode
}

func (node *Node) Remove(delta int) (*Node, int) {
	tmp := node
	for i := 0; i < delta; i++ {
		tmp = tmp.left
	}
	tmp.left.right = tmp.right
	tmp.right.left = tmp.left
	return tmp.right, tmp.value
}


func InitList (value int) (*Node) {
	node := &Node{value, nil, nil}
	node.left = node
	node.right = node
	return node
}


func main() {
	line := lib.GetItems("inputs/9.txt")[0]
	re, _ := regexp.Compile(`(\d+) .* worth (\d+).*`)
	matches := re.FindStringSubmatch(line)
	playersCount, _ := strconv.Atoi(matches[1])
	if playersCount == 0 { panic("Division by zero!") }
	last, _ := strconv.Atoi(matches[2])
	players := make([]int, playersCount)
	activePlayer := 0
	activeNode := InitList(0)
	value := 0
	for i := 1; i <= last * 100; i++ {
		if i % 23 != 0 {
			activeNode = activeNode.Insert(i)
		} else {
			activeNode, value = activeNode.Remove(7)
			players[activePlayer] += value + i
		}
		activePlayer = (activePlayer + 1) % playersCount
	}
	_, max := lib.Max(players)
	fmt.Printf("%d\n", max)
}