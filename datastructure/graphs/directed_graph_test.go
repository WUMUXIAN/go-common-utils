package graphs

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDirectedGraph(t *testing.T) {
	directedGraph := NewDirectedGraph(6)
	Convey("Create A New Directed Graph With 6 Verteices", t, func() {
		Convey("The Created Graph Should Have 6 Vertices And 0 Edges", func() {
			So(directedGraph.GetVertexCount(), ShouldEqual, 6)
			So(directedGraph.GetEdgeCount(), ShouldEqual, 0)
		})
	})
	Convey("Add An Edge Between Wrong Vertices", t, func() {
		err := directedGraph.AddEdge(-1, 2)
		Convey("One Vertex Has Index < 0, Should Fail With Vertex Not Found", func() {
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		err = directedGraph.AddEdge(0, 7)
		Convey("One Vertex Has Index >= VertexCount, Should Also Fail With Vertex Not Found", func() {
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
	})
	Convey("Add An Edge Between Correct Vertices", t, func() {
		Convey("Add Edge Between 0 And 1 Should Be Successful", func() {
			err := directedGraph.AddEdge(0, 1)
			So(err, ShouldBeNil)
			So(directedGraph.GetEdgeCount(), ShouldEqual, 1)
		})
		Convey("Add Edge Between 2 And 5 Should Be Successful", func() {
			err := directedGraph.AddEdge(2, 5)
			So(err, ShouldBeNil)
			So(directedGraph.GetEdgeCount(), ShouldEqual, 2)
		})
		Convey("Add Edge Between 1 And 4 Should Be Successful", func() {
			err := directedGraph.AddEdge(1, 4)
			So(err, ShouldBeNil)
			So(directedGraph.GetEdgeCount(), ShouldEqual, 3)
		})
		Convey("Add The Rest Of The Edges Should Be Successful", func() {
			directedGraph.AddEdge(0, 2)
			directedGraph.AddEdge(0, 3)
			directedGraph.AddEdge(0, 4)
			directedGraph.AddEdge(1, 5)
			directedGraph.AddEdge(3, 5)
			directedGraph.AddEdge(4, 5)
			So(directedGraph.GetEdgeCount(), ShouldEqual, 9)
		})
	})
	Convey("Get Adjacent Vertices And Degree", t, func() {
		Convey("Get Adjacent Vertices And Degree For Wrong Vertex", func() {
			_, err := directedGraph.GetAdjacentVertices(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedGraph.GetAdjacentVertices(8)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedGraph.GetVertexInDegree(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedGraph.GetVertexInDegree(8)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedGraph.GetVertexOutDegree(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedGraph.GetVertexOutDegree(8)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		for i := 0; i < directedGraph.GetVertexCount(); i++ {
			Convey(fmt.Sprintf("Get Adjacent And Degree For Vertex %d Should Be Successful", i), func() {
				adjs, err := directedGraph.GetAdjacentVertices(i)
				So(err, ShouldBeNil)
				inDegree, err := directedGraph.GetVertexInDegree(i)
				So(err, ShouldBeNil)
				outDegree, err := directedGraph.GetVertexOutDegree(i)
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
	Convey("Print The Graph", t, func() {
		res := directedGraph.Print()
		Convey("Should Be Successful", func() {
			So(res, ShouldEqual, `Vertex Count: 6, Edge Count: 9
Vertex 0: [1 2 3 4]
Vertex 1: [4 5]
Vertex 2: [5]
Vertex 3: [5]
Vertex 4: [5]
Vertex 5: []
`)
		})
	})
	Convey("Reverse The Graph And Print", t, func() {
		reversedGraph := directedGraph.Reverse()
		Convey("Should Be Successful", func() {
			So(reversedGraph.Print(), ShouldEqual, `Vertex Count: 6, Edge Count: 9
Vertex 0: []
Vertex 1: [0]
Vertex 2: [0]
Vertex 3: [0]
Vertex 4: [0 1]
Vertex 5: [1 2 3 4]
`)
		})
	})
	Convey("Depth First Searcch The Graph", t, func() {
		Convey("Starting From A Non-Existed Vertex Should Return Error", func() {
			_, err := directedGraph.DFS(10)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedGraph.DFSRecursively(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 1 Should Return 1, 4, 5", func() {
			vertices, err := directedGraph.DFS(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 4, 5})

			vertices, err = directedGraph.DFSRecursively(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 4, 5})
		})
	})

	Convey("Get The Path Between Two Vertices In The Graph", t, func() {
		Convey("Use Non-Existed Vertices Should Return Error", func() {
			_, err := directedGraph.GetDFSPath(-1, 2)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedGraph.GetDFSPath(0, 10)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 0 And End At 5 Should Return 0, 1, 4, 5", func() {
			vertices, err := directedGraph.GetDFSPath(0, 5)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{0, 1, 4, 5})
		})
		Convey("Starting From Vertex 2 And End At 3 Should Return No Path", func() {
			vertices, err := directedGraph.GetDFSPath(2, 3)
			So(err, ShouldBeError, errors.New("path not found"))
			So(vertices, ShouldResemble, []int(nil))
		})
		Convey("Some Graphs There's No Path Between Two Vertex", func() {
			udGraph := NewDirectedGraph(3)
			udGraph.AddEdge(0, 1)
			_, err := udGraph.GetDFSPath(0, 2)
			So(err, ShouldBeError, errors.New("path not found"))
		})
	})

	Convey("Breadth First Searcch The Graph", t, func() {
		Convey("Starting From A Non-Existed Vertex Should Return Error", func() {
			_, err := directedGraph.BFS(10)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedGraph.BFS(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 1 Should Return 1, 4, 5", func() {
			vertices, err := directedGraph.BFS(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 4, 5})
		})
	})

	Convey("Get The BFS Path Between Two Vertices In The Graph", t, func() {
		Convey("Use Non-Existed Vertices Should Return Error", func() {
			_, err := directedGraph.GetBFSPath(-1, 2)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = directedGraph.GetBFSPath(0, 10)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 0 And End At 5 Should Return 0, 1, 5", func() {
			vertices, err := directedGraph.GetBFSPath(0, 5)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{0, 1, 5})
		})
		Convey("Starting From Vertex 2 And End At 5 Should Return 2, 5", func() {
			vertices, err := directedGraph.GetBFSPath(2, 5)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{2, 5})
		})
		Convey("Some Graphs There's No Path Between Two Vertex", func() {
			udGraph := NewDirectedGraph(3)
			udGraph.AddEdge(0, 1)
			_, err := udGraph.GetBFSPath(0, 2)
			So(err, ShouldBeError, errors.New("path not found"))
		})
	})

	Convey("Get The Connected Component Of A Graph", t, func() {
		Convey("Use A Fully Connected Graph", func() {
			graph := NewDirectedGraph(4)
			graph.AddEdge(0, 1)
			graph.AddEdge(1, 2)
			graph.AddEdge(2, 3)
			So(graph.GetConnectedComponents(), ShouldResemble, [][]int{{0, 1, 2, 3}})
		})
		Convey("Use A Graph That Has Two Sub-Graphes", func() {
			graph := NewUnDirectedGraph(4)
			graph.AddEdge(0, 1)
			graph.AddEdge(2, 3)
			So(graph.GetConnectedComponents(), ShouldResemble, [][]int{[]int{0, 1}, []int{2, 3}})
		})
		Convey("Use A Graph That Has Three Sub-Graphes", func() {
			graph := NewUnDirectedGraph(4)
			graph.AddEdge(0, 1)
			So(graph.GetConnectedComponents(), ShouldResemble, [][]int{[]int{0, 1}, []int{2}, []int{3}})
		})
	})

	Convey("Get Cyclic Path In A Graph", t, func() {
		Convey("Use A Graph That Has Normal Cyclic Path", func() {
			graph := NewDirectedGraph(4)
			graph.AddEdge(0, 1)
			graph.AddEdge(1, 2)
			graph.AddEdge(2, 3)
			graph.AddEdge(3, 0)
			So(graph.GetCyclicPath(), ShouldResemble, []int{0, 1, 2, 3, 0})
		})
		Convey("Use A Graph That Also Has Normal Cyclic Path", func() {
			graph := NewDirectedGraph(6)
			graph.AddEdge(0, 1)
			graph.AddEdge(1, 2)
			graph.AddEdge(2, 3)
			graph.AddEdge(3, 4)
			graph.AddEdge(3, 5)
			graph.AddEdge(4, 5)
			graph.AddEdge(4, 2)
			So(graph.GetCyclicPath(), ShouldResemble, []int{2, 3, 4, 2})
		})
		Convey("Use A Graph That Does Not Have Normal Cyclic Path", func() {
			graph := NewDirectedGraph(4)
			graph.AddEdge(0, 1)
			graph.AddEdge(1, 2)
			graph.AddEdge(2, 3)
			So(graph.GetCyclicPath(), ShouldResemble, []int(nil))
		})
	})

	Convey("Get Bipartite Parts Of A Graph", t, func() {
		Convey("Use A Bipartite Graph", func() {
			graph := NewDirectedGraph(4)
			graph.AddEdge(0, 1)
			graph.AddEdge(1, 2)
			graph.AddEdge(0, 3)
			graph.AddEdge(3, 2)
			So(graph.GetBipartiteParts(), ShouldResemble, [][]int{[]int{1, 3}, []int{0, 2}})
		})
		Convey("Use A Non-Bipartite Graph", func() {
			graph := NewDirectedGraph(4)
			graph.AddEdge(0, 1)
			graph.AddEdge(1, 2)
			graph.AddEdge(1, 3)
			graph.AddEdge(2, 3)
			So(graph.GetBipartiteParts(), ShouldBeEmpty)
		})
	})

	Convey("Get Order Of The Graph", t, func() {
		graph := NewDirectedGraph(5)
		graph.AddEdge(0, 1)
		graph.AddEdge(1, 4)
		graph.AddEdge(0, 2)
		graph.AddEdge(2, 4)
		graph.AddEdge(0, 3)
		graph.AddEdge(3, 4)
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
		graph = NewDirectedGraph(8)
		graph.AddEdge(0, 1)
		graph.AddEdge(1, 2)
		graph.AddEdge(2, 5)
		graph.AddEdge(5, 7)
		graph.AddEdge(1, 3)
		graph.AddEdge(3, 5)
		graph.AddEdge(0, 4)
		graph.AddEdge(4, 6)
		graph.AddEdge(6, 7)
		topology := graph.GetTopologyOrder()
		Convey("TopologyOrder (ReversePostOrder) Should Be 0, 4, 6, 1, 3, 2, 5, 7", func() {
			So(topology, ShouldResemble, []int{0, 4, 6, 1, 3, 2, 5, 7})
		})

		// Test when there's a cyclic path in the graph.
		graph = NewDirectedGraph(4)
		graph.AddEdge(0, 1)
		graph.AddEdge(1, 2)
		graph.AddEdge(2, 3)
		graph.AddEdge(3, 1)
		topology = graph.GetTopologyOrder()
		Convey("When Graph Has Cyclic Path, Topology Order Should Return Nil", func() {
			So(topology, ShouldBeNil)
		})
	})

	Convey("Get Strongly Connected Components Of A Graph", t, func() {
		graph := NewDirectedGraph(5)
		graph.AddEdge(0, 1)
		graph.AddEdge(1, 4)
		graph.AddEdge(0, 2)
		graph.AddEdge(2, 4)
		graph.AddEdge(0, 3)
		graph.AddEdge(3, 4)
		Convey("Strongly Connected Components Should [[4] [3] [2] [1] [0]]", func() {
			So(graph.GetStronglyConnectedComponent(), ShouldResemble, [][]int{[]int{4}, []int{3}, []int{2}, []int{1}, []int{0}})
		})

		graph = NewDirectedGraph(2)
		graph.AddEdge(0, 1)
		graph.AddEdge(1, 0)
		Convey("Strongly Connected Components Should [[0, 1]]", func() {
			So(graph.GetStronglyConnectedComponent(), ShouldResemble, [][]int{[]int{0, 1}})
		})

		graph = NewDirectedGraph(4)
		graph.AddEdge(0, 1)
		graph.AddEdge(1, 2)
		graph.AddEdge(2, 3)
		graph.AddEdge(2, 0)
		graph.AddEdge(3, 0)
		Convey("Strongly Connected Components Should [[0, 1, 2, 3]]", func() {
			So(graph.GetStronglyConnectedComponent(), ShouldResemble, [][]int{[]int{0, 1, 2, 3}})
		})

		graph = NewDirectedGraph(4)
		graph.AddEdge(0, 1)
		graph.AddEdge(1, 2)
		graph.AddEdge(2, 0)
		Convey("Strongly Connected Components Should [[3] [0, 1, 2]]", func() {
			So(graph.GetStronglyConnectedComponent(), ShouldResemble, [][]int{[]int{3}, []int{0, 1, 2}})
		})
	})

	Convey("Get Transtive Clousure From A Graph", t, func() {
		graph := NewDirectedGraph(4)
		graph.AddEdge(0, 1)
		graph.AddEdge(1, 2)
		graph.AddEdge(2, 0)
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
