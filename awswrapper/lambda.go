package awswrapper

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// S3Service represents a S3 service.
type LambdaService struct {
	region  string
	service *lambda.Lambda
}

var (
	lambdaServices = make(map[string]*LambdaService)
)

// GetLambdaService gets a ecr service for a specific region
func GetLambdaService(region string) *LambdaService {
	if lambdaService, ok := lambdaServices[region]; ok {
		return lambdaService
	}
	svc := lambda.New(sess, &aws.Config{
		Region: aws.String(region),
	})
	lambdaService := &LambdaService{
		region:  region,
		service: svc,
	}
	lambdaServices[region] = lambdaService
	return lambdaService
}

func (o *LambdaService) Invoke(functionName string, payload []byte, isEvent bool) (err error) {
	invocationType := "RequestResponse"
	if isEvent {
		invocationType = "Event"
	}
	out, err := o.service.Invoke(&lambda.InvokeInput{FunctionName: aws.String(functionName), InvocationType: aws.String(invocationType), Payload: payload})
	fmt.Println(out)
	return err
}
