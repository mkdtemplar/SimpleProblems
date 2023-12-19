package main

import "fmt"

type Node struct {
	Neighbors []*Node
}
type Graph struct {
	nodes map[int]*Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*Node),
	}
}
func (g *Graph) AddNode(nodeID int) {
	if _, exists := g.nodes[nodeID]; !exists {
		newNode := &Node{
			Neighbors: []*Node{},
		}
		g.nodes[nodeID] = newNode
		fmt.Println("New node added to graph")
	} else {
		fmt.Println("Node already exists!")
	}
}
func (g *Graph) AddEdge(nodeID1, nodeID2 int) {
	node1 := g.nodes[nodeID1]
	node2 := g.nodes[nodeID2]

	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node1)
}
func (g *Graph) removeEdge(node, neighbor *Node) {
	index := -1
	for i, n := range node.Neighbors {
		if n == neighbor {
			index = i
			break
		}
	}
	if index != -1 {
		node.Neighbors = append(node.Neighbors[:index], node.Neighbors[index+1:]...)
	}
}

func (g *Graph) RemoveEdge(node, neighbor *Node) {
	g.removeEdge(node, neighbor)
	g.removeEdge(neighbor, node)
	fmt.Println("Edge successfully removed")
}
func (g *Graph) RemoveNode(nodeID int) {
	node, exists := g.nodes[nodeID]
	if !exists {
		fmt.Println("Node doesn't exist")
		return
	}

	for _, neighbor := range node.Neighbors {
		g.RemoveEdge(node, neighbor)
	}
	delete(g.nodes, nodeID)
	fmt.Println("Node deleted successfully")
}

func main() {
	g := NewGraph()
	g.AddNode(10)
	g.AddNode(20)
	g.AddNode(30)
	g.AddNode(40)
	g.AddEdge(10, 20)
	g.AddEdge(20, 30)
	g.AddEdge(20, 40)
	g.AddEdge(30, 40)

}
