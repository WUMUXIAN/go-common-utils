package graphs

import (
	"errors"
	"fmt"
)

type UnDirectedGraph struct {
	vertexCount      int
	edgeCount        int
	adjacentVertices [][]int
}

// NewUnDirectedGraph initalises a new undirected graph with vertexCount vertices.
func NewUnDirectedGraph(vertexCount int) *UnDirectedGraph {
	return &UnDirectedGraph{
		vertexCount, 0, make([][]int, vertexCount),
	}
}

func (u *UnDirectedGraph) isVertexValid(vertex int) bool {
	return vertex >= 0 && vertex < u.vertexCount
}

// GetVertexCount gets vertex count
func (u *UnDirectedGraph) GetVertexCount() int {
	return u.vertexCount
}

// GetEdgeCount gets the edge count
func (u *UnDirectedGraph) GetEdgeCount() int {
	return u.edgeCount
}

// AddEdge adds an edge to the graph
func (u *UnDirectedGraph) AddEdge(vertex1, vertex2 int) error {
	if u.isVertexValid(vertex1) && u.isVertexValid(vertex2) {
		u.adjacentVertices[vertex1] = append(u.adjacentVertices[vertex1], vertex2)
		u.adjacentVertices[vertex2] = append(u.adjacentVertices[vertex2], vertex1)
		u.edgeCount++
		return nil
	}
	return errors.New("vertex not found")
}

// GetAdjacentVertices gets all adjacent vertices for a given vertex
func (u *UnDirectedGraph) GetAdjacentVertices(vertex int) ([]int, error) {
	if u.isVertexValid(vertex) {
		return u.adjacentVertices[vertex], nil
	}
	return nil, errors.New("vertex not found")
}

// Get the degree of a given vertex
func (u *UnDirectedGraph) GetVertexDegree(vertex int) (int, error) {
	if u.isVertexValid(vertex) {
		return len(u.adjacentVertices[vertex]), nil
	}
	return 0, errors.New("vertex not found")
}

func (u *UnDirectedGraph) Print() string {
	res := ""
	res += fmt.Sprintf("Vertex Count: %d, Edge Count: %d\n", u.vertexCount, u.edgeCount)
	for vertex, adjacentVertices := range u.adjacentVertices {
		res += fmt.Sprintf("Vertex %d: %v\n", vertex, adjacentVertices)
	}
	return res
}
