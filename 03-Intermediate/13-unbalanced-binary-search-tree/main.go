package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) add(value int) *Tree {
	t.root = t.addNode(t.root, value)
	return t
}

func (t *Tree) addNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{value: value}
	}

	if value < root.value {
		root.left = t.addNode(root.left, value)
	} else {
		root.right = t.addNode(root.right, value)
	}

	return root
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// Create a queue and insert root
	var queue []*TreeNode
	queue = append(queue, root)

	// Create result slice
	var result [][]int

	// Process as long as queue is not empty
	for len(queue) > 0 {
		// Get the current size or length of the queue.
		// This indicates the total number of nodes that are part of current level
		sz := len(queue)
		var level []int
		for i := 0; i < sz; i++ {
			// Remove a node
			node := queue[0]
			queue = queue[1:]

			// Visit the node. Here visiting means collecting it into the output array
			level = append(level, node.value)

			// Insert children of the node into the queue
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		// level is filled with one level of nodes' values. Insert this into the final
		// result
		result = append(result, level)
	}
	// result is ready to be returned
	return result
}

func main() {

	root := &Tree{}
	root = root.add(8)

	root.add(4)
	root.add(15)
	root.add(1)
	root.add(7)

	root.add(5)

	fmt.Println(levelOrder(root.root))
}
