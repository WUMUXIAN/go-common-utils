package awswrapper

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SQSService defines the attributes of a SQS service
type SQSService struct {
	region  string
	service *sqs.SQS
}

var (
	sqsServices = make(map[string]*SQSService)
)

// GetSQSService gets a sqs service for a specific region
func GetSQSService(region string) *SQSService {
	if sqsService, ok := sqsServices[region]; ok {
		return sqsService
	}

	svc := sqs.New(sess, &aws.Config{
		Region: aws.String(region),
	})
	sqsService := &SQSService{
		region:  region,
		service: svc,
	}
	sqsServices[region] = sqsService
	return sqsService
}

// SendMessage sends a message payload to a named queue
func (o *SQSService) SendMessage(queueName, payload string) (messageID *string, err error) {
	var queueURL *sqs.GetQueueUrlOutput
	queueURL, err = o.service.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		return
	}
	qURL := queueURL.QueueUrl
	var result *sqs.SendMessageOutput
	result, err = o.service.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(0),
		MessageBody:  aws.String(payload),
		QueueUrl:     qURL,
	})
	if err != nil {
		return
	}
	messageID = result.MessageId

	return
}

var (
	defaultBatchSize = 10
)

// MessageHandler defines an interface for message handler function
type MessageHandler func(*sqs.Message) error

// Consumer defines the attributes of a consumer
type Consumer struct {
	SqsService                    SQSService
	Handler                       MessageHandler
	Consuming                     bool
	PollTimeout                   time.Duration
	StopChan                      chan struct{} // a channel to signal to stop consumer
	StoppedChan                   chan struct{} // a channel to signal that consumer is stopped
	QName                         string
	MaxNumberOfMessagesPerRequest int64
	WaitTimeSeconds               int64
	VisibilityTimeOut             int64 // in seconds
	RetryCount                    int8
	RetryDelay                    time.Duration // in seconds
}

// NewConsumer creates a newConsumer object
func NewConsumer(mySQS *SQSService, messageHandler MessageHandler, pollTimeOut time.Duration, qName string, maxNumberOfMessagesPerRequest, waitTimeSeconds, visibilityTimeoutSeconds int64, retryDelay time.Duration, retryCount int8) *Consumer {
	return &Consumer{
		SqsService:                    *mySQS,
		Handler:                       messageHandler,
		Consuming:                     false,
		PollTimeout:                   pollTimeOut,
		StopChan:                      make(chan struct{}),
		QName:                         qName,
		StoppedChan:                   make(chan struct{}),
		WaitTimeSeconds:               waitTimeSeconds,
		MaxNumberOfMessagesPerRequest: maxNumberOfMessagesPerRequest,
		VisibilityTimeOut:             visibilityTimeoutSeconds,
		RetryCount:                    retryCount,
		RetryDelay:                    retryDelay, // in seconds
	}
}

// Start starts a consumer to consume message in background
func (c *Consumer) Start() {
	if c.Consuming {
		return
	}

	go func() {
		// close the StoppedChan when this func exits
		defer close(c.StoppedChan)
		c.Consuming = true
		consumerTimer := time.NewTimer(c.PollTimeout)
		for {
			select {
			case <-consumerTimer.C:
				for i := 0; i < defaultBatchSize; i++ {
					c.consume()
				}
				consumerTimer.Reset(c.PollTimeout)
			case <-c.StopChan:
				fmt.Println("stopped channel")
				return
			}
		}
	}()
}

// Stop stops a consumer from consuming messages once it completes existing payload
func (c *Consumer) Stop() {
	log.Println("Stopping...")
	close(c.StopChan) // tell it to stop
	<-c.StoppedChan   // wait for it to have stopped
	log.Println("Stopped...")
}

// consume is a helper function for the consumer to consume messages
func (c *Consumer) consume() {
	// Need to convert the queue name into a URL. Make the GetQueueUrl
	// API call to retrieve the URL. This is needed for receiving messages
	// from the queue.
	resultURL, err := c.SqsService.service.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(c.QName),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == sqs.ErrCodeQueueDoesNotExist {
			exitErrorf("Unable to find queue %q.", c.QName)
		}
		exitErrorf("Unable to queue %q, %v.", c.QName, err)
	}

	// Receive a message from the SQS queue with long polling enabled.
	result, err := c.SqsService.service.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            resultURL.QueueUrl,
		AttributeNames:      aws.StringSlice([]string{"SentTimestamp"}),
		MaxNumberOfMessages: aws.Int64(c.MaxNumberOfMessagesPerRequest),
		WaitTimeSeconds:     aws.Int64(c.WaitTimeSeconds),
		VisibilityTimeout:   aws.Int64(c.VisibilityTimeOut),
	})
	if err != nil {
		exitErrorf("Unable to receive message from queue %q, %v.", c.QName, err)
	}

	fmt.Printf("Received %d messages.\n", len(result.Messages))
	for i, msg := range result.Messages {
		for j := int8(1); j < c.RetryCount+1; j++ {
			err = c.Handler(msg)
			if err == nil {
				c.deleteMessage(resultURL.QueueUrl, result.Messages[i])
			} else if j == c.RetryCount {
				fmt.Println("unable to process message. maximum retry failed.")
				_, err = c.SqsService.SendMessage("DeadLetterQueue", result.Messages[i].String())
				if err != nil {
					fmt.Println("error while sending message to DLQ: ", err)
				}
				c.deleteMessage(resultURL.QueueUrl, result.Messages[i])
			} else {
				fmt.Println("consuming message failed: ", err)
				time.Sleep(c.RetryDelay * time.Second)
			}
		}
	}
}

// deleteMessage is a helper function for the consumer to delete messages
func (c *Consumer) deleteMessage(qURL *string, msg *sqs.Message) {
	_, err := c.SqsService.service.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      qURL,
		ReceiptHandle: msg.ReceiptHandle,
	})
	if err != nil {
		fmt.Println("Delete Error", err)
		return
	}
	fmt.Println("Message Deleted")
}

func exitErrorf(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
