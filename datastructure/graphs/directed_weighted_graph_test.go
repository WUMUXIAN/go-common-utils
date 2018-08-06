package graphs

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDirectedWeightedGraph(t *testing.T) {
	directedWeightedGraph := NewDirectedWeightedGraph(6)
	Convey("Create A New Directed Weighted Graph With 6 Verteices", t, func() {
		Convey("The Created Graph Should Have 6 Vertices And 0 Edges", func() {
			So(directedWeightedGraph.GetVertexCount(), ShouldEqual, 6)
			So(directedWeightedGraph.GetEdgeCount(), ShouldEqual, 0)
		})
	})
	Convey("Add An Edge Between Wrong Vertices", t, func() {
		err := directedWeightedGraph.AddEdge(-1, 2, 0)
		Convey("One Vertex Has Index < 0, Should Fail With Vertex Not Found", func() {
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		err = directedWeightedGraph.AddEdge(0, 7, 0)
		Convey("One Vertex Has Index >= VertexCount, Should Also Fail With Vertex Not Found", func() {
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
	})
	Convey("Add An Edge Between Correct Vertices", t, func() {
		Convey("Add Edge Between 0 And 1 Should Be Successful", func() {
			err := directedWeightedGraph.AddEdge(0, 1, 2)
			So(err, ShouldBeNil)
			So(directedWeightedGraph.GetEdgeCount(), ShouldEqual, 1)
		})
		Convey("Add Edge Between 2 And 5 Should Be Successful", func() {
			err := directedWeightedGraph.AddEdge(2, 5, 2)
			So(err, ShouldBeNil)
			So(directedWeightedGraph.GetEdgeCount(), ShouldEqual, 2)
		})
		Convey("Add Edge Between 1 And 4 Should Be Successful", func() {
			err := directedWeightedGraph.AddEdge(1, 4, 3)
			So(err, ShouldBeNil)
			So(directedWeightedGraph.GetEdgeCount(), ShouldEqual, 3)
		})
		Convey("Add The Rest Of The Edges Should Be Successful", func() {
			directedWeightedGraph.AddEdge(0, 2, 4)
			directedWeightedGraph.AddEdge(0, 3, 5)
			directedWeightedGraph.AddEdge(0, 4, 6)
			directedWeightedGraph.AddEdge(1, 5, 7)
			directedWeightedGraph.AddEdge(3, 5, 8)
			directedWeightedGraph.AddEdge(4, 5, 9)
			So(directedWeightedGraph.GetEdgeCount(), ShouldEqual, 9)
		})
	})
	Convey("Get Adjacent Vertices And Degree", t, func() {
		Convey("Get Adjacent Vertices And Degree For Wrong Vertex", func() {
			_, err := directedWeightedGraph.GetAdjacentVertices(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedWeightedGraph.GetAdjacentVertices(8)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedWeightedGraph.GetVertexInDegree(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedWeightedGraph.GetVertexInDegree(8)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedWeightedGraph.GetVertexOutDegree(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedWeightedGraph.GetVertexOutDegree(8)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		for i := 0; i < directedWeightedGraph.GetVertexCount(); i++ {
			Convey(fmt.Sprintf("Get Adjacent And Degree For Vertex %d Should Be Successful", i), func() {
				adjs, err := directedWeightedGraph.GetAdjacentVertices(i)
				So(err, ShouldBeNil)
				inDegree, err := directedWeightedGraph.GetVertexInDegree(i)
				So(err, ShouldBeNil)
				outDegree, err := directedWeightedGraph.GetVertexOutDegree(i)
				So(err, ShouldBeNil)
				switch i {
				case 0:
					So(adjs, ShouldResemble, []int{1, 2, 3, 4})
					So(inDegree, ShouldEqual, 0)
					So(outDegree, ShouldEqual, 4)
				case 1:
					So(adjs, ShouldResemble, []int{4, 5})
					So(inDegree, ShouldEqual, 1)
					So(outDegree, ShouldEqual, 2)
				case 2:
					So(adjs, ShouldResemble, []int{5})
					So(inDegree, ShouldEqual, 1)
					So(outDegree, ShouldEqual, 1)
				case 3:
					So(adjs, ShouldResemble, []int{5})
					So(inDegree, ShouldEqual, 1)
					So(outDegree, ShouldEqual, 1)
				case 4:
					So(adjs, ShouldResemble, []int{5})
					So(inDegree, ShouldEqual, 2)
					So(outDegree, ShouldEqual, 1)
				case 5:
					So(adjs, ShouldResemble, []int(nil))
					So(inDegree, ShouldEqual, 4)
					So(outDegree, ShouldEqual, 0)
				}
			})
		}
	})
	Convey("Get Edges", t, func() {
		edges := directedWeightedGraph.GetEdges()
		Convey("Print Edges Should Work", func() {
			for i, edge := range edges {
				switch i {
				case 0:
					So(edge.Print(), ShouldEqual, "Edge (0 -> 1), weight: 2.00")
				case 1:
					So(edge.Print(), ShouldEqual, "Edge (0 -> 2), weight: 4.00")
				case 2:
					So(edge.Print(), ShouldEqual, "Edge (0 -> 3), weight: 5.00")
				case 3:
					So(edge.Print(), ShouldEqual, "Edge (0 -> 4), weight: 6.00")
				case 4:
					So(edge.Print(), ShouldEqual, "Edge (1 -> 4), weight: 3.00")
				case 5:
					So(edge.Print(), ShouldEqual, "Edge (1 -> 5), weight: 7.00")
				case 6:
					So(edge.Print(), ShouldEqual, "Edge (2 -> 5), weight: 2.00")
				case 7:
					So(edge.Print(), ShouldEqual, "Edge (3 -> 5), weight: 8.00")
				case 8:
					So(edge.Print(), ShouldEqual, "Edge (4 -> 5), weight: 9.00")
				}
			}
			So(edges[1].GetFrom(), ShouldEqual, 0)
			So(edges[2].GetTo(), ShouldEqual, 3)
			So(edges[0].Compare(edges[1]), ShouldEqual, -1)
			So(edges[0].Compare(edges[6]), ShouldEqual, 0)
			So(edges[7].Compare(edges[6]), ShouldEqual, 1)
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
		res := directedWeightedGraph.Print()
		Convey("Should Be Successful", func() {
			So(res, ShouldEqual, `Vertex Count: 6, Edge Count: 9
Vertex 0: [{0 1 2} {0 2 4} {0 3 5} {0 4 6}]
Vertex 1: [{1 4 3} {1 5 7}]
Vertex 2: [{2 5 2}]
Vertex 3: [{3 5 8}]
Vertex 4: [{4 5 9}]
Vertex 5: []
`)
		})
	})
	Convey("Reverse The Graph And Print", t, func() {
		reversedGraph := directedWeightedGraph.Reverse()
		Convey("Should Be Successful", func() {
			So(reversedGraph.Print(), ShouldEqual, `Vertex Count: 6, Edge Count: 9
Vertex 0: []
Vertex 1: [{1 0 2}]
Vertex 2: [{2 0 4}]
Vertex 3: [{3 0 5}]
Vertex 4: [{4 0 6} {4 1 3}]
Vertex 5: [{5 1 7} {5 2 2} {5 3 8} {5 4 9}]
`)
		})
	})
	Convey("Depth First Searcch The Graph", t, func() {
		Convey("Starting From A Non-Existed Vertex Should Return Error", func() {
			_, err := directedWeightedGraph.DFS(10)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedWeightedGraph.DFSRecursively(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 1 Should Return 1, 4, 5", func() {
			vertices, err := directedWeightedGraph.DFS(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 4, 5})

			vertices, err = directedWeightedGraph.DFSRecursively(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 4, 5})
		})
	})

	Convey("Get The Path Between Two Vertices In The Graph", t, func() {
		Convey("Use Non-Existed Vertices Should Return Error", func() {
			_, err := directedWeightedGraph.GetDFSPath(-1, 2)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedWeightedGraph.GetDFSPath(0, 10)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 0 And End At 5 Should Return 0, 1, 4, 5", func() {
			vertices, err := directedWeightedGraph.GetDFSPath(0, 5)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{0, 1, 4, 5})
		})
		Convey("Starting From Vertex 2 And End At 3 Should Return No Path", func() {
			vertices, err := directedWeightedGraph.GetDFSPath(2, 3)
			So(err, ShouldBeError, errors.New("path not found"))
			So(vertices, ShouldResemble, []int(nil))
		})
		Convey("Some Graphs There's No Path Between Two Vertex", func() {
			dwGraph := NewDirectedWeightedGraph(3)
			dwGraph.AddEdge(0, 1, 0)
			_, err := dwGraph.GetDFSPath(0, 2)
			So(err, ShouldBeError, errors.New("path not found"))
		})
	})

	Convey("Breadth First Searcch The Graph", t, func() {
		Convey("Starting From A Non-Existed Vertex Should Return Error", func() {
			_, err := directedWeightedGraph.BFS(10)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedWeightedGraph.BFS(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 1 Should Return 1, 4, 5", func() {
			vertices, err := directedWeightedGraph.BFS(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 4, 5})
		})
	})

	Convey("Get The BFS Path Between Two Vertices In The Graph", t, func() {
		Convey("Use Non-Existed Vertices Should Return Error", func() {
			_, err := directedWeightedGraph.GetBFSPath(-1, 2)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedWeightedGraph.GetBFSPath(0, 10)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 0 And End At 5 Should Return 0, 1, 5", func() {
			vertices, err := directedWeightedGraph.GetBFSPath(0, 5)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{0, 1, 5})
		})
		Convey("Starting From Vertex 2 And End At 5 Should Return 2, 5", func() {
			vertices, err := directedWeightedGraph.GetBFSPath(2, 5)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{2, 5})
		})
		Convey("Some Graphs There's No Path Between Two Vertex", func() {
			dwGraph := NewDirectedWeightedGraph(3)
			dwGraph.AddEdge(0, 1, 0)
			_, err := dwGraph.GetBFSPath(0, 2)
			So(err, ShouldBeError, errors.New("path not found"))
		})
	})

	Convey("Get The Connected Component Of A Graph", t, func() {
		Convey("Use A Fully Connected Graph", func() {
			graph := NewDirectedWeightedGraph(4)
			graph.AddEdge(0, 1, 0)
			graph.AddEdge(1, 2, 1)
			graph.AddEdge(2, 3, 2)
			So(graph.GetConnectedComponents(), ShouldResemble, [][]int{{0, 1, 2, 3}})
		})
		Convey("Use A Graph That Has Two Sub-Graphes", func() {
			graph := NewDirectedWeightedGraph(4)
			graph.AddEdge(0, 1, 0)
			graph.AddEdge(2, 3, 1)
			So(graph.GetConnectedComponents(), ShouldResemble, [][]int{[]int{0, 1}, []int{2, 3}})
		})
		Convey("Use A Graph That Has Three Sub-Graphes", func() {
			graph := NewDirectedWeightedGraph(4)
			graph.AddEdge(0, 1, 0)
			So(graph.GetConnectedComponents(), ShouldResemble, [][]int{[]int{0, 1}, []int{2}, []int{3}})
		})
	})

	Convey("Get Cyclic Path In A Graph", t, func() {
		Convey("Use A Graph That Has Normal Cyclic Path", func() {
			graph := NewDirectedWeightedGraph(4)
			graph.AddEdge(0, 1, 0)
			graph.AddEdge(1, 2, 1)
			graph.AddEdge(2, 3, 2)
			graph.AddEdge(3, 0, 3)
			So(graph.GetCyclicPath(), ShouldResemble, []int{0, 1, 2, 3, 0})
		})
		Convey("Use A Graph That Also Has Normal Cyclic Path", func() {
			graph := NewDirectedWeightedGraph(6)
			graph.AddEdge(0, 1, 0)
			graph.AddEdge(1, 2, 1)
			graph.AddEdge(2, 3, 2)
			graph.AddEdge(3, 4, 3)
			graph.AddEdge(3, 5, 4)
			graph.AddEdge(4, 5, 5)
			graph.AddEdge(4, 2, 6)
			So(graph.GetCyclicPath(), ShouldResemble, []int{2, 3, 4, 2})
		})
		Convey("Use A Graph That Does Not Have Normal Cyclic Path", func() {
			graph := NewDirectedWeightedGraph(4)
			graph.AddEdge(0, 1, 0)
			graph.AddEdge(1, 2, 1)
			graph.AddEdge(2, 3, 2)
			So(graph.GetCyclicPath(), ShouldResemble, []int(nil))
		})
	})

	Convey("Get Bipartite Parts Of A Graph", t, func() {
		Convey("Use A Bipartite Graph", func() {
			graph := NewDirectedWeightedGraph(4)
			graph.AddEdge(0, 1, 0)
			graph.AddEdge(1, 2, 1)
			graph.AddEdge(0, 3, 2)
			graph.AddEdge(3, 2, 3)
			So(graph.GetBipartiteParts(), ShouldResemble, [][]int{[]int{1, 3}, []int{0, 2}})
		})
		Convey("Use A Non-Bipartite Graph", func() {
			graph := NewDirectedWeightedGraph(4)
			graph.AddEdge(0, 1, 0)
			graph.AddEdge(1, 2, 1)
			graph.AddEdge(1, 3, 2)
			graph.AddEdge(2, 3, 3)
			So(graph.GetBipartiteParts(), ShouldBeEmpty)
		})
	})

	Convey("Get Order Of The Graph", t, func() {
		graph := NewDirectedWeightedGraph(5)
		graph.AddEdge(0, 1, 0)
		graph.AddEdge(1, 4, 1)
		graph.AddEdge(0, 2, 2)
		graph.AddEdge(2, 4, 3)
		graph.AddEdge(0, 3, 4)
		graph.AddEdge(3, 4, 5)
		pre, post, reversePost := graph.GetOrder()
		Convey("Preorder Should Be 0, 1, 4, 2, 3", func() {
			So(pre, ShouldResemble, []int{0, 1, 4, 2, 3})
		})
		Convey("PostOrder Should Be 4, 1, 2, 3, 0", func() {
			So(post, ShouldResemble, []int{4, 1, 2, 3, 0})
		})
		Convey("TopologyOrder (ReversePostOrder) Should Be 0, 3, 2, 1, 4", func() {
			So(reversePost, ShouldResemble, []int{0, 3, 2, 1, 4})
		})

		// Test again the topology order is correct.
		graph = NewDirectedWeightedGraph(8)
		graph.AddEdge(0, 1, 0)
		graph.AddEdge(1, 2, 1)
		graph.AddEdge(2, 5, 2)
		graph.AddEdge(5, 7, 3)
		graph.AddEdge(1, 3, 4)
		graph.AddEdge(3, 5, 5)
		graph.AddEdge(0, 4, 6)
		graph.AddEdge(4, 6, 7)
		graph.AddEdge(6, 7, 8)
		topology := graph.GetTopologyOrder()
		Convey("TopologyOrder (ReversePostOrder) Should Be 0, 4, 6, 1, 3, 2, 5, 7", func() {
			So(topology, ShouldResemble, []int{0, 4, 6, 1, 3, 2, 5, 7})
		})

		// Test when there's a cyclic path in the graph.
		graph = NewDirectedWeightedGraph(4)
		graph.AddEdge(0, 1, 0)
		graph.AddEdge(1, 2, 1)
		graph.AddEdge(2, 3, 2)
		graph.AddEdge(3, 1, 1)
		topology = graph.GetTopologyOrder()
		Convey("When Graph Has Cyclic Path, Topology Order Should Return Nil", func() {
			So(topology, ShouldBeNil)
		})
	})

	Convey("Get Strongly Connected Components Of A Graph", t, func() {
		graph := NewDirectedWeightedGraph(5)
		graph.AddEdge(0, 1, 1)
		graph.AddEdge(1, 4, 1)
		graph.AddEdge(0, 2, 1)
		graph.AddEdge(2, 4, 1)
		graph.AddEdge(0, 3, 1)
		graph.AddEdge(3, 4, 1)
		Convey("Strongly Connected Components Should [[4] [3] [2] [1] [0]]", func() {
			So(graph.GetStronglyConnectedComponent(), ShouldResemble, [][]int{[]int{4}, []int{3}, []int{2}, []int{1}, []int{0}})
		})

		graph = NewDirectedWeightedGraph(2)
		graph.AddEdge(0, 1, 1)
		graph.AddEdge(1, 0, 1)
		Convey("Strongly Connected Components Should [[0, 1]]", func() {
			So(graph.GetStronglyConnectedComponent(), ShouldResemble, [][]int{[]int{0, 1}})
		})

		graph = NewDirectedWeightedGraph(4)
		graph.AddEdge(0, 1, 1)
		graph.AddEdge(1, 2, 1)
		graph.AddEdge(2, 3, 1)
		graph.AddEdge(2, 0, 1)
		graph.AddEdge(3, 0, 1)
		Convey("Strongly Connected Components Should [[0, 1, 2, 3]]", func() {
			So(graph.GetStronglyConnectedComponent(), ShouldResemble, [][]int{[]int{0, 1, 2, 3}})
		})

		graph = NewDirectedWeightedGraph(4)
		graph.AddEdge(0, 1, 1)
		graph.AddEdge(1, 2, 1)
		graph.AddEdge(2, 0, 1)
		Convey("Strongly Connected Components Should [[3] [0, 1, 2]]", func() {
			So(graph.GetStronglyConnectedComponent(), ShouldResemble, [][]int{[]int{3}, []int{0, 1, 2}})
		})
	})

	Convey("Get Transtive Clousure From A Graph", t, func() {
		graph := NewDirectedWeightedGraph(4)
		graph.AddEdge(0, 1, 1)
		graph.AddEdge(1, 2, 1)
		graph.AddEdge(2, 0, 1)
		compoents := graph.GetStronglyConnectedComponent()
		tc := NewTransitiveClousure(graph)
		Convey("Transtive Clousure Should Be Correct", func() {
			So(tc.graph, ShouldResemble, [][]bool{[]bool{true, true, true, false}, []bool{true, true, true, false}, []bool{true, true, true, false}, []bool{false, false, false, true}})
		})
		Convey("Check Against Connected Compoents Should Be Correct", func() {
			So(compoents, ShouldResemble, [][]int{[]int{3}, []int{0, 1, 2}})
			So(tc.graph[compoents[1][0]][compoents[1][2]], ShouldBeTrue)
			So(tc.graph[compoents[1][2]][compoents[1][0]], ShouldBeTrue)
			So(tc.graph[compoents[1][2]][compoents[0][0]], ShouldBeFalse)
			So(tc.graph[compoents[1][1]][compoents[0][0]], ShouldBeFalse)
			So(tc.graph[compoents[1][0]][compoents[0][0]], ShouldBeFalse)
		})
	})
}
