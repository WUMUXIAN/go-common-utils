package awswrapper

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestSendSESEmail(t *testing.T) {
	err := GetSESService("us-east-1").SendSESEmail("test", "Tectus DreamLab", "no-reply@tectusdreamlab.com", "wumuxian1988@gmail.com", "sample data", "html", "ServiceName", "vinspection", "Env", "test")
	Convey("Send Email Via SES Service", t, func() {
		So(err, ShouldBeNil)
	})

	err = GetSESService("us-east-1").SendSESEmail("test", "Tectus DreamLab", "no-reply@tectusdreamlab.com", "wumuxian1988@gmail.com", "sample data", "text", "ServiceName", "vinspection", "Env", "test")
	Convey("Send Email Via SES Service", t, func() {
		So(err, ShouldBeNil)
	})
}
