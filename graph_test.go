package graph

import (
	"math"
	"testing"
)

func TestCreateGraph(t *testing.T) {
	graph := New()
	if graph == nil {
		t.Fatalf("Unexpected error: nil graph")
	}
}

func TestAddVertices(t *testing.T) {
	graph := New()

	testCases := []struct {
		name string
		err  bool
	}{
		{name: "foo"},
		{name: "foo", err: true},
		{name: "bar"},
	}

	vertices := 0
	for i, tc := range testCases {
		if err := graph.AddVertex(tc.name); err == nil && tc.err {
			t.Errorf("tc[%v]: Expected error", i)
		} else if err != nil && !tc.err {
			t.Errorf("tc[%v]: Unexpected error: %v", i, err)
		}

		if !tc.err {
			vertices++
		}
	}

	if len(graph.Vertices) != vertices {
		t.Errorf("Expected %v vertices, got %v", vertices, len(graph.Vertices))
	}
}

func TestAddEdge(t *testing.T) {
	graph := New()
	graph.AddVertex("a")
	graph.AddVertex("b")

	testCases := []struct {
		from   string
		to     string
		weight float64
		err    bool
	}{
		// valid weigth
		{from: "a", to: "b", weight: 10},
		{from: "a", to: "b", weight: 0},
		// invalid weight
		{from: "a", to: "b", weight: -10, err: true},
		{from: "a", to: "b", weight: math.Inf(-1), err: true},
		{from: "a", to: "b", weight: math.Inf(+1), err: true},
		// invalid/undefined vertices
		{from: "a", to: "-", err: true},
		{from: "-", to: "b", err: true},
	}

	edges := 0
	for i, tc := range testCases {
		if err := graph.AddEdge(tc.from, tc.to, tc.weight); err == nil && tc.err {
			t.Errorf("tc[%v]: Expected error", i)
		} else if err != nil && !tc.err {
			t.Errorf("tc[%v]: Unexpected error: %v", i, err)
		}

		if !tc.err {
			edges++
		}
	}

	if len(graph.Edges) != edges {
		t.Errorf("Expected %v edges, got %v", edges, len(graph.Edges))
	}
}
