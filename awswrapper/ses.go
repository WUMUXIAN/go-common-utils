package awswrapper

import (
	"net/mail"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

// SESService represents a SES service.
type SESService struct {
	region  string
	service *ses.SES
}

var (
	sesServices = make(map[string]*SESService)
)

// GetSESService gets a SES service for a specific region
func GetSESService(region string) *SESService {
	if sesService, ok := sesServices[region]; ok {
		return sesService
	}
	svc := ses.New(sess, &aws.Config{
		Region: aws.String(region),
	})
	sesService := &SESService{
		region:  region,
		service: svc,
	}
	sesServices[region] = sesService
	return sesService
}

// SendSESEmail sends email using AWS SES api
// Parameters:
// subject: subject of the email
// fromAddr: email address it's from
// fromName: whom it's from
// toAddr: email address to send to
// ccAddr: email address to send cc
// data: the body of the email
// dataType: the type of the body, text or html
// tags: tags on the email, comes in as key value pair, e.g. "serviceName", "proceq", "env", "dev"
func (o *SESService) SendSESEmail(subject, fromName, fromAddr, toAddr string, ccAddr *string, data, dataType string, tags ...string) error {
	from := mail.Address{
		Name:    fromName,
		Address: fromAddr,
	}

	var body *ses.Body
	switch dataType {
	case "text":
		body = &ses.Body{ // Required
			Text: &ses.Content{
				Data:    aws.String(data), // Required
				Charset: aws.String("utf-8"),
			},
		}
	case "html":
		body = &ses.Body{ // Required
			Html: &ses.Content{
				Data:    aws.String(data), // Required
				Charset: aws.String("utf-8"),
			},
		}

	}

	destination := &ses.Destination{ // Required
		ToAddresses: []*string{
			aws.String(toAddr), // Required
		},
	}
	if ccAddr != nil {
		destination.CcAddresses = []*string{ccAddr}
	}

	params := &ses.SendEmailInput{
		Destination: destination,
		Message: &ses.Message{ // Required
			Body: body,
			Subject: &ses.Content{ // Required
				Data:    aws.String(subject), // Required
				Charset: aws.String("utf-8"),
			},
		},
		Source:               aws.String(from.String()), // Required
		ConfigurationSetName: aws.String("shared"),
	}

	if len(tags) > 0 && len(tags)%2 == 0 {
		messageTags := make([]*ses.MessageTag, 0)
		for i := 0; i < len(tags)-1; i += 2 {
			name := tags[i]
			value := tags[i+1]
			messageTags = append(messageTags, &ses.MessageTag{
				Name:  aws.String(name),
				Value: aws.String(value),
			})
		}
		params.Tags = messageTags
	}
	_, err := o.service.SendEmail(params)
	return err
}
