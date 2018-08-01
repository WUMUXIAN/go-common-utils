package graphs

// Graph defines a graph interface
type Graph interface {
	GetVertexCount() int
	GetEdgeCount() int
	AddEdge(vertex1, vertex2 int) error
	GetAdjacentVertices(vertex int) ([]int, error)
	GetVertexDegree(vertex int) (int, error)
}
