package graph

import (
	"fmt"
	"math"
)

// ShortestPath returns weight of the shortest possible path
// found between two vertices in a weighted directed graph.
// Returns infinity and an error if there's no such path found.
func (g *Graph) ShortestPath(from, to string) (float64, error) {
	var ok bool
	From, ok := g.Vertices[from]
	if !ok {
		return math.Inf(+1), fmt.Errorf("Vertex \"%v\" doesn't exist", from)
	}
	To, ok := g.Vertices[to]
	if !ok {
		return math.Inf(+1), fmt.Errorf("Vertex \"%v\" doesn't exist", to)
	}

	weight := g.Dijkstra(From, To)

	if math.IsInf(weight, +1) {
		return weight, fmt.Errorf("Path between \"%v\" and \"%v\" doesn't exist", from, to)
	}

	return weight, nil
}

type nodes map[*Vertex]float64

// Dijkstra implements Dijkstra's Shortest Path algorithm in
// a weighted directed graph.
func (g *Graph) Dijkstra(from, to *Vertex) float64 {
	if from == to {
		return 0
	}

	todo := make(nodes)
	visited := make(nodes)

	visited[from] = 0
	current := from
	for {
		for _, edge := range current.Neighborhood {
			if _, ok := visited[edge.To]; ok {
				continue
			}
			if weight, ok := todo[edge.To]; ok && weight < visited[current]+edge.Weight {
				continue
			}
			todo[edge.To] = visited[current] + edge.Weight
		}
		if len(todo) == 0 {
			break
		}

		min := math.Inf(+1)
		for node, weight := range todo {
			if weight <= min {
				current = node
				min = weight
			}
		}
		delete(todo, current)
		visited[current] = min
	}

	if weigth, ok := visited[to]; ok {
		return weigth
	}
	return math.Inf(+1)
}
