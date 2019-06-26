package awswrapper

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"time"
)

// S3Service represents a S3 service.
type ECSService struct {
	region  string
	service *ecs.ECS
}

var (
	ecsServices = make(map[string]*ECSService)
)

// GetECSService gets a ecr service for a specific region
func GetECSService(region string) *ECSService {
	if ecsService, ok := ecsServices[region]; ok {
		return ecsService
	}
	svc := ecs.New(sess, &aws.Config{
		Region: aws.String(region),
	})
	ecsService := &ECSService{
		region:  region,
		service: svc,
	}
	ecsServices[region] = ecsService
	return ecsService
}

// RunTask starts a new task using the specified task definition and in given region
func (o *ECSService) RunTask(cluster, container, taskDefinitionName string, envVariables map[string]string, commands []string, fargate bool, subnetIDs, securityGroupIDs []string, timeout time.Duration) (err error) {
	override := &ecs.TaskOverride{}
	if container != "" || len(envVariables) != 0 || len(commands) != 0 {
		containerOverride := &ecs.ContainerOverride{}
		if container != "" {
			containerOverride.Name = aws.String(container)
		}
		if len(envVariables) != 0 {
			keyValues := make([]*ecs.KeyValuePair, 0)
			for name, value := range envVariables {
				keyValues = append(keyValues, &ecs.KeyValuePair{Name: aws.String(name), Value: aws.String(value)})
			}
			containerOverride.Environment = keyValues
		}
		if len(commands) != 0 {
			containerOverride.Command = aws.StringSlice(commands)
		}
		override.SetContainerOverrides([]*ecs.ContainerOverride{containerOverride})
	}

	var params *ecs.RunTaskInput

	launchType := "EC2"
	assignPublicIP := "DISABLED"
	if fargate {
		launchType = "FARGATE"
		assignPublicIP = "ENABLED"
	}

	if len(subnetIDs) > 0 {
		vpcConfiguration := &ecs.AwsVpcConfiguration{
			AssignPublicIp: aws.String(assignPublicIP),
			Subnets:        aws.StringSlice(subnetIDs),
			SecurityGroups: aws.StringSlice(securityGroupIDs),
		}
		network := &ecs.NetworkConfiguration{
			AwsvpcConfiguration: vpcConfiguration,
		}
		params = &ecs.RunTaskInput{
			Cluster:              aws.String(cluster),
			TaskDefinition:       aws.String(taskDefinitionName),
			Overrides:            override,
			NetworkConfiguration: network,
			LaunchType:           aws.String(launchType),
		}
	} else {
		params = &ecs.RunTaskInput{
			Cluster:        aws.String(cluster),
			TaskDefinition: aws.String(taskDefinitionName),
			Overrides:      override,
			LaunchType:     aws.String(launchType),
		}
	}

	out, err := o.service.RunTask(params)
	fmt.Println(out)
	return err
}
