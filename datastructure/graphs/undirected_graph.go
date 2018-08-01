package graphs

import (
	"errors"
	"fmt"
)

// UnDirectedGraph defines a undirected graph
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

// GetVertexDegree gets the degree of a given vertex
func (u *UnDirectedGraph) GetVertexDegree(vertex int) (int, error) {
	if u.isVertexValid(vertex) {
		return len(u.adjacentVertices[vertex]), nil
	}
	return 0, errors.New("vertex not found")
}

// Print prints the graph.
func (u *UnDirectedGraph) Print() string {
	res := ""
	res += fmt.Sprintf("Vertex Count: %d, Edge Count: %d\n", u.vertexCount, u.edgeCount)
	for vertex, adjacentVertices := range u.adjacentVertices {
		res += fmt.Sprintf("Vertex %d: %v\n", vertex, adjacentVertices)
	}
	return res
}

func (u *UnDirectedGraph) dfsRecursively(startingVertex int, visited *[]bool) (vertices []int) {
	if !(*visited)[startingVertex] {
		vertices = append(vertices, startingVertex)
		(*visited)[startingVertex] = true
	}

	adjs, _ := u.GetAdjacentVertices(startingVertex)
	for _, v := range adjs {
		if !(*visited)[v] {
			vertices = append(vertices, u.dfsRecursively(v, visited)...)
		}
	}
	return
}

// DFSRecursively does a dfs search using rescursive method
func (u *UnDirectedGraph) DFSRecursively(startingVertex int) (vertices []int, err error) {
	if !u.isVertexValid(startingVertex) {
		return nil, errors.New("vertex not found")
	}
	visited := make([]bool, u.vertexCount)
	return u.dfsRecursively(startingVertex, &visited), nil
}

// DFS does a depth first search
func (u *UnDirectedGraph) DFS(startingVertex int) (vertices []int, err error) {
	if !u.isVertexValid(startingVertex) {
		return nil, errors.New("vertex not found")
	}
	visited := make([]bool, u.vertexCount)
	stack := []int{startingVertex}

	for {
		if len(stack) == 0 {
			break
		}
		// pop stack
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[vertex] {
			vertices = append(vertices, vertex)
			visited[vertex] = true
		}

		// get all its adjacent vertices.
		adjs, _ := u.GetAdjacentVertices(vertex)
		for i := len(adjs) - 1; i >= 0; i-- {
			if !visited[adjs[i]] {
				stack = append(stack, adjs[i])
			}
		}
	}

	return
}

// GetPath gets the path from startingVertex to endingVertex
func (u *UnDirectedGraph) GetPath(startingVertex int, endingVertex int) (path []int, err error) {
	if !u.isVertexValid(startingVertex) || !u.isVertexValid(endingVertex) {
		return nil, errors.New("vertex not found")
	}
	vertices, _ := u.DFS(startingVertex)

	for i := range vertices {
		if vertices[i] == endingVertex {
			path = vertices[:i+1]
			return
		}
	}
	return nil, errors.New("path not found")
}
