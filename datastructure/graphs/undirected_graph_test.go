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

}
