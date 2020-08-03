package awswrapper

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSQSService(t *testing.T) {
	mySQS := GetSQSService("us-east-1")
	Convey("Should be able to create a SQS service", t, func() {
		So(mySQS, ShouldNotBeNil)
	})

	myConsumer := NewConsumer(mySQS, handleMessage, 200, "TestSQS", 10, 0, 5, 10, 3)
	Convey("Should be able to create consumers", t, func() {
		So(myConsumer, ShouldNotBeNil)
		So(myConsumer.Consuming, ShouldBeFalse)
	})

	myConsumer.Start()
	Convey("Should be able to start consumer", t, func() {
		time.Sleep(10000) // wait for goroutine to start
		So(myConsumer.Consuming, ShouldBeTrue)
	})

	messageID, err := mySQS.SendMessage("TestSQS", "Hello world 1...!")
	Convey("Should be able to send messages to an existing queue", t, func() {
		So(err, ShouldBeNil)
		So(messageID, ShouldNotBeNil)
	})

	messageID, err = mySQS.SendMessage("NonExistingQueue", "Hello world 2...!")
	Convey("Should not exit when sending messages to non existing queue", t, func() {
		So(err, ShouldNotBeNil)
		So(messageID, ShouldBeNil)
	})

	myConsumer.Stop()
	Convey("Should be able to stop consumer", t, func() {
		So(myConsumer.Consuming, ShouldBeFalse)
	})
}

func handleMessage(msg string) error {
	fmt.Println(msg)
	return nil
}
