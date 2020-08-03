package awswrapper

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
	"testing"
)

func TestSQS(t *testing.T) {
	mySQS := GetSQSService("us-east-1")
	myConsumer := NewConsumer(mySQS, handleMessage, 200, "TestSQS", 10, 0, 5, 10, 3)
	myConsumer.Start()
	messageID, err := mySQS.SendMessage("TestSQS", "Hello my world...!")
	fmt.Println(err)
	fmt.Println(*messageID)
	myConsumer.Stop()
}

func handleMessage(msg *sqs.Message) error {
	fmt.Println(*msg)
	return nil
}
