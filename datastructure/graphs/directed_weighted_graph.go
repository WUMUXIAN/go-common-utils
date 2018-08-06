package graphs

import (
	"errors"
	"fmt"
)

// DirectedWeightedGraph defines a directed wegithed graph
type DirectedWeightedGraph struct {
	DirectedGraph
	adjacentEdges [][]DirectedWeightedEdge
}

// NewDirectedWeightedGraph initalises a new directed weighted graph with vertexCount vertices.
func NewDirectedWeightedGraph(vertexCount int) *DirectedWeightedGraph {
	return &DirectedWeightedGraph{
		DirectedGraph{
			UnDirectedGraph{vertexCount: vertexCount, adjacentVertices: make([][]int, vertexCount)},
			make([]int, vertexCount), nil, nil, nil,
		},
		make([][]DirectedWeightedEdge, vertexCount),
	}
}

// GetAdjacentVertices gets all adjacent vertices for a given vertex
func (d *DirectedWeightedGraph) GetAdjacentVertices(vertex int) ([]int, error) {
	return d.DirectedGraph.GetAdjacentVertices(vertex)
}

// AddEdge adds an edge to the directed weighted graph
func (d *DirectedWeightedGraph) AddEdge(from, to int, weight float64) error {
	if !d.isVertexValid(from) || !d.isVertexValid(to) {
		return errors.New("vertex not found")
	}
	d.DirectedGraph.AddEdge(from, to)
	d.adjacentEdges[from] = append(d.adjacentEdges[from], DirectedWeightedEdge{from, to, weight})
	return nil
}

// GetEdges prints all edges.
func (d *DirectedWeightedGraph) GetEdges() []DirectedWeightedEdge {
	res := make([]DirectedWeightedEdge, 0)
	for i := 0; i < d.vertexCount; i++ {
		edges := d.adjacentEdges[i]
		for _, edge := range edges {
			res = append(res, edge)
		}
	}
	return res
}

// Reverse reversees a directed graph, a.k.a revere all edges.
func (d *DirectedWeightedGraph) Reverse() (uv *DirectedWeightedGraph) {
	uv = NewDirectedWeightedGraph(d.vertexCount)
	for i := 0; i < d.vertexCount; i++ {
		for _, adj := range d.adjacentEdges[i] {
			uv.AddEdge(adj.GetTo(), adj.GetFrom(), adj.GetWeight())
		}
	}
	return
}

// Print prints the graph.
func (u *DirectedWeightedGraph) Print() string {
	res := ""
	res += fmt.Sprintf("Vertex Count: %d, Edge Count: %d\n", u.vertexCount, u.edgeCount)
	for vertex, adjacentEdges := range u.adjacentEdges {
		res += fmt.Sprintf("Vertex %d: %v\n", vertex, adjacentEdges)
	}
	return res
}
