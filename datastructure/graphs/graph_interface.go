package graphs

// Graph defines a graph interface
type Graph interface {
	GetVertexCount() int
	GetEdgeCount() int
	GetAdjacentVertices(vertex int) ([]int, error)
	GetBFSPath(verext1, vertex2 int) ([]int, error)
}
