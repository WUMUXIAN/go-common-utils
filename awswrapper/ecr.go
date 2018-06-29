package awswrapper

import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// ECRService represents anecr service.
type ECRService struct {
	region  string
	service *ecr.ECR
}

var (
	ecrServices = make(map[string]*ECRService)
)

// GetECRService gets a ecr service for a specific region
func GetECRService(region string) *ECRService {
	if ecrService, ok := ecrServices[region]; ok {
		return ecrService
	}
	svc := ecr.New(sess, &aws.Config{
		Region: aws.String(region),
	})
	ecrService := &ECRService{
		region:  region,
		service: svc,
	}
	ecrServices[region] = ecrService
	return ecrService
}

// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(r1, r2 *ecr.ImageDetail) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(imageDetails []*ecr.ImageDetail) {
	sorter := &imageDetailSorter{
		imageDetails: imageDetails,
		by:           by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(sorter)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type imageDetailSorter struct {
	imageDetails []*ecr.ImageDetail
	by           func(r1, r2 *ecr.ImageDetail) bool // Closure used in the Less method.
}

func (o *imageDetailSorter) Len() int { return len(o.imageDetails) }
func (o *imageDetailSorter) Swap(i, j int) {
	o.imageDetails[i], o.imageDetails[j] = o.imageDetails[j], o.imageDetails[i]
}
func (o *imageDetailSorter) Less(i, j int) bool {
	return o.by(o.imageDetails[i], o.imageDetails[j])
}

// ListImages returns all tagged images for a given repo
func (o *ECRService) ListImages(repoName string) (tags [][]string) {
	tags = make([][]string, 0)

	result, err := o.service.DescribeImages(&ecr.DescribeImagesInput{
		RepositoryName: aws.String(repoName),
		Filter: &ecr.DescribeImagesFilter{
			TagStatus: aws.String("TAGGED"),
		},
	})
	if err != nil {
		fmt.Println("Error listing images")
		return
	}

	byPushedAtDesc := func(r1, r2 *ecr.ImageDetail) bool {
		return aws.TimeValue(r1.ImagePushedAt).Unix() > aws.TimeValue(r2.ImagePushedAt).Unix()
	}
	By(byPushedAtDesc).Sort(result.ImageDetails)

	for _, result := range result.ImageDetails {
		tags = append(tags, aws.StringValueSlice(result.ImageTags))
	}
	return
}
