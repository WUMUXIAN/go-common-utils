package trees

import (
	"testing"

	"github.com/WUMUXIAN/go-common-utils/datastructure/shared"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBinaryTree(t *testing.T) {
	tree := &BinaryTree{
		Root:       nil,
		Comparator: shared.IntComparator,
	}

	Convey("Traverse Breadth First (Level by Level) For Empty Tree Should Give Empty Result", t, func() {
		So(tree.BreadthFirstTraverse(), ShouldBeEmpty)
	})
	Convey("Traverse Depth First For Empty Tree Should Give Empty Result", t, func() {
		So(tree.DepthFirstTraverse(), ShouldBeEmpty)
	})

	tree.Root = &BinaryTreeNode{
		Key: 5,
	}
	tree.Root.Left = &BinaryTreeNode{
		Key: 4,
	}
	tree.Root.Right = &BinaryTreeNode{
		Key: 3,
	}
	tree.Root.Left.Left = &BinaryTreeNode{
		Key: 2,
	}
	tree.Root.Right.Right = &BinaryTreeNode{
		Key: 1,
	}
	Convey("Traverse Breadth First (Level by Level) Should Give Corrent Result", t, func() {
		So(tree.BreadthFirstTraverse(), ShouldResemble, []interface{}{5, 4, 3, 2, 1})
	})

	Convey("Traverse Depth First Should Give Corrent Result", t, func() {
		So(tree.DepthFirstTraverse(), ShouldResemble, []interface{}{5, 4, 2, 3, 1})
	})
}
