package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (n *TreeNode) String() string {
	return strconv.Itoa(n.value)
}

func (b *Tree) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()

}

func (b *Tree) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}

func (b *Tree) inOrderTraversalByNode(sb *strings.Builder, root *TreeNode) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	b.inOrderTraversalByNode(sb, root.right)
}

func (b *Tree) add(value int) *Tree {
	b.root = b.addNode(b.root, value)
	return b
}

func (b *Tree) addNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{value: value}
	}

	if value < root.value {
		root.left = b.addNode(root.left, value)
	} else {
		root.right = b.addNode(root.right, value)
	}

	return root
}

func (b *Tree) search(value int) (*TreeNode, bool) {
	return b.searchByNode(b.root, value)
}

func (b *Tree) searchByNode(root *TreeNode, value int) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	if value == root.value {
		return root, true
	} else if value < root.value {
		return b.searchByNode(root.left, value)
	} else {
		return b.searchByNode(root.right, value)
	}
}

func (b *Tree) remove(value int) {
	b.removeByNode(b.root, value)
}

func (b *Tree) removeByNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return root
	}
	if value > root.value {
		root.right = b.removeByNode(root.right, value)
	} else if value < root.value {
		root.left = b.removeByNode(root.left, value)
	} else {
		if root.left == nil {
			return root.left
		} else {
			temp := root.left
			for temp.right != nil {
				temp = temp.right
			}
			root.value = temp.value
			root.left = b.removeByNode(root.left, value)
		}
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

	tree := &Tree{}
	tree = tree.add(8)

	tree.add(4)
	tree.add(15)
	tree.add(1)
	tree.add(7)
	tree.add(5)

	fmt.Println(levelOrder(tree.root))
	fmt.Println(tree.search(15))

	//tree.remove(15)
	fmt.Println(tree)
}
