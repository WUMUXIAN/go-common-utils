package graphs

import (
	"errors"
	"fmt"
	"math"

	"github.com/TectusDreamlab/go-common-utils/datastructure/shared"
	"github.com/TectusDreamlab/go-common-utils/datastructure/trees"
)

// DirectedWeightedGraph defines a directed wegithed graph
type DirectedWeightedGraph struct {
	DirectedGraph
	adjacentEdges [][]DirectedWeightedEdge
	weightTo      []float64
	edgeTo        []*DirectedWeightedEdge
}

// NewDirectedWeightedGraph initalises a new directed weighted graph with vertexCount vertices.
func NewDirectedWeightedGraph(vertexCount int) *DirectedWeightedGraph {
	return &DirectedWeightedGraph{
		DirectedGraph{
			UnDirectedGraph{vertexCount: vertexCount, adjacentVertices: make([][]int, vertexCount)},
			make([]int, vertexCount), nil, nil, nil,
		},
		make([][]DirectedWeightedEdge, vertexCount),
		make([]float64, vertexCount),
		make([]*DirectedWeightedEdge, vertexCount),
	}
}

// GetAdjacentVertices gets all adjacent vertices for a given vertex
func (d *DirectedWeightedGraph) GetAdjacentVertices(vertex int) ([]int, error) {
	return d.DirectedGraph.GetAdjacentVertices(vertex)
}

// GetAdjacentEdges gets all adjacent edges for a given vertex
func (d *DirectedWeightedGraph) GetAdjacentEdges(vertex int) ([]DirectedWeightedEdge, error) {
	if !d.isVertexValid(vertex) {
		return nil, errors.New("vertex not found")
	}
	return d.adjacentEdges[vertex], nil
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
func (d *DirectedWeightedGraph) Print() string {
	res := ""
	res += fmt.Sprintf("Vertex Count: %d, Edge Count: %d\n", d.vertexCount, d.edgeCount)
	for vertex, adjacentEdges := range d.adjacentEdges {
		res += fmt.Sprintf("Vertex %d: %v\n", vertex, adjacentEdges)
	}
	return res
}

// DijkstraShortestPath gets the shortest path (mimimum weight) from a start vertex to an end vertex
func (d *DirectedWeightedGraph) DijkstraShortestPath(s, v int) (path []DirectedWeightedEdge, distance float64, err error) {
	if !d.isVertexValid(s) || !d.isVertexValid(v) {
		err = errors.New("vertext not found")
		return
	}

	d.edgeTo = make([]*DirectedWeightedEdge, d.vertexCount)
	d.weightTo = make([]float64, d.vertexCount)
	for i := range d.weightTo {
		d.weightTo[i] = math.MaxFloat64
	}
	d.weightTo[s] = float64(0)
	indexedPriorityQueue := trees.NewIndexedPriorityQueue(d.vertexCount, trees.HeapTypeMin, shared.Float64Comparator)
	indexedPriorityQueue.Insert(s, d.weightTo[s])
	for !indexedPriorityQueue.IsEmpty() {
		var targetV int
		targetV, _, err = indexedPriorityQueue.Pop()
		if err != nil {
			return
		}
		var adjs []DirectedWeightedEdge
		adjs, err = d.GetAdjacentEdges(targetV)
		if err != nil {
			return
		}
		// relax all the edges of this vertex
		for _, e := range adjs {
			// s -> targetV -> w weight
			w := e.GetTo()
			weight := d.weightTo[targetV] + e.GetWeight()
			if d.weightTo[w] > weight {
				d.weightTo[w] = weight
				d.edgeTo[w] = &e
				if indexedPriorityQueue.Contains(w) {
					indexedPriorityQueue.ChangeValue(w, weight)
				} else {
					indexedPriorityQueue.Insert(w, weight)
				}
			}
		}
	}
	if d.weightTo[v] == math.MaxFloat64 {
		err = errors.New("path not found")
		return
	}
	distance = d.weightTo[v]
	e := d.edgeTo[v]
	for e != nil {
		path = append([]DirectedWeightedEdge{*e}, path...)
		e = d.edgeTo[e.GetFrom()]
	}
	return
}
