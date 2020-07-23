package awswrapper

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"time"
)

// S3Service represents a S3 service.
type LambdaService struct {
	region  string
	service *lambda.Lambda
}

type HTTPClientSettings struct {
	Connect          time.Duration
	ConnKeepAlive    time.Duration
	ExpectContinue   time.Duration
	IdleConn         time.Duration
	MaxAllIdleConns  int
	MaxHostIdleConns int
	ResponseHeader   time.Duration
	TLSHandshake     time.Duration
}

var (
	lambdaServices = make(map[string]*LambdaService)
)

// GetLambdaService gets a ecr service for a specific region
func GetLambdaService(region string) *LambdaService {
	if lambdaService, ok := lambdaServices[region]; ok {
		return lambdaService
	}

	httpClient, err := NewHTTPClientWithSettings(HTTPClientSettings{
		Connect:          2 * time.Minute,
		ExpectContinue:   1 * time.Second,
		IdleConn:         90 * time.Second,
		ConnKeepAlive:    2 * time.Minute,
		MaxAllIdleConns:  100,
		MaxHostIdleConns: 10,
		ResponseHeader:   2 * time.Minute,
		TLSHandshake:     10 * time.Second,
	})
	if err != nil {
		fmt.Println("Got an error creating custom HTTP client:")
		fmt.Println(err)
		return nil
	}

	svc := lambda.New(sess, &aws.Config{
		Region:     aws.String(region),
		HTTPClient: httpClient,
	})
	lambdaService := &LambdaService{
		region:  region,
		service: svc,
	}
	lambdaServices[region] = lambdaService
	return lambdaService
}

func (o *LambdaService) Invoke(functionName string, payload []byte, isEvent bool) (*lambda.InvokeOutput, error) {
	invocationType := "RequestResponse"
	if isEvent {
		invocationType = "Event"
	}

	return o.service.Invoke(&lambda.InvokeInput{FunctionName: aws.String(functionName), InvocationType: aws.String(invocationType), Payload: payload})
}

func NewHTTPClientWithSettings(httpSettings HTTPClientSettings) (*http.Client, error) {
	var client http.Client
	tr := &http.Transport{
		ResponseHeaderTimeout: httpSettings.ResponseHeader,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: httpSettings.ConnKeepAlive,
			DualStack: true,
			Timeout:   httpSettings.Connect,
		}).DialContext,
		MaxIdleConns:          httpSettings.MaxAllIdleConns,
		IdleConnTimeout:       httpSettings.IdleConn,
		TLSHandshakeTimeout:   httpSettings.TLSHandshake,
		MaxIdleConnsPerHost:   httpSettings.MaxHostIdleConns,
		ExpectContinueTimeout: httpSettings.ExpectContinue,
	}

	// So client makes HTTP/2 requests
	err := http2.ConfigureTransport(tr)
	if err != nil {
		return &client, err
	}

	return &http.Client{
		Transport: tr,
	}, nil
}
