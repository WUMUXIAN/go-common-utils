// Package awswrapper wraps over the AWS official GO SDK to simplify the usage.
package awswrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	sess *session.Session
)

func init() {
	// Set the us-east-1 as the default region.
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsWest2RegionID),
	}))
}
