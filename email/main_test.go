package email

import (
	"errors"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEmail(t *testing.T) {
	Convey("Valid Format Function\n", t, func() {
		Convey("abc.com Should Not Be An Email", func() {
			So(ValidateFormat("abc.com"), ShouldBeError, errors.New("invalid format"))
		})
		Convey("123@@@@.com Should Not Be An Email", func() {
			So(ValidateFormat("123@@@@.com"), ShouldBeError, errors.New("invalid format"))
		})
		Convey("123@.com.com Should Not Be An Email", func() {
			So(ValidateFormat("123@.com.com"), ShouldBeError, errors.New("invalid format"))
		})
		Convey("()^*^%^@gmail.com Should Not Be An Email", func() {
			So(ValidateFormat("()^*^%^@gmail.com"), ShouldBeError, errors.New("invalid format"))
		})
		Convey("abc@com Should Be An Email", func() {
			So(ValidateFormat("abc@com"), ShouldBeNil)
		})
		Convey("email@xxx.com Should Be An Email", func() {
			So(ValidateFormat("email@xxx.com"), ShouldBeNil)
		})
	})

	Convey("Valid Reachability Function\n", t, func() {
		Convey("Test Reachability For Valid Email", func() {
			So(ValidateReachability(os.Getenv("GMAIL_USERNAME"), time.Second), ShouldBeNil)
		})
		Convey("fakeemail@somewebsite.com Should Not Be Reachable", func() {
			So(ValidateReachability("fakeemail@somewebsite.com", time.Second*5), ShouldBeError, errors.New("unresolvable host"))
		})
	})

	Convey("Send Out An Email Via SMTP\n", t, func() {
		err := SendSMTP(os.Getenv("GMAIL_USERNAME"), "This is a test email, do not reply", "Test Content",
			os.Getenv("GMAIL_USERNAME"), os.Getenv("GMAIL_PASSWORD"), "smtp.gmail.com", "Test Bot", "no-reply@gmail.com")
		So(err, ShouldBeNil)
	})
}
