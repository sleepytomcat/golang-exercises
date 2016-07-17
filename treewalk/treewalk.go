package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

func (tree *TreeNode) String() string {
	if tree == nil {
		return "."
	} else {
		leftSubTree := tree.left.String()
		rightSubTree := tree.right.String()
		return "[" + leftSubTree + "]" + strconv.Itoa(tree.value) + "[" + rightSubTree + "]"
	}
}

func (tree *TreeNode) Walk(output *chan int) {
	if tree == nil {
		return
	} else {
		tree.WalkRecursive(output)
	}
	close(*output)
}

func (tree *TreeNode) WalkRecursive(output *chan int) {
	if tree == nil {
		return
	} else {
		*output <- tree.value
		tree.left.WalkRecursive(output)
		tree.right.WalkRecursive(output)
	}
}

func main() {
	tree := &TreeNode{
		&TreeNode{
			nil,
			nil,
			8},
		&TreeNode{
			nil,
			nil,
			4},
		10}

	output := make(chan int)
	go tree.Walk(&output)
	for value := range(output) {
		fmt.Println(value)
	}
}
