package awswrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"testing"
)
import . "github.com/smartystreets/goconvey/convey"

func TestSendSESEmail(t *testing.T) {
	err := GetSESService("us-east-1").SendSESEmail("test", "Tectus DreamLab", "no-reply@tectusdreamlab.com", "wumuxian1988@gmail.com", nil, "sample data", "html", "ServiceName", "vinspection", "Env", "test")
	Convey("Send Email Via SES Service", t, func() {
		So(err, ShouldBeNil)
	})

	err = GetSESService("us-east-1").SendSESEmail("test", "Tectus DreamLab", "no-reply@tectusdreamlab.com", "wumuxian1988@gmail.com", aws.String("shebin.vincent@screeningeagle.com"), "sample data", "text", "ServiceName", "vinspection", "Env", "test")
	Convey("Send Email Via SES Service", t, func() {
		So(err, ShouldBeNil)
	})
}
