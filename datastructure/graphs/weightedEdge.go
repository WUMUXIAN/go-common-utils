package graphs

import (
	"errors"
	"fmt"
)

// WeightedEdge defines a weighted edge
type WeightedEdge struct {
	vertex1 int
	vertex2 int
	weight  float64
}

// GetWeight gets the weight on the edge
func (w *WeightedEdge) GetWeight() float64 {
	return w.weight
}

// GetVertex1 gets of the vertex of the edge.
func (w *WeightedEdge) GetVertex1() int {
	return w.vertex1
}

// GetOther gets the other vertex based on the given vertex of the edge.
func (w *WeightedEdge) GetOther(vertex int) (int, error) {
	if vertex == w.vertex1 {
		return w.vertex2, nil
	} else if vertex == w.vertex2 {
		return w.vertex1, nil
	}
	return 0, errors.New("vertex not found")
}

// Compare compares the weight of two edges.
func (w *WeightedEdge) Compare(w1 WeightedEdge) int {
	if w.GetWeight() > w1.GetWeight() {
		return 1
	} else if w.GetWeight() < w1.GetWeight() {
		return -1
	}
	return 0
}

// Print prints the edge.
func (w *WeightedEdge) Print() string {
	return fmt.Sprintf("Edge (%d, %d), weight: %0.2f", w.vertex1, w.vertex2, w.weight)
}
