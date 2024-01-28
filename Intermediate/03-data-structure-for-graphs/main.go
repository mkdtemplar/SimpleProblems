package main

import "fmt"

type Vertex[T comparable] struct {
	key      T
	adjacent []*Vertex[T]
}

type Graph[T comparable] struct {
	vertices []*Vertex[T]
}

func (g *Graph[T]) AddVertex(k T) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added becouse allredy axists", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex[T]{key: k})
	}
}

func (g *Graph[T]) AddEdge(from T, to T) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v --> %v) ", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("edge allreday exists (%v --> %v) ", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph[T]) getVertex(vertex T) *Vertex[T] {
	for i, v := range g.vertices {
		if v.key == vertex {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph[T]) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
}

func contains[T comparable](s []*Vertex[T], k T) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func main() {
	test := Graph[int]{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(1, 2)
	test.AddEdge(1, 3)
	test.AddEdge(2, 4)

	test.Print()
}
