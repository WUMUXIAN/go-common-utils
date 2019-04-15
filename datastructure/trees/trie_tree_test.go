package trees

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrieTree(t *testing.T) {
	trieTree := NewTrieTree()
	Convey("New Trie Tree Should Contain No Strings\n", t, func() {
		So(trieTree.GetAll(), ShouldResemble, []string{})
	})

	trieTree.Insert("idea")
	trieTree.Insert("beam")

	Convey("Inserting Should Work As Expected\n", t, func() {
		So(trieTree.GetAll(), ShouldResemble, []string{"beam", "idea"})
	})
	trieTree.Insert("app")
	trieTree.Insert("apple")
	trieTree.Insert("apple")
	trieTree.Insert("attitude")
	trieTree.Insert("attitude")
	trieTree.Insert("attitude")
	trieTree.Insert("application")
	trieTree.Insert("application")
	trieTree.Insert("application")
	trieTree.Insert("application")

	Convey("Inserting Should Work As Expected\n", t, func() {
		So(trieTree.GetAll(), ShouldResemble, []string{"app", "apple", "application", "attitude", "beam", "idea"})
	})

	Convey("Search Should Work As Expected\n", t, func() {
		So(trieTree.Search("app"), ShouldEqual, true)
		So(trieTree.Search("apple"), ShouldEqual, true)
		So(trieTree.Search("attitude"), ShouldEqual, true)
		So(trieTree.Search("application"), ShouldEqual, true)
		So(trieTree.Search("ap"), ShouldEqual, false)
		So(trieTree.Search(""), ShouldEqual, false)
		So(trieTree.Search("attit"), ShouldEqual, false)
		So(trieTree.Search("applicationa"), ShouldEqual, false)
		So(trieTree.Search("appel"), ShouldEqual, false)
	})

	Convey("Count Should Work As Expected\n", t, func() {
		So(trieTree.Count("app"), ShouldEqual, 1)
		So(trieTree.Count("apple"), ShouldEqual, 2)
		So(trieTree.Count("attitude"), ShouldEqual, 3)
		So(trieTree.Count("application"), ShouldEqual, 4)
		So(trieTree.Count("ap"), ShouldEqual, 0)
		So(trieTree.Count(""), ShouldEqual, 0)
		So(trieTree.Count("attit"), ShouldEqual, 0)
		So(trieTree.Count("applicationa"), ShouldEqual, 0)
		So(trieTree.Count("appel"), ShouldEqual, 0)
	})

	Convey("Prefix Searching Work As Expected\n", t, func() {
		So(trieTree.StartsWith(""), ShouldResemble, trieTree.GetAll())
		So(trieTree.StartsWith("a"), ShouldResemble, []string{"app", "apple", "application", "attitude"})
		So(trieTree.StartsWith("ap"), ShouldResemble, []string{"app", "apple", "application"})
		So(trieTree.StartsWith("at"), ShouldResemble, []string{"attitude"})
		So(trieTree.StartsWith("applicat"), ShouldResemble, []string{"application"})
		So(trieTree.StartsWith("be"), ShouldResemble, []string{"beam"})
	})
}
