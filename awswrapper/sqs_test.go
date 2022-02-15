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
	Convey("Should be able to send a single message to an existing queue", t, func() {
		So(err, ShouldBeNil)
		So(messageID, ShouldNotBeNil)
	})

	failedMessageIDs, successMessageIDs, err := mySQS.SendMessageBatch("TestSQS", []string{"Batch hello world 1...!", "Batch hello world 2...!"})
	Convey("Should be able to send multiple messages to an existing queue", t, func() {
		So(err, ShouldBeNil)
		So(successMessageIDs, ShouldHaveLength, 2)
		So(failedMessageIDs, ShouldHaveLength, 0)
	})

	messageID, err = mySQS.SendMessage("NonExistingQueue", "Hello world 2...!")
	Convey("Should not exit when sending single message to non existing queue", t, func() {
		So(err, ShouldNotBeNil)
		So(messageID, ShouldBeNil)
	})

	failedMessageIDs, successMessageIDs, err = mySQS.SendMessageBatch("NonExistingQueue", []string{"Batch hello world 1...!", "Batch hello world 2...!"})
	Convey("Should not exit when sending multiple messages to non existing queue", t, func() {
		So(err, ShouldNotBeNil)
		So(successMessageIDs, ShouldBeNil)
		So(failedMessageIDs, ShouldBeNil)
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
