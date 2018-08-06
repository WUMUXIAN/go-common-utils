package graphs

import (
	"fmt"
)

// DirectedWeightedEdge defines a weighted edge
type DirectedWeightedEdge struct {
	from   int
	to     int
	weight float64
}

// GetWeight gets the weight on the edge
func (d *DirectedWeightedEdge) GetWeight() float64 {
	return d.weight
}

// GetVertex1 gets of the vertex of the edge.
func (d *DirectedWeightedEdge) GetFrom() int {
	return d.from
}

// GetOther gets the other vertex based on the given vertex of the edge.
func (d *DirectedWeightedEdge) GetTo() int {
	return d.to
}

// Compare compares the weight of two edges.
func (d *DirectedWeightedEdge) Compare(d1 DirectedWeightedEdge) int {
	if d.GetWeight() > d1.GetWeight() {
		return 1
	} else if d.GetWeight() < d1.GetWeight() {
		return -1
	}
	return 0
}

// Print prints the edge.
func (d *DirectedWeightedEdge) Print() string {
	return fmt.Sprintf("Edge (%d -> %d), weight: %0.2f", d.from, d.to, d.weight)
}

// CompareDirectedEdges compares edges
func CompareDirectedEdges(e1, e2 interface{}) int {
	edge1, _ := e1.(DirectedWeightedEdge)
	edge2, _ := e2.(DirectedWeightedEdge)
	if edge1.GetWeight() > edge2.GetWeight() {
		return 1
	} else if edge1.GetWeight() < edge2.GetWeight() {
		return -1
	}
	return 0
}
