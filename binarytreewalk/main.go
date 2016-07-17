package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	left *TreeNode
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

func main() {
	tree := &TreeNode{&TreeNode{nil, nil, 8}, &TreeNode{nil, nil, 4}, 10}
	fmt.Println(tree)
}
