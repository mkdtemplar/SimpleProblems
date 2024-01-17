package main

import "fmt"

// Graph structure
type Graph struct {
	vertices []*Vertex
}

// Vertex Adjacent Vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex will add a vertex to a graph
func (g *Graph) AddVertex(vertex int) error {
	if contains(g.vertices, vertex) {
		err := fmt.Errorf("vertex %d already exists", vertex)
		return err
	} else {
		v := &Vertex{
			key: vertex,
		}
		g.vertices = append(g.vertices, v)
	}
	return nil
}

// AddEdge will add ad edge from a vertex to a vertex
func (g *Graph) AddEdge(to, from int) error {
	toVertex := g.getVertex(to)
	fromVertex := g.getVertex(from)
	if toVertex == nil || fromVertex == nil {
		return fmt.Errorf("not a valid edge from %d ---> %d", from, to)
	} else if contains(fromVertex.adjacent, toVertex.key) {
		return fmt.Errorf("edge from vertex %d ---> %d already exists", fromVertex.key, toVertex.key)
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
	return nil
}

// getVertex will return a vertex point if exists or return nil
func (g *Graph) getVertex(vertex int) *Vertex {
	for i, v := range g.vertices {
		if v.key == vertex {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(v []*Vertex, key int) bool {
	for _, v := range v {
		if v.key == key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("%d -> ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%d ", v.key)
		}
		fmt.Println()
	}
}

func errorHandling(err error) {
	if err != nil {
		return
	}
}

func PrintEgDirectedGraph() {
	g := &Graph{}
	err := g.AddVertex(1)
	errorHandling(err)
	err = g.AddVertex(2)
	errorHandling(err)

	g.AddVertex(3)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(1, 3)
	g.AddEdge(3, 1)
	g.Print()
}

func main() {
	PrintEgDirectedGraph()

}
