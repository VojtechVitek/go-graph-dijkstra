package graph

import (
	"fmt"
	"math"
)

// Graph represents set of vertices and and edges.
type Graph struct {
	// Map instead of a slice for quick "Name"-based search
	Vertices map[string]*Vertex
	Edges    []*Edge
}

// Vertex represents a node in a directed weighted graph.
type Vertex struct {
	Name         string
	Neighborhood []*Edge
}

// Edge represents directed weighted path between the
// two vertices in a graph.
type Edge struct {
	From, To *Vertex
	Weight   float64
}

// New creates and returns an instance of a graph.
func New() *Graph {
	var graph Graph
	graph.Vertices = make(map[string]*Vertex, 0)
	graph.Edges = make([]*Edge, 0)

	return &graph
}

// AddVertex adds a node to the graph.
func (g *Graph) AddVertex(name string) error {
	if _, ok := g.Vertices[name]; ok {
		return fmt.Errorf("Vertex \"%v\" already exist", name)
	}
	g.Vertices[name] = &Vertex{
		Name:         name,
		Neighborhood: make([]*Edge, 0),
	}
	return nil
}

// AddEdge adds a directed weighted path between the
// two vertices in the graph.
func (g *Graph) AddEdge(from, to string, weight float64) error {
	if weight < 0.0 || weight == math.Inf(+1) {
		return fmt.Errorf("Weight must not be negative number or infinity")
	}

	var ok bool
	From, ok := g.Vertices[from]
	if !ok {
		return fmt.Errorf("Vertex \"%v\" doesn't exist", from)
	}
	To, ok := g.Vertices[to]
	if !ok {
		return fmt.Errorf("Vertex \"%v\" doesn't exist", from)
	}

	edge := Edge{
		From:   From,
		To:     To,
		Weight: weight,
	}

	g.Edges = append(g.Edges, &edge)
	From.Neighborhood = append(From.Neighborhood, &edge)

	return nil
}
