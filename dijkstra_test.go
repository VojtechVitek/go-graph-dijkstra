package graph

import (
	"math"
	"testing"
)

type tcVertex string

type tcEdge struct {
	from   string
	to     string
	weight float64
}

type tcPath struct {
	from   string
	to     string
	err    bool
	weight float64
}

func TestShortestPath(t *testing.T) {
	testCases := []struct {
		vertices []tcVertex
		edges    []tcEdge
		paths    []tcPath
	}{
		{ // tc[0]
			[]tcVertex{"a"},
			[]tcEdge{},
			[]tcPath{
				{from: "a", to: "a", weight: 0},
			},
		},
		{ // tc[1]
			[]tcVertex{"a", "b"},
			[]tcEdge{
				{from: "a", to: "b", weight: 10},
			},
			[]tcPath{
				{from: "a", to: "b", weight: 10},
				{from: "b", to: "a", weight: math.Inf(+1), err: true},
			},
		},
		{ // tc[2]
			[]tcVertex{"a", "b"},
			[]tcEdge{
				{from: "a", to: "b", weight: 100},
				{from: "a", to: "b", weight: 10},
				{from: "a", to: "b", weight: 1},
				{from: "b", to: "a", weight: 2},
				{from: "b", to: "a", weight: 20},
				{from: "b", to: "a", weight: 200},
			},
			[]tcPath{
				{from: "a", to: "b", weight: 1},
				{from: "b", to: "a", weight: 2},
			},
		},
		{ // tc[3]
			[]tcVertex{"a", "b", "x", "y"},
			[]tcEdge{
				{from: "a", to: "b", weight: 10},
				{from: "x", to: "y", weight: math.Inf(+1)},
			},
			[]tcPath{
				{from: "a", to: "b", weight: 10},
				{from: "x", to: "y", weight: math.Inf(+1), err: true},
				{from: "a", to: "x", weight: math.Inf(+1), err: true},
			},
		},
		{ // tc[4]
			[]tcVertex{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
			[]tcEdge{
				{from: "a", to: "b", weight: 10},
				{from: "b", to: "c", weight: 5},
				{from: "b", to: "d", weight: 20},
				{from: "c", to: "d", weight: 10},
				{from: "d", to: "e", weight: 10},
				{from: "e", to: "f", weight: 10},
				{from: "f", to: "g", weight: 10},
				{from: "g", to: "i", weight: 10},
				{from: "a", to: "h", weight: 50},
				{from: "h", to: "i", weight: 50},
				{from: "h", to: "d", weight: 15},
				{from: "d", to: "h", weight: 15},
				{from: "d", to: "b", weight: 15},
				{from: "i", to: "a", weight: 70},
			},
			[]tcPath{
				{from: "a", to: "b", weight: 10},
				{from: "a", to: "d", weight: 25},
				{from: "d", to: "g", weight: 30},
				{from: "a", to: "h", weight: 40},
				{from: "d", to: "i", weight: 40},
				{from: "a", to: "i", weight: 65},
				{from: "h", to: "b", weight: 30},
				{from: "h", to: "c", weight: 35},
				{from: "i", to: "g", weight: 125},
				{from: "h", to: "a", weight: 120},
			},
		},
	}

	for i, tc := range testCases {
		graph := New()
		for _, vertex := range tc.vertices {
			graph.AddVertex(string(vertex))
		}
		for _, edge := range tc.edges {
			graph.AddEdge(edge.from, edge.to, edge.weight)
		}
		for _, path := range tc.paths {
			weight, err := graph.ShortestPath(path.from, path.to)
			if path.err && err == nil {
				t.Errorf("tc[%v]: From \"%v\" to \"%v\": Expected error", i, path.from, path.to)
			}
			if !path.err && err != nil {
				t.Errorf("tc[%v]: From \"%v\" to \"%v\": Unexpected error: %v", i, path.from, path.to, err)
			}
			if weight != path.weight {
				t.Errorf("tc[%v]: From \"%v\" to \"%v\": Expected %v, got %v", i, path.from, path.to, path.weight, weight)
			}
		}
	}
}
