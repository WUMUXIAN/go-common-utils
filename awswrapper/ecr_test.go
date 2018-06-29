package awswrapper

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListImages(t *testing.T) {
	images := GetECRService("us-east-1").ListImages("vinspection")
	Convey("List All Vinspection Images", t, func() {
		So(images, ShouldNotBeEmpty)
		fmt.Println(images)
	})
	images = GetECRService("us-east-1").ListImages("convertor")
	Convey("List All Convertor Images", t, func() {
		So(images, ShouldNotBeEmpty)
		fmt.Println(images)
	})
	images = GetECRService("us-east-1").ListImages("not_existed")
	Convey("Try to List Non-existed Repo", t, func() {
		So(images, ShouldBeEmpty)
		fmt.Println(images)
	})
}
