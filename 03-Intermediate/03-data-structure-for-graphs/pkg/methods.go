package pkg

import "fmt"

type Vertex[T comparable] struct {
	Value    T
	Adjacent []*Vertex[T]
}

type Graph[T comparable] struct {
	vertices []*Vertex[T]
}

func (g *Graph[T]) AddVertex(k T) []*Vertex[T] {
	if Contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added becouse allredy axists", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex[T]{Value: k})
	}
	return g.vertices
}

func (g *Graph[T]) AddEdge(from T, to T) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v --> %v) ", from, to)
		fmt.Println(err.Error())
	} else if Contains(fromVertex.Adjacent, to) {
		err := fmt.Errorf("edge allreday exists (%v --> %v) ", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
	}
}

func (g *Graph[T]) AddEdgeMultiGraph(from T, to T) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
}

func (g *Graph[T]) GetVertex(vertex T) *Vertex[T] {
	for i, v := range g.vertices {
		if v.Value == vertex {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph[T]) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : egdes -> ", v.Value)
		for _, v := range v.Adjacent {
			fmt.Printf(" %v ", v.Value)
		}
	}
}

func Contains[T comparable](s []*Vertex[T], k T) bool {
	for _, v := range s {
		if k == v.Value {
			return true
		}
	}
	return false
}
