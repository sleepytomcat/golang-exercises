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

func Isomorphic(tree1 *TreeNode, tree2 *TreeNode) bool {
	switch {
	case tree1 == nil && tree2 == nil:
		return true
	case tree1 == nil || tree2 == nil: // note the case above
		return false
	case Isomorphic(tree1.left, tree2.left) && Isomorphic(tree1.right, tree2.right):
		return true
	case Isomorphic(tree1.left, tree2.right) && Isomorphic(tree1.right, tree2.left):
		return true
	default:
		return false
	}
}

func main() {
	tree1 := &TreeNode{&TreeNode{nil, nil, 8}, &TreeNode{nil, nil, 4}, 10}
	tree2 := &TreeNode{nil, nil, 42}

	fmt.Println(tree1)
	fmt.Println(Isomorphic(tree1, tree1))
	fmt.Println(Isomorphic(nil, nil))
	fmt.Println(Isomorphic(tree1, tree2))
}
