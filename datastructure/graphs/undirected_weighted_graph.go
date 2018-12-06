package graphs

import (
	"errors"
	"fmt"

	"github.com/TectusDreamlab/go-common-utils/datastructure/trees"
)

// UnDirectedWeightGraph defines a undirected graph
type UnDirectedWeightGraph struct {
	UnDirectedGraph

	adjacentEdges [][]WeightedEdge
	mimimumWeight float64
}

// NewUnDirectedWeightedGraph initalises a new undirected weighted graph with vertexCount vertices.
func NewUnDirectedWeightedGraph(vertexCount int) *UnDirectedWeightGraph {
	return &UnDirectedWeightGraph{
		UnDirectedGraph{vertexCount, 0, make([][]int, vertexCount), nil, nil, nil, nil},
		make([][]WeightedEdge, vertexCount), 0,
	}
}

// AddEdge adds an edge to the graph
func (u *UnDirectedWeightGraph) AddEdge(vertex1, vertex2 int, weight float64) error {
	if u.isVertexValid(vertex1) && u.isVertexValid(vertex2) {
		u.UnDirectedGraph.AddEdge(vertex1, vertex2)
		u.adjacentEdges[vertex1] = append(u.adjacentEdges[vertex1], WeightedEdge{vertex1, vertex2, weight})
		u.adjacentEdges[vertex2] = append(u.adjacentEdges[vertex2], WeightedEdge{vertex2, vertex1, weight})
		return nil
	}
	return errors.New("vertex not found")
}

// GetEdges prints all edges.
func (u *UnDirectedWeightGraph) GetEdges() []WeightedEdge {
	res := make([]WeightedEdge, 0)
	for i := 0; i < u.vertexCount; i++ {
		edges := u.adjacentEdges[i]
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
	for vertex, adjacentEdges := range u.adjacentEdges {
		res += fmt.Sprintf("Vertex %d: %v\n", vertex, adjacentEdges)
	}
	return res
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
	edges := u.adjacentEdges[vertex]
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
	for priorityQueue.Peek() != nil {
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
