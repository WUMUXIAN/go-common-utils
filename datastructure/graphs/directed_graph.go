package graphs

import (
	"errors"
)

// DirectedGraph defines a undirected graph
type DirectedGraph struct {
	UnDirectedGraph
	indegree    []int
	preorder    []int
	postorder   []int
	reversePost []int
}

// TransitiveClousure presents a transitive clousure.
type TransitiveClousure struct {
	graph [][]bool
}

// NewTransitiveClousure creates a transitive clousure from a given directed graph.
func NewTransitiveClousure(dGraph *DirectedGraph) *TransitiveClousure {
	tc := &TransitiveClousure{make([][]bool, dGraph.GetVertexCount())}
	for i := 0; i < dGraph.GetVertexCount(); i++ {
		tc.graph[i] = make([]bool, dGraph.GetVertexCount())
		for j := 0; j < dGraph.GetVertexCount(); j++ {
			if i == j {
				tc.graph[i][j] = true
			} else {
				_, err := dGraph.GetBFSPath(j, i)
				tc.graph[i][j] = (err == nil)
			}
		}
	}
	return tc
}

// NewDirectedGraph initalises a new directed graph with vertexCount vertices.
func NewDirectedGraph(vertexCount int) *DirectedGraph {
	return &DirectedGraph{
		UnDirectedGraph{vertexCount: vertexCount, adjacentVertices: make([][]int, vertexCount)},
		make([]int, vertexCount), nil, nil, nil,
	}
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

// GetCyclicPath gets a cyclic path in the graph, if not found, return nil.
func (u *DirectedGraph) GetCyclicPath() (path []int) {
	// loop through all vertices
	u.visited = make([]bool, u.vertexCount)
	u.pathTo = make([]int, u.vertexCount)
	marked := make([]bool, u.vertexCount)
	for i := 0; i < u.vertexCount; i++ {
		if !u.visited[i] && len(path) == 0 {
			u.dfsForCyclicPath(i, &marked, &path)
		}
	}
	return
}

// GetTopologyOrder get topology order.
func (u *DirectedGraph) GetTopologyOrder() (topology []int) {
	// if we have cyclic path, the topology order doesn't make sense.
	if len(u.GetCyclicPath()) != 0 {
		return
	}
	_, _, topology = u.GetOrder()
	return
}

// GetOrder get vertex orders (pre, post and reverse-post[topology])
func (u *DirectedGraph) GetOrder() (pre, post, reversePost []int) {
	u.preorder = make([]int, 0)
	u.postorder = make([]int, 0)
	u.reversePost = make([]int, 0)
	u.visited = make([]bool, u.vertexCount)

	for i := 0; i < u.vertexCount; i++ {
		if !u.visited[i] {
			u.dfsForOrder(i)
		}
	}

	return u.preorder, u.postorder, u.reversePost
}

func (u *DirectedGraph) dfsForOrder(vertex int) {
	u.visited[vertex] = true

	u.preorder = append(u.preorder, vertex)
	adjs, _ := u.GetAdjacentVertices(vertex)
	for _, adj := range adjs {
		if !u.visited[adj] {
			u.dfsForOrder(adj)
		}
	}

	u.reversePost = append([]int{vertex}, u.reversePost...)
	u.postorder = append(u.postorder, vertex)
}

func (u *DirectedGraph) dfsForCyclicPath(vertex int, inStack *[]bool, path *[]int) {
	u.visited[vertex] = true

	// put this vertex in stack.
	(*inStack)[vertex] = true
	adjs, _ := u.GetAdjacentVertices(vertex)
	for _, adj := range adjs {
		// We have the path found, so we should return only.
		if len(*path) != 0 {
			return
		}
		if !u.visited[adj] {
			u.pathTo[adj] = vertex
			u.dfsForCyclicPath(adj, inStack, path)
		} else if (*inStack)[adj] {
			// We've got a loop.
			for v := vertex; v != adj; v = u.pathTo[v] {
				// fmt.Println("endless loop happending", v, u.pathTo[v])
				(*path) = append([]int{v}, (*path)...)
			}
			(*path) = append([]int{adj}, (*path)...)
			(*path) = append((*path), adj)
		}
	}

	// When we reach here, the current path has reach an end, pop it from the stack
	(*inStack)[vertex] = false
}

// GetStronglyConnectedComponent gets all strongly connected component, each component contains a set of vertices
// It uses Kosaraju’s algorithm. (Reverse Graph -> Reverse Post Order -> DFS)
func (u *DirectedGraph) GetStronglyConnectedComponent() (stronglyConnectedComponent [][]int) {
	stronglyConnectedComponent = make([][]int, 0)

	// Get the reversed graph.
	graph := u.Reverse()

	// Get the post-reverse order of the reversed graph.
	_, _, postReversed := graph.GetOrder()

	u.visited = make([]bool, u.vertexCount)
	u.pathTo = make([]int, u.vertexCount)
	// Do a DFS using the postReversed order
	for _, v := range postReversed {
		if !u.visited[v] {
			vertices := u.dfsRecursively(v, &u.visited)
			stronglyConnectedComponent = append(stronglyConnectedComponent, vertices)
		}
	}

	return
}
