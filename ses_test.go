package aws

import "testing"

func TestSendSESEmail(t *testing.T) {
	err := GetSESService("us-east-1").SendSESEmail("test", "Tectus DreamLab", "no-reply@tectusdreamlab.com", "wumuxian1988@gmail.com", "sample data", "html", "ServiceName", "procep", "Env", "test")
	if err != nil {
		t.Fatal(err)
	}
}
