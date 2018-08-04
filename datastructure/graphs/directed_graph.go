package graphs

import (
	"errors"
	"fmt"
)

// DirectedGraph defines a undirected graph
type DirectedGraph struct {
	vertexCount        int
	edgeCount          int
	adjacentVertices   [][]int
	visited            []bool
	pathTo             []int
	distanceTo         []int
	connectedComponent [][]int
	indegree           []int
}

// NewDirectedGraph initalises a new directed graph with vertexCount vertices.
func NewDirectedGraph(vertexCount int) *DirectedGraph {
	return &DirectedGraph{
		vertexCount, 0, make([][]int, vertexCount), nil, nil, nil, nil, make([]int, vertexCount),
	}
}

func (u *DirectedGraph) isVertexValid(vertex int) bool {
	return vertex >= 0 && vertex < u.vertexCount
}

// GetVertexCount gets vertex count
func (u *DirectedGraph) GetVertexCount() int {
	return u.vertexCount
}

// GetEdgeCount gets the edge count
func (u *DirectedGraph) GetEdgeCount() int {
	return u.edgeCount
}

// AddEdge adds an edge to the graph
func (u *DirectedGraph) AddEdge(vertex1, vertex2 int) error {
	if u.isVertexValid(vertex1) && u.isVertexValid(vertex2) {
		u.adjacentVertices[vertex1] = append(u.adjacentVertices[vertex1], vertex2)
		u.edgeCount++
		u.indegree[vertex2]++
		return nil
	}
	return errors.New("vertex not found")
}

// GetAdjacentVertices gets all adjacent vertices for a given vertex
func (u *DirectedGraph) GetAdjacentVertices(vertex int) ([]int, error) {
	if u.isVertexValid(vertex) {
		return u.adjacentVertices[vertex], nil
	}
	return nil, errors.New("vertex not found")
}

// GetVertexInDegree gets in degree for a given vertex
func (u *DirectedGraph) GetVertexInDegree(vertex int) (int, error) {
	if u.isVertexValid(vertex) {
		return u.indegree[vertex], nil
	}
	return 0, errors.New("vertex not found")
}

// GetVertexOutDegree gets the out degree of a given vertex
func (u *DirectedGraph) GetVertexOutDegree(vertex int) (int, error) {
	if u.isVertexValid(vertex) {
		return len(u.adjacentVertices[vertex]), nil
	}
	return 0, errors.New("vertex not found")
}

// Print prints the graph.
func (u *DirectedGraph) Print() string {
	res := ""
	res += fmt.Sprintf("Vertex Count: %d, Edge Count: %d\n", u.vertexCount, u.edgeCount)
	for vertex, adjacentVertices := range u.adjacentVertices {
		res += fmt.Sprintf("Vertex %d: %v\n", vertex, adjacentVertices)
	}
	return res
}

// Reverse reversees a directed graph, a.k.a revere all edges.
func (u *DirectedGraph) Reverse() (uv *DirectedGraph) {
	uv = NewDirectedGraph(u.vertexCount)
	for i := 0; i < u.vertexCount; i++ {
		for _, adj := range u.adjacentVertices[i] {
			uv.AddEdge(adj, i)
		}
	}
	return
}

func (u *DirectedGraph) dfsRecursively(startingVertex int, visited *[]bool) (vertices []int) {
	vertices = append(vertices, startingVertex)
	(*visited)[startingVertex] = true

	adjs, _ := u.GetAdjacentVertices(startingVertex)
	for _, v := range adjs {
		if !(*visited)[v] {
			vertices = append(vertices, u.dfsRecursively(v, visited)...)
			u.pathTo[v] = startingVertex
		}
	}
	return
}

// DFSRecursively does a dfs search using rescursive method
func (u *DirectedGraph) DFSRecursively(startingVertex int) (vertices []int, err error) {
	if !u.isVertexValid(startingVertex) {
		return nil, errors.New("vertex not found")
	}
	u.visited = make([]bool, u.vertexCount)
	u.pathTo = make([]int, u.vertexCount)
	return u.dfsRecursively(startingVertex, &u.visited), nil
}

// DFS does a depth first search
func (u *DirectedGraph) DFS(startingVertex int) (vertices []int, err error) {
	if !u.isVertexValid(startingVertex) {
		return nil, errors.New("vertex not found")
	}
	u.visited = make([]bool, u.vertexCount)
	u.pathTo = make([]int, u.vertexCount)
	stack := []int{startingVertex}

	for {
		if len(stack) == 0 {
			break
		}
		// pop stack
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// only if this vertex has not been visited, we mark it as visited and add into result.
		if !u.visited[vertex] {
			vertices = append(vertices, vertex)
			u.visited[vertex] = true
		}

		// get all its adjacent vertices.
		adjs, _ := u.GetAdjacentVertices(vertex)
		for i := len(adjs) - 1; i >= 0; i-- {
			// only add to stack if it's not visited yet.
			if !u.visited[adjs[i]] {
				stack = append(stack, adjs[i])
				u.pathTo[adjs[i]] = vertex
			}
		}
	}

	return
}

// BFS does a breadth first search starting from startingVertex in graph
func (u *DirectedGraph) BFS(startingVertex int) (vertices []int, err error) {
	if !u.isVertexValid(startingVertex) {
		return nil, errors.New("vertex not found")
	}
	u.visited = make([]bool, u.vertexCount)
	u.pathTo = make([]int, u.vertexCount)
	u.distanceTo = make([]int, u.vertexCount)
	queue := []int{startingVertex}
	u.visited[startingVertex] = true

	for {
		if len(queue) == 0 {
			break
		}
		// dequeue
		vertex := queue[0]
		queue = queue[1:]
		vertices = append(vertices, vertex)

		// get all its adjacent vertices.
		adjs, _ := u.GetAdjacentVertices(vertex)
		for i := 0; i < len(adjs); i++ {
			if !u.visited[adjs[i]] {
				queue = append(queue, adjs[i])
				u.visited[adjs[i]] = true
				u.pathTo[adjs[i]] = vertex
				u.distanceTo[adjs[i]] = u.distanceTo[vertex] + 1
			}
		}
	}
	return
}

// GetDFSPath gets the path from startingVertex to endingVertex using DFS
func (u *DirectedGraph) GetDFSPath(startingVertex int, endingVertex int) (path []int, err error) {
	if !u.isVertexValid(startingVertex) || !u.isVertexValid(endingVertex) {
		return nil, errors.New("vertex not found")
	}

	u.pathTo = make([]int, u.vertexCount)
	u.visited = make([]bool, u.vertexCount)
	u.DFS(startingVertex)

	if !u.visited[endingVertex] {
		return nil, errors.New("path not found")
	}

	vertex := endingVertex
	for {
		path = append([]int{vertex}, path...)
		vertex = u.pathTo[vertex]
		if vertex == startingVertex {
			break
		}
	}
	path = append([]int{vertex}, path...)

	return
}

// GetBFSPath gets the BFS path from startingVertex to endingVertex.
// Using BFS, the path is also the mimimum path (mimimum number of edges).
func (u *DirectedGraph) GetBFSPath(startingVertex int, endingVertex int) (path []int, err error) {
	if !u.isVertexValid(startingVertex) || !u.isVertexValid(endingVertex) {
		return nil, errors.New("vertex not found")
	}

	u.pathTo = make([]int, u.vertexCount)
	u.distanceTo = make([]int, u.vertexCount)
	u.visited = make([]bool, u.vertexCount)

	u.BFS(startingVertex)

	if !u.visited[endingVertex] {
		return nil, errors.New("path not found")
	}

	vertex := endingVertex
	for {
		if u.distanceTo[vertex] != 0 {
			path = append([]int{vertex}, path...)
			vertex = u.pathTo[vertex]
		} else {
			path = append([]int{vertex}, path...)
			break
		}
	}
	return
}

// GetCyclicPath gets a cyclic path in the graph, if not found, return nil.
func (u *DirectedGraph) GetCyclicPath() (path []int) {
	// we run a DFS, to find loop.
	u.visited = make([]bool, u.vertexCount)

	// loop through all vertices
	for i := 0; i < u.vertexCount; i++ {
		// only if it's not visited yet.
		if !u.visited[i] {
			stack := []int{i}
			u.pathTo = make([]int, u.vertexCount)
			u.pathTo[i] = -1

			// dfs
		DFS:
			for {
				if len(stack) == 0 {
					break
				}

				// pop
				vertex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if !u.visited[vertex] {
					u.visited[vertex] = true
				}

				// get all of its adjs.
				adjs, _ := u.GetAdjacentVertices(vertex)
				for _, adj := range adjs {
					if !u.visited[adj] {
						stack = append(stack, adj)
						u.pathTo[adj] = vertex
					} else {
						// We have encountered a vertex that's been visited.
						for v := vertex; v != adj; v = u.pathTo[v] {
							path = append([]int{v}, path...)
						}
						path = append([]int{adj}, path...)
						path = append(path, adj)
						break DFS
					}
				}
			}
		}
	}
	return
}

// GetBipartiteParts gets the two parties if the graph is a bi-partite graph
func (u *DirectedGraph) GetBipartiteParts() (parts [][]int) {
	u.visited = make([]bool, u.vertexCount)
	color := make([]bool, u.vertexCount)
	for i := 0; i < u.vertexCount; i++ {
		if !u.visited[i] {
			stack := []int{i}

			// run a dfs.
			for {
				if len(stack) == 0 {
					break
				}

				vertex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if !u.visited[vertex] {
					u.visited[vertex] = true
				}

				adjs, _ := u.GetAdjacentVertices(vertex)
				for _, adj := range adjs {
					if !u.visited[adj] {
						color[adj] = !color[vertex]
						stack = append(stack, adj)
					} else if color[adj] == color[vertex] {
						return nil
					}
				}
			}

		}
	}
	parts = make([][]int, 2)
	for i, c := range color {
		if c {
			parts[0] = append(parts[0], i)
		} else {
			parts[1] = append(parts[1], i)
		}
	}
	return
}
