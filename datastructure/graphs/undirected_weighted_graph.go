package graphs

import (
	"errors"
	"fmt"

	"github.com/WUMUXIAN/go-common-utils/datastructure/trees"
)

// UnDirectedWeightGraph defines a undirected graph
type UnDirectedWeightGraph struct {
	vertexCount        int
	edgeCount          int
	adjacentVertices   [][]WeightedEdge
	visited            []bool
	pathTo             []int
	distanceTo         []int
	connectedComponent [][]int
	mimimumWeight      float64
}

// NewUnDirectedWeightedGraph initalises a new undirected weighted graph with vertexCount vertices.
func NewUnDirectedWeightedGraph(vertexCount int) *UnDirectedWeightGraph {
	return &UnDirectedWeightGraph{
		vertexCount, 0, make([][]WeightedEdge, vertexCount), nil, nil, nil, nil, 0,
	}
}

func (u *UnDirectedWeightGraph) isVertexValid(vertex int) bool {
	return vertex >= 0 && vertex < u.vertexCount
}

// GetVertexCount gets vertex count
func (u *UnDirectedWeightGraph) GetVertexCount() int {
	return u.vertexCount
}

// GetEdgeCount gets the edge count
func (u *UnDirectedWeightGraph) GetEdgeCount() int {
	return u.edgeCount
}

// AddEdge adds an edge to the graph
func (u *UnDirectedWeightGraph) AddEdge(vertex1, vertex2 int, weight float64) error {
	if u.isVertexValid(vertex1) && u.isVertexValid(vertex2) {
		u.adjacentVertices[vertex1] = append(u.adjacentVertices[vertex1], WeightedEdge{vertex1, vertex2, weight})
		u.adjacentVertices[vertex2] = append(u.adjacentVertices[vertex2], WeightedEdge{vertex2, vertex1, weight})
		u.edgeCount++
		return nil
	}
	return errors.New("vertex not found")
}

// GetAdjacentVertices gets all adjacent vertices for a given vertex
func (u *UnDirectedWeightGraph) GetAdjacentVertices(vertex int) ([]int, error) {
	if u.isVertexValid(vertex) {
		edges := u.adjacentVertices[vertex]
		res := make([]int, len(edges))
		for i := 0; i < len(res); i++ {
			res[i], _ = edges[i].GetOther(vertex)
		}
		return res, nil
	}
	return nil, errors.New("vertex not found")
}

// GetVertexDegree gets the degree of a given vertex
func (u *UnDirectedWeightGraph) GetVertexDegree(vertex int) (int, error) {
	if u.isVertexValid(vertex) {
		return len(u.adjacentVertices[vertex]), nil
	}
	return 0, errors.New("vertex not found")
}

// GetEdges prints all edges.
func (u *UnDirectedWeightGraph) GetEdges() []WeightedEdge {
	res := make([]WeightedEdge, 0)
	for i := 0; i < u.vertexCount; i++ {
		edges := u.adjacentVertices[i]
		selfLoop := false
		for _, edge := range edges {
			v, _ := edge.GetOther(i)
			// Add only once
			if i < v {
				res = append(res, edge)
			}
			// selfloop, we only add once.
			if i == v {
				if !selfLoop {
					res = append(res, edge)
					selfLoop = !selfLoop
				}
			}
		}
	}
	return res
}

// Print prints the graph.
func (u *UnDirectedWeightGraph) Print() string {
	res := ""
	res += fmt.Sprintf("Vertex Count: %d, Edge Count: %d\n", u.vertexCount, u.edgeCount)
	for vertex, adjacentVertices := range u.adjacentVertices {
		res += fmt.Sprintf("Vertex %d: %v\n", vertex, adjacentVertices)
	}
	return res
}

func (u *UnDirectedWeightGraph) dfsRecursively(startingVertex int, visited *[]bool) (vertices []int) {
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
func (u *UnDirectedWeightGraph) DFSRecursively(startingVertex int) (vertices []int, err error) {
	if !u.isVertexValid(startingVertex) {
		return nil, errors.New("vertex not found")
	}
	u.visited = make([]bool, u.vertexCount)
	u.pathTo = make([]int, u.vertexCount)
	return u.dfsRecursively(startingVertex, &u.visited), nil
}

// DFS does a depth first search
func (u *UnDirectedWeightGraph) DFS(startingVertex int) (vertices []int, err error) {
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
func (u *UnDirectedWeightGraph) BFS(startingVertex int) (vertices []int, err error) {
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
func (u *UnDirectedWeightGraph) GetDFSPath(startingVertex int, endingVertex int) (path []int, err error) {
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
func (u *UnDirectedWeightGraph) GetBFSPath(startingVertex int, endingVertex int) (path []int, err error) {
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

// GetConnectedComponents gets all the connected component of a graph
func (u *UnDirectedWeightGraph) GetConnectedComponents() (connectedCompoent [][]int) {
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

func (u *UnDirectedWeightGraph) selfLoop(vertex int) bool {
	adjs, _ := u.GetAdjacentVertices(vertex)
	for _, adj := range adjs {
		if adj == vertex {
			return true
		}
	}
	return false
}

func (u *UnDirectedWeightGraph) parallel(vertex1, vertex2 int) bool {
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
func (u *UnDirectedWeightGraph) GetCyclicPath() (path []int) {
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
			u.dfs(i, -1, &path)
		}
	}
	return
}

func (u *UnDirectedWeightGraph) dfs(vertex int, pathToVertex int, path *[]int) {
	u.visited[vertex] = true
	adjs, _ := u.GetAdjacentVertices(vertex)
	for _, adj := range adjs {
		// If we already found the cyclic path, don't do anything but quit the recursive loop
		if len(*path) != 0 {
			return
		}
		if !u.visited[adj] {
			u.pathTo[adj] = vertex
			u.dfs(adj, vertex, path)
		} else if pathToVertex != adj {
			// found it
			for v := vertex; v != adj; v = u.pathTo[v] {
				(*path) = append([]int{v}, (*path)...)
			}
			(*path) = append([]int{adj}, (*path)...)
			(*path) = append((*path), adj)
		}
	}
}

// GetBipartiteParts gets the two parties if the graph is a bi-partite graph
func (u *UnDirectedWeightGraph) GetBipartiteParts() (parts [][]int) {
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

// LazyPrimMinimumSpanningTree gets the mimimum spanning tree of the give directed weighted graph.
func (u *UnDirectedWeightGraph) LazyPrimMinimumSpanningTree() ([]WeightedEdge, float64) {
	connectedC := u.GetConnectedComponents()
	if len(connectedC) != 1 {
		return nil, 0
	}

	priorityQueue := &trees.Heap{
		HeapType:   trees.HeapTypeMin,
		Comparator: CompareEdges,
	}
	u.mimimumWeight = 0
	u.visited = make([]bool, u.vertexCount)
	res := make([]WeightedEdge, 0)

	for i := 0; i < u.vertexCount; i++ {
		if !u.visited[i] {
			u.prim(i, priorityQueue, &res)
		}
	}
	return res, u.mimimumWeight
}

func (u *UnDirectedWeightGraph) scanAdjsAndEnqueue(vertex int, priorityQueue *trees.Heap) {
	u.visited[vertex] = true
	edges := u.adjacentVertices[vertex]
	for _, edge := range edges {
		// if the other vertex is already visited, this edge is already handled, skip.
		w, _ := edge.GetOther(vertex)
		if !u.visited[w] {
			priorityQueue.Insert(edge)
		}
	}
	return
}

func (u *UnDirectedWeightGraph) prim(vertex int, priorityQueue *trees.Heap, res *[]WeightedEdge) {
	u.scanAdjsAndEnqueue(vertex, priorityQueue)
	for {
		if priorityQueue.Peek() == nil {
			break
		}

		// get the top edge.
		edge := priorityQueue.Pop().(WeightedEdge)
		v := edge.GetVertex1()
		w, _ := edge.GetOther(v)
		// if two nodes are in the tree already, skip.
		if u.visited[v] && u.visited[w] {
			continue
		}

		// add this edge to the result
		(*res) = append((*res), edge)
		u.mimimumWeight += edge.GetWeight()

		// for the other vertex, if it's not visited yet, add its adjs into the queue.
		if !u.visited[w] {
			u.scanAdjsAndEnqueue(w, priorityQueue)
		}
	}
}
