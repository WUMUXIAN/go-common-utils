package graphs

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnDirectedWeightGraph(t *testing.T) {
	unDirectedWeightedGraph := NewUnDirectedWeightedGraph(6)
	Convey("Create A New UnDirected Weighted Graph With 6 Verteices", t, func() {
		Convey("The Created Graph Should Have 6 Vertices And 0 Edges", func() {
			So(unDirectedWeightedGraph.GetVertexCount(), ShouldEqual, 6)
			So(unDirectedWeightedGraph.GetEdgeCount(), ShouldEqual, 0)
		})
	})
	Convey("Add An Edge Between Wrong Vertices", t, func() {
		err := unDirectedWeightedGraph.AddEdge(-1, 2, 1)
		Convey("One Vertex Has Index < 0, Should Fail With Vertex Not Found", func() {
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		err = unDirectedWeightedGraph.AddEdge(0, 7, 2)
		Convey("One Vertex Has Index >= VertexCount, Should Also Fail With Vertex Not Found", func() {
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
	})
	Convey("Add An Edge Between Correct Vertices", t, func() {
		Convey("Add Edge Between 0 And 1 Should Be Successful", func() {
			err := unDirectedWeightedGraph.AddEdge(0, 1, 2)
			So(err, ShouldBeNil)
			So(unDirectedWeightedGraph.GetEdgeCount(), ShouldEqual, 1)
		})
		Convey("Add Edge Between 2 And 5 Should Be Successful", func() {
			err := unDirectedWeightedGraph.AddEdge(2, 5, 3)
			So(err, ShouldBeNil)
			So(unDirectedWeightedGraph.GetEdgeCount(), ShouldEqual, 2)
		})
		Convey("Add Edge Between 1 And 4 Should Be Successful", func() {
			err := unDirectedWeightedGraph.AddEdge(1, 4, 4)
			So(err, ShouldBeNil)
			So(unDirectedWeightedGraph.GetEdgeCount(), ShouldEqual, 3)
		})
		Convey("Add The Rest Of The Edges Should Be Successful", func() {
			unDirectedWeightedGraph.AddEdge(0, 2, 1)
			unDirectedWeightedGraph.AddEdge(0, 3, 2)
			unDirectedWeightedGraph.AddEdge(0, 4, 3)
			unDirectedWeightedGraph.AddEdge(1, 5, 4)
			unDirectedWeightedGraph.AddEdge(3, 5, 5)
			unDirectedWeightedGraph.AddEdge(4, 5, 6)
			So(unDirectedWeightedGraph.GetEdgeCount(), ShouldEqual, 9)
		})
	})
	Convey("Get Adjacent Vertices And Degree", t, func() {
		Convey("Get Adjacent Vertices And Degree For Wrong Vertex", func() {
			_, err := unDirectedWeightedGraph.GetAdjacentVertices(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = unDirectedWeightedGraph.GetAdjacentVertices(8)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = unDirectedWeightedGraph.GetVertexDegree(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = unDirectedWeightedGraph.GetVertexDegree(8)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		for i := 0; i < unDirectedWeightedGraph.GetVertexCount(); i++ {
			Convey(fmt.Sprintf("Get Adjacent And Degree For Vertex %d Should Be Successful", i), func() {
				adjs, err := unDirectedWeightedGraph.GetAdjacentVertices(i)
				So(err, ShouldBeNil)
				degree, err := unDirectedWeightedGraph.GetVertexDegree(i)
				So(err, ShouldBeNil)
				switch i {
				case 0:
					So(adjs, ShouldResemble, []int{1, 2, 3, 4})
					So(degree, ShouldEqual, 4)
				case 1:
					So(adjs, ShouldResemble, []int{0, 4, 5})
					So(degree, ShouldEqual, 3)
				case 2:
					So(adjs, ShouldResemble, []int{5, 0})
					So(degree, ShouldEqual, 2)
				case 3:
					So(adjs, ShouldResemble, []int{0, 5})
					So(degree, ShouldEqual, 2)
				case 4:
					So(adjs, ShouldResemble, []int{1, 0, 5})
					So(degree, ShouldEqual, 3)
				case 5:
					So(adjs, ShouldResemble, []int{2, 1, 3, 4})
					So(degree, ShouldEqual, 4)
				}
			})
		}
	})
	Convey("Get Edges", t, func() {
		edges := unDirectedWeightedGraph.GetEdges()
		Convey("Print Edges Should Work", func() {
			for i, edge := range edges {
				switch i {
				case 0:
					So(edge.Print(), ShouldEqual, "Edge (0, 1), weight: 2.00")
				case 1:
					So(edge.Print(), ShouldEqual, "Edge (0, 2), weight: 1.00")
				case 2:
					So(edge.Print(), ShouldEqual, "Edge (0, 3), weight: 2.00")
				case 3:
					So(edge.Print(), ShouldEqual, "Edge (0, 4), weight: 3.00")
				case 4:
					So(edge.Print(), ShouldEqual, "Edge (1, 4), weight: 4.00")
				case 5:
					So(edge.Print(), ShouldEqual, "Edge (1, 5), weight: 4.00")
				case 6:
					So(edge.Print(), ShouldEqual, "Edge (2, 5), weight: 3.00")
				case 7:
					So(edge.Print(), ShouldEqual, "Edge (3, 5), weight: 5.00")
				case 8:
					So(edge.Print(), ShouldEqual, "Edge (4, 5), weight: 6.00")
				}
			}
			So(edges[1].GetVertex1(), ShouldEqual, 0)
			other, err := edges[2].GetOther(0)
			So(err, ShouldBeNil)
			So(other, ShouldEqual, 3)
			other, err = edges[2].GetOther(3)
			So(err, ShouldBeNil)
			So(other, ShouldEqual, 0)
			So(edges[0].Compare(edges[1]), ShouldEqual, 1)
			So(edges[0].Compare(edges[2]), ShouldEqual, 0)
			So(edges[1].Compare(edges[2]), ShouldEqual, -1)
			_, err = edges[2].GetOther(5)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		uwg := NewUnDirectedWeightedGraph(2)
		uwg.AddEdge(0, 0, 1)
		uwg.AddEdge(0, 1, 1)
		uwg.AddEdge(1, 0, 1)
		Convey("Self Loop And Parallel Edges Should Be Handled", func() {
			edges := uwg.GetEdges()
			for i, edge := range edges {
				switch i {
				case 0:
					So(edge.Print(), ShouldEqual, "Edge (0, 0), weight: 1.00")
				case 1:
					So(edge.Print(), ShouldEqual, "Edge (0, 1), weight: 1.00")
				case 2:
					So(edge.Print(), ShouldEqual, "Edge (0, 1), weight: 1.00")
				}
			}
		})
	})
	Convey("Print The Graph", t, func() {
		res := unDirectedWeightedGraph.Print()
		Convey("Should Be Successful", func() {
			So(res, ShouldEqual, `Vertex Count: 6, Edge Count: 9
Vertex 0: [{0 1 2} {0 2 1} {0 3 2} {0 4 3}]
Vertex 1: [{1 0 2} {1 4 4} {1 5 4}]
Vertex 2: [{2 5 3} {2 0 1}]
Vertex 3: [{3 0 2} {3 5 5}]
Vertex 4: [{4 1 4} {4 0 3} {4 5 6}]
Vertex 5: [{5 2 3} {5 1 4} {5 3 5} {5 4 6}]
`)
		})
	})
	Convey("Depth First Searcch The Graph", t, func() {
		Convey("Starting From A Non-Existed Vertex Should Return Error", func() {
			_, err := unDirectedWeightedGraph.DFS(10)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = unDirectedWeightedGraph.DFSRecursively(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 1 Should Return 1, 0, 2, 5, 3, 4", func() {
			vertices, err := unDirectedWeightedGraph.DFS(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 0, 2, 5, 3, 4})

			vertices, err = unDirectedWeightedGraph.DFSRecursively(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 0, 2, 5, 3, 4})
		})
	})

	Convey("Get The Path Between Two Vertices In The Graph", t, func() {
		Convey("Use Non-Existed Vertices Should Return Error", func() {
			_, err := unDirectedWeightedGraph.GetDFSPath(-1, 2)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = unDirectedWeightedGraph.GetDFSPath(0, 10)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 0 And End At 5 Should Return 0 1 4 5", func() {
			vertices, err := unDirectedWeightedGraph.GetDFSPath(0, 5)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{0, 1, 4, 5})
		})
		Convey("Starting From Vertex 2 And End At 3 Should Return 2 5 1 3", func() {
			vertices, err := unDirectedWeightedGraph.GetDFSPath(2, 3)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{2, 5, 1, 0, 3})
		})
		Convey("Some Graphs There's No Path Between Two Vertex", func() {
			udGraph := NewUnDirectedWeightedGraph(3)
			udGraph.AddEdge(0, 1, 1)
			_, err := udGraph.GetDFSPath(0, 2)
			So(err, ShouldBeError, errors.New("path not found"))
		})
	})

	Convey("Breadth First Searcch The Graph", t, func() {
		Convey("Starting From A Non-Existed Vertex Should Return Error", func() {
			_, err := unDirectedWeightedGraph.BFS(10)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = unDirectedWeightedGraph.BFS(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 1 Should Return 1, 0, 4, 5, 2, 3", func() {
			vertices, err := unDirectedWeightedGraph.BFS(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 0, 4, 5, 2, 3})
		})
	})

	Convey("Get The BFS Path Between Two Vertices In The Graph", t, func() {
		Convey("Use Non-Existed Vertices Should Return Error", func() {
			_, err := unDirectedWeightedGraph.GetBFSPath(-1, 2)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = unDirectedWeightedGraph.GetBFSPath(0, 10)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 0 And End At 5 Should Return 0, 1, 5", func() {
			vertices, err := unDirectedWeightedGraph.GetBFSPath(0, 5)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{0, 1, 5})
		})
		Convey("Starting From Vertex 2 And End At 3 Should Return 2, 5, 3", func() {
			vertices, err := unDirectedWeightedGraph.GetBFSPath(2, 3)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{2, 5, 3})
		})
		Convey("Some Graphs There's No Path Between Two Vertex", func() {
			udGraph := NewUnDirectedWeightedGraph(3)
			udGraph.AddEdge(0, 1, 2)
			_, err := udGraph.GetBFSPath(0, 2)
			So(err, ShouldBeError, errors.New("path not found"))
		})
	})

	Convey("Get The Connected Component Of A Graph", t, func() {
		Convey("Use A Fully Connected Graph", func() {
			uGraph := NewUnDirectedWeightedGraph(4)
			uGraph.AddEdge(0, 1, 1)
			uGraph.AddEdge(1, 2, 2)
			uGraph.AddEdge(2, 0, 3)
			uGraph.AddEdge(0, 3, 4)
			So(uGraph.GetConnectedComponents(), ShouldResemble, [][]int{{0, 1, 2, 3}})
		})
		Convey("Use A Graph That Has Two Sub-Graphes", func() {
			uGraph := NewUnDirectedWeightedGraph(4)
			uGraph.AddEdge(0, 1, 1)
			uGraph.AddEdge(2, 3, 2)
			So(uGraph.GetConnectedComponents(), ShouldResemble, [][]int{[]int{0, 1}, []int{2, 3}})
		})
		Convey("Use A Graph That Has Three Sub-Graphes", func() {
			uGraph := NewUnDirectedWeightedGraph(4)
			uGraph.AddEdge(0, 1, 3)
			So(uGraph.GetConnectedComponents(), ShouldResemble, [][]int{[]int{0, 1}, []int{2}, []int{3}})
		})
	})

	Convey("Get Cyclic Path In A Graph", t, func() {
		Convey("Use A Graph With Self Loop", func() {
			uGraph := NewUnDirectedWeightedGraph(4)
			uGraph.AddEdge(0, 0, 1)
			So(uGraph.GetCyclicPath(), ShouldResemble, []int{0, 0})
		})
		Convey("Use A Graph With Parallel Edge", func() {
			uGraph := NewUnDirectedWeightedGraph(4)
			uGraph.AddEdge(0, 1, 1)
			uGraph.AddEdge(1, 0, 1)
			So(uGraph.GetCyclicPath(), ShouldResemble, []int{0, 1, 0})
		})
		Convey("Use A Graph That Has Normal Cyclic Path", func() {
			uGraph := NewUnDirectedWeightedGraph(4)
			uGraph.AddEdge(0, 1, 2)
			uGraph.AddEdge(1, 2, 3)
			uGraph.AddEdge(2, 3, 4)
			uGraph.AddEdge(3, 1, 5)
			So(uGraph.GetCyclicPath(), ShouldResemble, []int{1, 2, 3, 1})
		})
		Convey("Use A Graph That Does Not Have Normal Cyclic Path", func() {
			uGraph := NewUnDirectedWeightedGraph(4)
			uGraph.AddEdge(0, 1, 1)
			uGraph.AddEdge(1, 2, 2)
			uGraph.AddEdge(2, 3, 3)
			So(uGraph.GetCyclicPath(), ShouldResemble, []int(nil))
		})
	})

	Convey("Get Bipartite Parts Of A Graph", t, func() {
		Convey("Use A Bipartite Graph", func() {
			uGraph := NewUnDirectedWeightedGraph(4)
			uGraph.AddEdge(0, 1, 1)
			uGraph.AddEdge(1, 2, 2)
			uGraph.AddEdge(0, 3, 3)
			uGraph.AddEdge(3, 2, 4)
			So(uGraph.GetBipartiteParts(), ShouldResemble, [][]int{[]int{1, 3}, []int{0, 2}})
		})
		Convey("Use A Non-Bipartite Graph", func() {
			uGraph := NewUnDirectedWeightedGraph(4)
			uGraph.AddEdge(0, 1, 1)
			uGraph.AddEdge(1, 2, 2)
			uGraph.AddEdge(1, 3, 3)
			uGraph.AddEdge(2, 3, 4)
			So(uGraph.GetBipartiteParts(), ShouldBeEmpty)
		})
	})

	Convey("LazyPrimMinimumSpanningTree Algorithm", t, func() {
		Convey("A Fully Connected Graph Should Have Minimum Spanning Tree(s)", func() {
			uGraph := NewUnDirectedWeightedGraph(8)
			uGraph.AddEdge(0, 7, 0.16)
			uGraph.AddEdge(0, 2, 0.26)
			uGraph.AddEdge(0, 4, 0.38)
			uGraph.AddEdge(0, 6, 0.58)
			uGraph.AddEdge(1, 2, 0.36)
			uGraph.AddEdge(1, 3, 0.29)
			uGraph.AddEdge(1, 5, 0.32)
			uGraph.AddEdge(1, 7, 0.19)
			uGraph.AddEdge(2, 3, 0.17)
			uGraph.AddEdge(2, 7, 0.34)
			uGraph.AddEdge(2, 6, 0.40)
			uGraph.AddEdge(3, 6, 0.52)
			uGraph.AddEdge(4, 5, 0.35)
			uGraph.AddEdge(4, 6, 0.93)
			uGraph.AddEdge(4, 7, 0.37)
			uGraph.AddEdge(5, 7, 0.28)
			edges, weight := uGraph.LazyPrimMinimumSpanningTree()
			So(edges, ShouldResemble, []WeightedEdge{
				WeightedEdge{vertex1: 0, vertex2: 7, weight: 0.16},
				WeightedEdge{vertex1: 7, vertex2: 1, weight: 0.19},
				WeightedEdge{vertex1: 0, vertex2: 2, weight: 0.26},
				WeightedEdge{vertex1: 2, vertex2: 3, weight: 0.17},
				WeightedEdge{vertex1: 7, vertex2: 5, weight: 0.28},
				WeightedEdge{vertex1: 5, vertex2: 4, weight: 0.35},
				WeightedEdge{vertex1: 2, vertex2: 6, weight: 0.4}})
			So(weight, ShouldEqual, 1.81)
		})
		Convey("A Non-Connected Graph Should Have Minimum Spanning Tree(s)", func() {
			uGraph := NewUnDirectedWeightedGraph(4)
			uGraph.AddEdge(0, 1, 0.16)
			uGraph.AddEdge(0, 2, 0.26)
			edges, weight := uGraph.LazyPrimMinimumSpanningTree()
			So(edges, ShouldHaveLength, 0)
			So(weight, ShouldEqual, 0)
		})
		Convey("A Connected Graph With Same Weights Should Have Minimum Spanning Tree(s)", func() {
			uGraph := NewUnDirectedWeightedGraph(4)
			uGraph.AddEdge(0, 1, 0.1)
			uGraph.AddEdge(0, 2, 0.1)
			uGraph.AddEdge(0, 3, 0.1)
			edges, weight := uGraph.LazyPrimMinimumSpanningTree()
			So(edges, ShouldResemble, []WeightedEdge{
				WeightedEdge{vertex1: 0, vertex2: 1, weight: 0.1},
				WeightedEdge{vertex1: 0, vertex2: 2, weight: 0.1},
				WeightedEdge{vertex1: 0, vertex2: 3, weight: 0.1},
			})
			So(fmt.Sprintf("%0.2f", weight), ShouldEqual, "0.30")
		})
	})
}
