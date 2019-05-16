package graphs

import (
	"container/list"
	"errors"
	"fmt"
)

// UnDirectedGraph defines a undirected graph
type UnDirectedGraph struct {
	vertexCount        int
	edgeCount          int
	adjacentVertices   [][]int
	visited            []bool
	pathTo             []int
	distanceTo         []int
	connectedComponent [][]int
	edges              [][]int
}

// NewUnDirectedGraph initalises a new undirected graph with vertexCount vertices.
func NewUnDirectedGraph(vertexCount int) *UnDirectedGraph {
	return &UnDirectedGraph{
		vertexCount, 0, make([][]int, vertexCount), nil, nil, nil, nil, make([][]int, 0),
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
		u.edges = append(u.edges, []int{vertex1, vertex2})
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
func (u *UnDirectedGraph) DFSRecursively(startingVertex int) (vertices []int, err error) {
	if !u.isVertexValid(startingVertex) {
		return nil, errors.New("vertex not found")
	}
	u.visited = make([]bool, u.vertexCount)
	u.pathTo = make([]int, u.vertexCount)
	return u.dfsRecursively(startingVertex, &u.visited), nil
}

// DFS does a depth first search
func (u *UnDirectedGraph) DFS(startingVertex int) (vertices []int, err error) {
	if !u.isVertexValid(startingVertex) {
		return nil, errors.New("vertex not found")
	}
	u.visited = make([]bool, u.vertexCount)
	u.pathTo = make([]int, u.vertexCount)
	stack := list.New()
	stack.PushBack(startingVertex)

	for stack.Len() > 0 {
		// pop stack
		v := stack.Remove(stack.Back())
		vertex := v.(int)

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
				stack.PushBack(adjs[i])
				u.pathTo[adjs[i]] = vertex
			}
		}
	}

	return
}

// BFS does a breadth first search starting from startingVertex in graph
func (u *UnDirectedGraph) BFS(startingVertex int) (vertices []int, err error) {
	if !u.isVertexValid(startingVertex) {
		return nil, errors.New("vertex not found")
	}
	u.visited = make([]bool, u.vertexCount)
	u.pathTo = make([]int, u.vertexCount)
	u.distanceTo = make([]int, u.vertexCount)
	queue := list.New()
	queue.PushBack(startingVertex)
	u.visited[startingVertex] = true

	for queue.Len() > 0 {
		// dequeue
		v := queue.Remove(queue.Front())
		vertex := v.(int)
		vertices = append(vertices, vertex)

		// get all its adjacent vertices.
		adjs, _ := u.GetAdjacentVertices(vertex)
		for i := 0; i < len(adjs); i++ {
			if !u.visited[adjs[i]] {
				queue.PushBack(adjs[i])
				u.visited[adjs[i]] = true
				u.pathTo[adjs[i]] = vertex
				u.distanceTo[adjs[i]] = u.distanceTo[vertex] + 1
			}
		}
	}
	return
}

// GetDFSPath gets the path from startingVertex to endingVertex using DFS
func (u *UnDirectedGraph) GetDFSPath(startingVertex int, endingVertex int) (path []int, err error) {
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
	for vertex != startingVertex {
		path = append([]int{vertex}, path...)
		vertex = u.pathTo[vertex]
	}
	path = append([]int{vertex}, path...)

	return
}

// GetBFSPath gets the BFS path from startingVertex to endingVertex.
// Using BFS, the path is also the mimimum path (mimimum number of edges).
func (u *UnDirectedGraph) GetBFSPath(startingVertex int, endingVertex int) (path []int, err error) {
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
	for u.distanceTo[vertex] != 0 {
		path = append([]int{vertex}, path...)
		vertex = u.pathTo[vertex]
	}
	path = append([]int{vertex}, path...)
	return
}

// GetConnectedComponents gets all the connected component of a graph
func (u *UnDirectedGraph) GetConnectedComponents() (connectedCompoent [][]int) {
	u.visited = make([]bool, u.vertexCount)
	u.pathTo = make([]int, u.vertexCount)
	u.connectedComponent = make([][]int, 0)

	for i := 0; i < u.vertexCount; i++ {
		if !u.visited[i] {
			vertices := u.dfsRecursively(i, &u.visited)
			u.connectedComponent = append(u.connectedComponent, vertices)
		}
	}
	return u.connectedComponent
}

// selfLoop checks for a given vertex, whether there is a loop exist (connects to itself.)
func (u *UnDirectedGraph) selfLoop(vertex int) bool {
	adjs, _ := u.GetAdjacentVertices(vertex)
	for _, adj := range adjs {
		if adj == vertex {
			return true
		}
	}
	return false
}

// parallel checks whether two vertices are connected to each other in parallel. (connected more than once.)
func (u *UnDirectedGraph) parallel(vertex1, vertex2 int) bool {
	adjs, _ := u.GetAdjacentVertices(vertex1)
	count := 0
	for _, adj := range adjs {
		if adj == vertex2 {
			count++
		}
		if count == 2 {
			return true
		}
	}
	return false
}

// GetCyclicPath gets a cyclic path in the graph, if not found, return nil.
func (u *UnDirectedGraph) GetCyclicPath() (path []int) {
	// Self loop, can return directly.
	for i := 0; i < u.vertexCount; i++ {
		if u.selfLoop(i) {
			return []int{i, i}
		}
	}

	// Parallel edges, can return directly.
	for i := 0; i < u.vertexCount-1; i++ {
		for j := i + 1; j < u.vertexCount; j++ {
			if u.parallel(i, j) {
				return []int{i, j, i}
			}
		}
	}

	u.visited = make([]bool, u.vertexCount)
	u.pathTo = make([]int, u.vertexCount)
	for i := 0; i < u.vertexCount; i++ {
		if !u.visited[i] && len(path) == 0 {
			u.dfsForCyclicPath(i, -1, &path)
		}
	}
	return
}

func (u *UnDirectedGraph) dfsForCyclicPath(vertex int, pathToVertex int, path *[]int) {
	u.visited[vertex] = true
	adjs, _ := u.GetAdjacentVertices(vertex)
	for _, adj := range adjs {
		// If we already found the cyclic path, don't do anything but quit the recursive loop
		if len(*path) != 0 {
			return
		}
		// if the adj is not visited, we add it to path and mark as visited.
		// and also dfs further down.
		if !u.visited[adj] {
			u.pathTo[adj] = vertex
			u.dfsForCyclicPath(adj, vertex, path)
		} else if pathToVertex != adj {
			// if the adj is already visited, and it's not equal to the pathToVertex, it means we have a cycle.
			// that's because this means there is an edge between this adj vertex and some previous visited vertex.
			// and that vertex is not the one we are coming from.
			for v := vertex; v != adj; v = u.pathTo[v] {
				(*path) = append([]int{v}, (*path)...)
			}
			(*path) = append([]int{adj}, (*path)...)
			(*path) = append((*path), adj)
		}
	}
}

// HasCycle using union-find method to find whether there is cycle.
func (u *UnDirectedGraph) HasCycle() bool {
	// let's init x sets, x = vertexCount. which means, each vertex is a set.
	parent := make([]int, u.vertexCount)
	// init them to -1
	for i := 0; i < u.vertexCount; i++ {
		parent[i] = -1
	}

	// we go through the edges one by one.
	for _, edge := range u.edges {
		setParent1 := findInSet(parent, edge[0])
		setParent2 := findInSet(parent, edge[1])

		// if these two edges are in the same union.
		if setParent1 == setParent2 {
			return true
		}
		parent = mergeSet(parent, setParent1, setParent2)
	}
	return false
}

func findInSet(parent []int, vertex int) int {
	// this vertex is the parent of the set it belongs to.
	if parent[vertex] == -1 {
		return vertex
	}
	// if not, recursively find its parent in the set.
	return findInSet(parent, parent[vertex])
}

func mergeSet(parent []int, vertex1, vertex2 int) []int {
	setParent1 := findInSet(parent, vertex1)
	setParent2 := findInSet(parent, vertex2)
	parent[setParent1] = setParent2
	return parent
}

// GetBipartiteParts gets the two parties if the graph is a bi-partite graph
func (u *UnDirectedGraph) GetBipartiteParts() (parts [][]int) {
	u.visited = make([]bool, u.vertexCount)
	color := make([]bool, u.vertexCount)
	for i := 0; i < u.vertexCount; i++ {
		if !u.visited[i] {
			stack := list.New()
			stack.PushBack(i)

			// run a dfs.
			for stack.Len() > 0 {
				// pop
				v := stack.Remove(stack.Back())
				vertex := v.(int)

				if !u.visited[vertex] {
					u.visited[vertex] = true
				}

				adjs, _ := u.GetAdjacentVertices(vertex)
				for _, adj := range adjs {
					// if this adjacent vertex is not visited yet, we set it to the different color.
					if !u.visited[adj] {
						color[adj] = !color[vertex]
						stack.PushBack(adj)
					} else if color[adj] == color[vertex] {
						// if this adjacent vertex is already visted and it has the same color as vertex.
						// it is not a bipartie graph.
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
