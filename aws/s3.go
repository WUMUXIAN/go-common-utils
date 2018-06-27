package aws

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3Service represents a S3 service.
type S3Service struct {
	region  string
	service *s3.S3
}

var (
	s3Services = make(map[string]*S3Service)
)

// GetS3Service gets a s3 service for a specific region
func GetS3Service(region string) *S3Service {
	if s3Service, ok := s3Services[region]; ok {
		return s3Service
	}
	svc := s3.New(sess, &aws.Config{
		Region: aws.String(region),
	})
	s3Service := &S3Service{
		region:  region,
		service: svc,
	}
	s3Services[region] = s3Service
	return s3Service
}

// CreateBucket creates a bucket with given name and in given region
func (o *S3Service) CreateBucket(bucketName, region string) error {
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}
	if region != "us-east-1" {
		input.CreateBucketConfiguration = &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(region),
		}
	}
	result, err := o.service.CreateBucket(input)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	return err
}

// DeleteBucket delete a bucket with given name and in given region
func (o *S3Service) DeleteBucket(bucketName string) error {
	// Before we can delete we need to remove everything.
	o.RemoveAllFromS3(bucketName, "/")

	result, err := o.service.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
	return err
}

// UploadToS3 uploads the content in byte to S3 with a specific bucket and path
// args. [0]. Make Public Or Not.
func (o *S3Service) UploadToS3(content []byte, bucketName string, path string, args ...interface{}) error {
	bodyBytes := bytes.NewReader(content)
	acl := s3.ObjectCannedACLPrivate
	if len(args) > 0 {
		for i, arg := range args {
			switch i {
			case 0:
				if arg.(bool) {
					acl = s3.ObjectCannedACLPublicRead
				}
			}
		}
	}
	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(path),
		Body:          bodyBytes,
		ContentLength: aws.Int64(int64(len(content))),
		ContentType:   aws.String(http.DetectContentType(content)),
		ACL:           aws.String(acl),
	}
	fmt.Println("Uploading Object: ", path)
	resp, err := o.service.PutObject(params)
	if err != nil {
		fmt.Println("bad response: %s", err)
		return err
	}
	fmt.Printf("response %s\n", awsutil.StringValue(resp))
	return nil
}

// ReadFromS3 read the content in byte from S3 with a specific bucket and path
func (o *S3Service) ReadFromS3(bucketName string, path string) (content []byte, err error) {
	resp, err := o.service.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	})

	fmt.Println("Downloading Object: ", path)

	if err != nil {
		fmt.Printf("bad response: %s\n", err)
		return
	}

	content, err = ioutil.ReadAll(resp.Body)
	return
}

func (o *S3Service) listS3Objects(bucketName string, path string, marker *string) (objectPaths []string, morePages bool, err error) {
	resp, err := o.service.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(path),
		Marker: marker,
	})

	fmt.Println("Listing All Objects Next Page: ", path)

	objectPaths = make([]string, 0)
	morePages = false

	if err != nil {
		fmt.Printf("bad response: %s\n", err)
		return
	}

	for _, content := range resp.Contents {
		objectPaths = append(objectPaths, aws.StringValue(content.Key))
	}

	morePages = aws.BoolValue(resp.IsTruncated)

	return
}

// ListAllS3 lists all objects with the path as prefix
func (o *S3Service) ListAllS3(bucketName string, path string) (objectPaths []string, err error) {
	objectPaths = make([]string, 0)

	pagedObjects, morePages, err := o.listS3Objects(bucketName, path, nil)

	if err != nil {
		return
	}

	objectPaths = append(objectPaths, pagedObjects...)
	for {
		if morePages {
			pagedObjects, morePages, _ = o.listS3Objects(bucketName, path, aws.String(objectPaths[len(objectPaths)-1]))
			objectPaths = append(objectPaths, pagedObjects...)
		} else {
			break
		}
	}
	return
}

// CopyWithInS3 copys a object from one place to another within a bucket on S3
func (o *S3Service) CopyWithInS3(sourceBucketName, sourcePath, destBucketName, destPath string, deleteAfterCopy bool) (err error) {
	if string(sourcePath[0]) != "/" {
		sourcePath = "/" + sourcePath
	}

	if string(destPath[0]) != "/" {
		destPath = "/" + destPath
	}

	_, err = o.service.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(destBucketName),
		CopySource: aws.String(sourceBucketName + sourcePath),
		Key:        aws.String(destPath),
	})

	fmt.Println("Copying Object From:", sourceBucketName+sourcePath, "to", destBucketName+destPath)

	if err != nil {
		fmt.Printf("bad response: %s\n", err)
		return
	}

	if deleteAfterCopy {
		return o.RemoveFromS3(sourceBucketName, sourcePath)
	}

	return
}

// RemoveFromS3 removes a object using its bucketname and path
func (o *S3Service) RemoveFromS3(bucketName string, path string) (err error) {
	resp, err := o.service.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	})

	fmt.Println("Removing Object: ", path)

	if err != nil {
		fmt.Printf("bad response: %s\n", err)
		return
	}

	fmt.Printf("response %s\n", awsutil.StringValue(resp))
	return
}

// RemoveAllFromS3 removes a folder with bucketname and path
func (o *S3Service) RemoveAllFromS3(bucketName string, path string) (err error) {
	objectPaths, err := o.ListAllS3(bucketName, path)

	if err != nil {
		fmt.Printf("bad response: %s\n", err)
		return
	}

	if len(objectPaths) == 0 {
		return
	}

	objects := make([]*s3.ObjectIdentifier, len(objectPaths))
	for i, ob := range objectPaths {
		object := &s3.ObjectIdentifier{
			Key: aws.String(ob),
		}
		objects[i] = object
	}

	fmt.Println("Objects: ", objects)

	d := &s3.Delete{
		Objects: objects,
	}
	_, err = o.service.DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: d,
	})

	if err != nil {
		fmt.Printf("bad response: %s\n", err)
		return
	}

	fmt.Println("Deleting All Objects in: ", path)

	return
}

// UploadToS3Concurrently uploads content to S3 concurrently
// args. [0]. Make Public Or Not.
func (o *S3Service) UploadToS3Concurrently(content []byte, bucketName string, path string, args ...interface{}) error {

	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(o.region),
	}))

	uploader := s3manager.NewUploader(session)
	bodyBytes := bytes.NewReader(content)
	acl := s3.ObjectCannedACLPrivate
	if len(args) > 0 {
		for i, arg := range args {
			switch i {
			case 0:
				if arg.(bool) {
					fmt.Println("Make it Public")
					acl = s3.ObjectCannedACLPublicRead
				}
			}
		}
	}
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(path),
		Body:        bodyBytes,
		ContentType: aws.String(http.DetectContentType(content)),
		ACL:         aws.String(acl),
	})
	fmt.Println("Uploading Object: ", path)
	if err != nil {
		if multierr, ok := err.(s3manager.MultiUploadFailure); ok {
			// Process error and its associated uploadID
			fmt.Println("Error:", multierr.Code(), multierr.Message(), multierr.UploadID())
		} else {
			// Process error generically
			fmt.Println("Error:", err.Error())
		}
	} else {
		fmt.Println("Upload Output: ", output)
	}

	return err
}

// ReadFromS3Concurrently reads content from S3 concurrently
func (o *S3Service) ReadFromS3Concurrently(bucketName string, path string) (content []byte, err error) {

	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(o.region),
	}))

	downloader := s3manager.NewDownloader(session)

	var buffer aws.WriteAtBuffer

	_, err = downloader.Download(&buffer, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	})

	fmt.Println("Downloading Object: ", path)

	if err != nil {
		fmt.Printf("bad response: %s\n", err)
	} else {
		content = buffer.Bytes()
	}
	return
}

// InitMultiPartUpload inits a multiple parts upload
func (o *S3Service) InitMultiPartUpload(bucketName, path string) (string, error) {
	input := &s3.CreateMultipartUploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	}
	result, err := o.service.CreateMultipartUpload(input)
	fmt.Println("Starting Multi Part Upload At Path: ", path)
	if err != nil {
		return "", err
	}
	return *result.UploadId, err
}

// UploadMultipart uploads a part to S3
func (o *S3Service) UploadMultipart(bucketName, path, uploadID string, partNumber int64, b []byte) (string, error) {
	input := &s3.UploadPartInput{
		Body:       aws.ReadSeekCloser(bytes.NewReader(b)),
		Bucket:     aws.String(bucketName),
		Key:        aws.String(path),
		PartNumber: aws.Int64(partNumber),
		UploadId:   aws.String(uploadID),
	}

	result, err := o.service.UploadPart(input)
	fmt.Println("Uploaded Multi Part: Path->", path, "PartNumber->", partNumber)
	if err != nil {
		return "", err
	}
	return *result.ETag, err
}

// CompleteMultipart marks a multi part upload as complete
func (o *S3Service) CompleteMultipart(bucketName, path, uploadID string, parts map[string]interface{}) error {
	completedParts := make([]*s3.CompletedPart, len(parts))
	for i, p := range parts {
		n, _ := strconv.ParseInt(i, 10, 64)
		completedParts[n-1] = &s3.CompletedPart{
			ETag:       aws.String(p.(string)),
			PartNumber: aws.Int64(n),
		}
	}
	input := &s3.CompleteMultipartUploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completedParts,
		},
		UploadId: aws.String(uploadID),
	}
	_, err := o.service.CompleteMultipartUpload(input)
	if err == nil {
		fmt.Println("Completed Multi Part For Path:", path)
	}
	return err
}

// AbortMultipart aborts a multipart upload
func (o *S3Service) AbortMultipart(bucketName, path, uploadID string) error {
	input := &s3.AbortMultipartUploadInput{
		Bucket:   aws.String(bucketName),
		Key:      aws.String(path),
		UploadId: aws.String(uploadID),
	}
	_, err := o.service.AbortMultipartUpload(input)
	if err == nil {
		fmt.Println("Aborted Multi Part For Path:", path)
	}
	return err
}

// GetPreSignedURL gets pre-signed URL that are valid for specified duration.
func (o *S3Service) GetPreSignedURL(bucketName, path string, validFor time.Duration) (string, error) {
	req, _ := o.service.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	})
	urlStr, err := req.Presign(validFor)

	if err != nil {
		log.Println("Failed to sign request", err)
		return "", err
	}

	return urlStr, nil
}
