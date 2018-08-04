package graphs

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnDirectedGraph(t *testing.T) {
	undirectedGraph := NewUnDirectedGraph(6)
	Convey("Create A New UnDirected Graph With 6 Verteices", t, func() {
		Convey("The Created Graph Should Have 6 Vertices And 0 Edges", func() {
			So(undirectedGraph.GetVertexCount(), ShouldEqual, 6)
			So(undirectedGraph.GetEdgeCount(), ShouldEqual, 0)
		})
	})
	Convey("Add An Edge Between Wrong Vertices", t, func() {
		err := undirectedGraph.AddEdge(-1, 2)
		Convey("One Vertex Has Index < 0, Should Fail With Vertex Not Found", func() {
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		err = undirectedGraph.AddEdge(0, 7)
		Convey("One Vertex Has Index >= VertexCount, Should Also Fail With Vertex Not Found", func() {
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
	})
	Convey("Add An Edge Between Correct Vertices", t, func() {
		Convey("Add Edge Between 0 And 1 Should Be Successful", func() {
			err := undirectedGraph.AddEdge(0, 1)
			So(err, ShouldBeNil)
			So(undirectedGraph.GetEdgeCount(), ShouldEqual, 1)
		})
		Convey("Add Edge Between 2 And 5 Should Be Successful", func() {
			err := undirectedGraph.AddEdge(2, 5)
			So(err, ShouldBeNil)
			So(undirectedGraph.GetEdgeCount(), ShouldEqual, 2)
		})
		Convey("Add Edge Between 1 And 4 Should Be Successful", func() {
			err := undirectedGraph.AddEdge(1, 4)
			So(err, ShouldBeNil)
			So(undirectedGraph.GetEdgeCount(), ShouldEqual, 3)
		})
		Convey("Add The Rest Of The Edges Should Be Successful", func() {
			undirectedGraph.AddEdge(0, 2)
			undirectedGraph.AddEdge(0, 3)
			undirectedGraph.AddEdge(0, 4)
			undirectedGraph.AddEdge(1, 5)
			undirectedGraph.AddEdge(3, 5)
			undirectedGraph.AddEdge(4, 5)
			So(undirectedGraph.GetEdgeCount(), ShouldEqual, 9)
		})
	})
	Convey("Get Adjacent Vertices And Degree", t, func() {
		Convey("Get Adjacent Vertices And Degree For Wrong Vertex", func() {
			_, err := undirectedGraph.GetAdjacentVertices(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = undirectedGraph.GetAdjacentVertices(8)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = undirectedGraph.GetVertexDegree(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = undirectedGraph.GetVertexDegree(8)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		for i := 0; i < undirectedGraph.GetVertexCount(); i++ {
			Convey(fmt.Sprintf("Get Adjacent And Degree For Vertex %d Should Be Successful", i), func() {
				adjs, err := undirectedGraph.GetAdjacentVertices(i)
				So(err, ShouldBeNil)
				degree, err := undirectedGraph.GetVertexDegree(i)
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
	Convey("Print The Graph", t, func() {
		res := undirectedGraph.Print()
		Convey("Should Be Successful", func() {
			So(res, ShouldEqual, `Vertex Count: 6, Edge Count: 9
Vertex 0: [1 2 3 4]
Vertex 1: [0 4 5]
Vertex 2: [5 0]
Vertex 3: [0 5]
Vertex 4: [1 0 5]
Vertex 5: [2 1 3 4]
`)
		})
	})
	Convey("Depth First Searcch The Graph", t, func() {
		Convey("Starting From A Non-Existed Vertex Should Return Error", func() {
			_, err := undirectedGraph.DFS(10)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = undirectedGraph.DFSRecursively(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 1 Should Return 1, 0, 2, 5, 3, 4", func() {
			vertices, err := undirectedGraph.DFS(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 0, 2, 5, 3, 4})

			vertices, err = undirectedGraph.DFSRecursively(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 0, 2, 5, 3, 4})
		})
	})

	Convey("Get The Path Between Two Vertices In The Graph", t, func() {
		Convey("Use Non-Existed Vertices Should Return Error", func() {
			_, err := undirectedGraph.GetDFSPath(-1, 2)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = undirectedGraph.GetDFSPath(0, 10)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 0 And End At 5 Should Return 0 1 4 5", func() {
			vertices, err := undirectedGraph.GetDFSPath(0, 5)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{0, 1, 4, 5})
		})
		Convey("Starting From Vertex 2 And End At 3 Should Return 2 5 1 3", func() {
			vertices, err := undirectedGraph.GetDFSPath(2, 3)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{2, 5, 1, 0, 3})
		})
		Convey("Some Graphs There's No Path Between Two Vertex", func() {
			udGraph := NewUnDirectedGraph(3)
			udGraph.AddEdge(0, 1)
			_, err := udGraph.GetDFSPath(0, 2)
			So(err, ShouldBeError, errors.New("path not found"))
		})
	})

	Convey("Breadth First Searcch The Graph", t, func() {
		Convey("Starting From A Non-Existed Vertex Should Return Error", func() {
			_, err := undirectedGraph.BFS(10)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = undirectedGraph.BFS(-1)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 1 Should Return 1, 0, 4, 5, 2, 3", func() {
			vertices, err := undirectedGraph.BFS(1)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{1, 0, 4, 5, 2, 3})
		})
	})

	Convey("Get The BFS Path Between Two Vertices In The Graph", t, func() {
		Convey("Use Non-Existed Vertices Should Return Error", func() {
			_, err := undirectedGraph.GetBFSPath(-1, 2)
			So(err, ShouldBeError, errors.New("vertex not found"))
			_, err = undirectedGraph.GetBFSPath(0, 10)
			So(err, ShouldBeError, errors.New("vertex not found"))
		})
		Convey("Starting From Vertex 0 And End At 5 Should Return 0, 1, 5", func() {
			vertices, err := undirectedGraph.GetBFSPath(0, 5)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{0, 1, 5})
		})
		Convey("Starting From Vertex 2 And End At 3 Should Return 2, 5, 3", func() {
			vertices, err := undirectedGraph.GetBFSPath(2, 3)
			So(err, ShouldBeNil)
			So(vertices, ShouldResemble, []int{2, 5, 3})
		})
		Convey("Some Graphs There's No Path Between Two Vertex", func() {
			udGraph := NewUnDirectedGraph(3)
			udGraph.AddEdge(0, 1)
			_, err := udGraph.GetBFSPath(0, 2)
			So(err, ShouldBeError, errors.New("path not found"))
		})
	})

	Convey("Get The Connected Component Of A Graph", t, func() {
		Convey("Use A Fully Connected Graph", func() {
			uGraph := NewUnDirectedGraph(4)
			uGraph.AddEdge(0, 1)
			uGraph.AddEdge(1, 2)
			uGraph.AddEdge(2, 0)
			uGraph.AddEdge(0, 3)
			So(uGraph.GetConnectedComponents(), ShouldResemble, [][]int{{0, 1, 2, 3}})
		})
		Convey("Use A Graph That Has Two Sub-Graphes", func() {
			uGraph := NewUnDirectedGraph(4)
			uGraph.AddEdge(0, 1)
			uGraph.AddEdge(2, 3)
			So(uGraph.GetConnectedComponents(), ShouldResemble, [][]int{[]int{0, 1}, []int{2, 3}})
		})
		Convey("Use A Graph That Has Three Sub-Graphes", func() {
			uGraph := NewUnDirectedGraph(4)
			uGraph.AddEdge(0, 1)
			So(uGraph.GetConnectedComponents(), ShouldResemble, [][]int{[]int{0, 1}, []int{2}, []int{3}})
		})
	})

	Convey("Get Cyclic Path In A Graph", t, func() {
		Convey("Use A Graph With Self Loop", func() {
			uGraph := NewUnDirectedGraph(4)
			uGraph.AddEdge(0, 0)
			So(uGraph.GetCyclicPath(), ShouldResemble, []int{0, 0})
		})
		Convey("Use A Graph With Parallel Edge", func() {
			uGraph := NewUnDirectedGraph(4)
			uGraph.AddEdge(0, 1)
			uGraph.AddEdge(1, 0)
			So(uGraph.GetCyclicPath(), ShouldResemble, []int{0, 1, 0})
		})
		Convey("Use A Graph That Has Normal Cyclic Path", func() {
			uGraph := NewUnDirectedGraph(4)
			uGraph.AddEdge(0, 1)
			uGraph.AddEdge(1, 2)
			uGraph.AddEdge(2, 3)
			uGraph.AddEdge(3, 0)
			So(uGraph.GetCyclicPath(), ShouldResemble, []int{0, 3, 2, 1, 0})
		})
		Convey("Use A Graph That Does Not Have Normal Cyclic Path", func() {
			uGraph := NewUnDirectedGraph(4)
			uGraph.AddEdge(0, 1)
			uGraph.AddEdge(1, 2)
			uGraph.AddEdge(2, 3)
			So(uGraph.GetCyclicPath(), ShouldResemble, []int(nil))
		})
	})

	Convey("Get Bipartite Parts Of A Graph", t, func() {
		Convey("Use A Bipartite Graph", func() {
			uGraph := NewUnDirectedGraph(4)
			uGraph.AddEdge(0, 1)
			uGraph.AddEdge(1, 2)
			uGraph.AddEdge(0, 3)
			uGraph.AddEdge(3, 2)
			So(uGraph.GetBipartiteParts(), ShouldResemble, [][]int{[]int{1, 3}, []int{0, 2}})
		})
		Convey("Use A Non-Bipartite Graph", func() {
			uGraph := NewUnDirectedGraph(4)
			uGraph.AddEdge(0, 1)
			uGraph.AddEdge(1, 2)
			uGraph.AddEdge(1, 3)
			uGraph.AddEdge(2, 3)
			So(uGraph.GetBipartiteParts(), ShouldBeEmpty)
		})
	})
}
