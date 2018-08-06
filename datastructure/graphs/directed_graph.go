package graphs

import (
	"errors"
)

// DirectedGraph defines a directed graph
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
func NewTransitiveClousure(dGraph Graph) *TransitiveClousure {
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
func (d *DirectedGraph) AddEdge(vertex1, vertex2 int) error {
	if d.isVertexValid(vertex1) && d.isVertexValid(vertex2) {
		d.adjacentVertices[vertex1] = append(d.adjacentVertices[vertex1], vertex2)
		d.edgeCount++
		d.indegree[vertex2]++
		return nil
	}
	return errors.New("vertex not found")
}

// GetVertexInDegree gets in degree for a given vertex
func (d *DirectedGraph) GetVertexInDegree(vertex int) (int, error) {
	if d.isVertexValid(vertex) {
		return d.indegree[vertex], nil
	}
	return 0, errors.New("vertex not found")
}

// GetVertexOutDegree gets the out degree of a given vertex
func (d *DirectedGraph) GetVertexOutDegree(vertex int) (int, error) {
	if d.isVertexValid(vertex) {
		return len(d.adjacentVertices[vertex]), nil
	}
	return 0, errors.New("vertex not found")
}

// Reverse reversees a directed graph, a.k.a revere all edges.
func (d *DirectedGraph) Reverse() (uv *DirectedGraph) {
	uv = NewDirectedGraph(d.vertexCount)
	for i := 0; i < d.vertexCount; i++ {
		for _, adj := range d.adjacentVertices[i] {
			uv.AddEdge(adj, i)
		}
	}
	return
}

// GetCyclicPath gets a cyclic path in the graph, if not found, return nil.
func (d *DirectedGraph) GetCyclicPath() (path []int) {
	// loop through all vertices
	d.visited = make([]bool, d.vertexCount)
	d.pathTo = make([]int, d.vertexCount)
	marked := make([]bool, d.vertexCount)
	for i := 0; i < d.vertexCount; i++ {
		if !d.visited[i] && len(path) == 0 {
			d.dfsForCyclicPath(i, &marked, &path)
		}
	}
	return
}

// GetTopologyOrder get topology order.
func (d *DirectedGraph) GetTopologyOrder() (topology []int) {
	// if we have cyclic path, the topology order doesn't make sense.
	if len(d.GetCyclicPath()) != 0 {
		return
	}
	_, _, topology = d.GetOrder()
	return
}

// GetOrder get vertex orders (pre, post and reverse-post[topology])
func (d *DirectedGraph) GetOrder() (pre, post, reversePost []int) {
	d.preorder = make([]int, 0)
	d.postorder = make([]int, 0)
	d.reversePost = make([]int, 0)
	d.visited = make([]bool, d.vertexCount)

	for i := 0; i < d.vertexCount; i++ {
		if !d.visited[i] {
			d.dfsForOrder(i)
		}
	}

	return d.preorder, d.postorder, d.reversePost
}

func (d *DirectedGraph) dfsForOrder(vertex int) {
	d.visited[vertex] = true

	d.preorder = append(d.preorder, vertex)
	adjs, _ := d.GetAdjacentVertices(vertex)
	for _, adj := range adjs {
		if !d.visited[adj] {
			d.dfsForOrder(adj)
		}
	}

	d.reversePost = append([]int{vertex}, d.reversePost...)
	d.postorder = append(d.postorder, vertex)
}

func (d *DirectedGraph) dfsForCyclicPath(vertex int, inStack *[]bool, path *[]int) {
	d.visited[vertex] = true

	// put this vertex in stack.
	(*inStack)[vertex] = true
	adjs, _ := d.GetAdjacentVertices(vertex)
	for _, adj := range adjs {
		// We have the path found, so we should return only.
		if len(*path) != 0 {
			return
		}
		if !d.visited[adj] {
			d.pathTo[adj] = vertex
			d.dfsForCyclicPath(adj, inStack, path)
		} else if (*inStack)[adj] {
			// We've got a loop.
			for v := vertex; v != adj; v = d.pathTo[v] {
				// fmt.Println("endless loop happending", v, d.pathTo[v])
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
func (d *DirectedGraph) GetStronglyConnectedComponent() (stronglyConnectedComponent [][]int) {
	stronglyConnectedComponent = make([][]int, 0)

	// Get the reversed graph.
	graph := d.Reverse()

	// Get the post-reverse order of the reversed graph.
	_, _, postReversed := graph.GetOrder()

	d.visited = make([]bool, d.vertexCount)
	d.pathTo = make([]int, d.vertexCount)
	// Do a DFS using the postReversed order
	for _, v := range postReversed {
		if !d.visited[v] {
			vertices := d.dfsRecursively(v, &d.visited)
			stronglyConnectedComponent = append(stronglyConnectedComponent, vertices)
		}
	}

	return
}
