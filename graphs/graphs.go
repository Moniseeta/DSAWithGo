package main

import (
	"fmt"
	"strconv"
)

type Graph struct {
	vertices []*Vertex
}
type Vertex struct {
	key      string
	adjacent []*Vertex
}

// AddVertex to graph
func (g *Graph) AddVertex(k string) {
	if contains(g.vertices, k) {
		fmt.Println("Vertex already exists")
		return
	}
	vertex := &Vertex{
		key: k,
	}
	g.vertices = append(g.vertices, vertex)
}

// AddEdge to graph
func (g *Graph) AddEdge(from, to string) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex
func (g Graph) getVertex(k string) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// Print prints adjacent vertices for each vertex in graph
func (g Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex: %v: ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
}

// contains checks if the vertex contained in the graph
func contains(vertices []*Vertex, k string) bool {
	for _, v := range vertices {
		if v.key == k {
			return true
		}
	}
	return false
}

func Init() *Graph {
	graph := &Graph{}
	return graph
}

func main() {
	graph := Init()
	for i := 0; i < 5; i++ {
		graph.AddVertex(strconv.Itoa(i))
	}

	graph.AddEdge("1", "2")
	graph.AddEdge("1", "2")
	graph.AddEdge("6", "2")
	graph.AddEdge("3", "2")
	graph.Print()
}
