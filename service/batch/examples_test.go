// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/batch"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To cancel a job
//
// This example cancels a job with the specified job ID.
func ExampleBatch_CancelJob_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.CancelJobInput{
		JobId:  aws.String("1d828f65-7a4d-42e8-996d-3b900ed59dc4"),
		Reason: aws.String("Cancelling job."),
	}

	result, err := svc.CancelJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a managed EC2 compute environment
//
// This example creates a managed compute environment with specific C4 instance types
// that are launched on demand. The compute environment is called C4OnDemand.
func ExampleBatch_CreateComputeEnvironment_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.CreateComputeEnvironmentInput{
		ComputeEnvironmentName: aws.String("C4OnDemand"),
		ComputeResources: &batch.ComputeResource{
			DesiredvCpus: aws.Int64(48),
			Ec2KeyPair:   aws.String("id_rsa"),
			InstanceRole: aws.String("ecsInstanceRole"),
			InstanceTypes: []*string{
				aws.String("c4.large"),
				aws.String("c4.xlarge"),
				aws.String("c4.2xlarge"),
				aws.String("c4.4xlarge"),
				aws.String("c4.8xlarge"),
			},
			MaxvCpus: aws.Int64(128),
			MinvCpus: aws.Int64(0),
			SecurityGroupIds: []*string{
				aws.String("sg-cf5093b2"),
			},
			Subnets: []*string{
				aws.String("subnet-220c0e0a"),
				aws.String("subnet-1a95556d"),
				aws.String("subnet-978f6dce"),
			},
			Tags: map[string]*string{
				"Name": aws.String("Batch Instance - C4OnDemand"),
			},
			Type: aws.String("EC2"),
		},
		ServiceRole: aws.String("arn:aws:iam::012345678910:role/AWSBatchServiceRole"),
		State:       aws.String("ENABLED"),
		Type:        aws.String("MANAGED"),
	}

	result, err := svc.CreateComputeEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a managed EC2 Spot compute environment
//
// This example creates a managed compute environment with the M4 instance type that
// is launched when the Spot bid price is at or below 20% of the On-Demand price for
// the instance type. The compute environment is called M4Spot.
func ExampleBatch_CreateComputeEnvironment_shared01() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.CreateComputeEnvironmentInput{
		ComputeEnvironmentName: aws.String("M4Spot"),
		ComputeResources: &batch.ComputeResource{
			BidPercentage: aws.Int64(20),
			DesiredvCpus:  aws.Int64(4),
			Ec2KeyPair:    aws.String("id_rsa"),
			InstanceRole:  aws.String("ecsInstanceRole"),
			InstanceTypes: []*string{
				aws.String("m4"),
			},
			MaxvCpus: aws.Int64(128),
			MinvCpus: aws.Int64(0),
			SecurityGroupIds: []*string{
				aws.String("sg-cf5093b2"),
			},
			SpotIamFleetRole: aws.String("arn:aws:iam::012345678910:role/aws-ec2-spot-fleet-role"),
			Subnets: []*string{
				aws.String("subnet-220c0e0a"),
				aws.String("subnet-1a95556d"),
				aws.String("subnet-978f6dce"),
			},
			Tags: map[string]*string{
				"Name": aws.String("Batch Instance - M4Spot"),
			},
			Type: aws.String("SPOT"),
		},
		ServiceRole: aws.String("arn:aws:iam::012345678910:role/AWSBatchServiceRole"),
		State:       aws.String("ENABLED"),
		Type:        aws.String("MANAGED"),
	}

	result, err := svc.CreateComputeEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a job queue with a single compute environment
//
// This example creates a job queue called LowPriority that uses the M4Spot compute
// environment.
func ExampleBatch_CreateJobQueue_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.CreateJobQueueInput{
		ComputeEnvironmentOrder: []*batch.ComputeEnvironmentOrder{
			{
				ComputeEnvironment: aws.String("M4Spot"),
				Order:              aws.Int64(1),
			},
		},
		JobQueueName: aws.String("LowPriority"),
		Priority:     aws.Int64(1),
		State:        aws.String("ENABLED"),
	}

	result, err := svc.CreateJobQueue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a job queue with multiple compute environments
//
// This example creates a job queue called HighPriority that uses the C4OnDemand compute
// environment with an order of 1 and the M4Spot compute environment with an order of
// 2.
func ExampleBatch_CreateJobQueue_shared01() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.CreateJobQueueInput{
		ComputeEnvironmentOrder: []*batch.ComputeEnvironmentOrder{
			{
				ComputeEnvironment: aws.String("C4OnDemand"),
				Order:              aws.Int64(1),
			},
			{
				ComputeEnvironment: aws.String("M4Spot"),
				Order:              aws.Int64(2),
			},
		},
		JobQueueName: aws.String("HighPriority"),
		Priority:     aws.Int64(10),
		State:        aws.String("ENABLED"),
	}

	result, err := svc.CreateJobQueue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a compute environment
//
// This example deletes the P2OnDemand compute environment.
func ExampleBatch_DeleteComputeEnvironment_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.DeleteComputeEnvironmentInput{
		ComputeEnvironment: aws.String("P2OnDemand"),
	}

	result, err := svc.DeleteComputeEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a job queue
//
// This example deletes the GPGPU job queue.
func ExampleBatch_DeleteJobQueue_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.DeleteJobQueueInput{
		JobQueue: aws.String("GPGPU"),
	}

	result, err := svc.DeleteJobQueue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To deregister a job definition
//
// This example deregisters a job definition called sleep10.
func ExampleBatch_DeregisterJobDefinition_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.DeregisterJobDefinitionInput{
		JobDefinition: aws.String("sleep10"),
	}

	result, err := svc.DeregisterJobDefinition(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe a compute environment
//
// This example describes the P2OnDemand compute environment.
func ExampleBatch_DescribeComputeEnvironments_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.DescribeComputeEnvironmentsInput{
		ComputeEnvironments: []*string{
			aws.String("P2OnDemand"),
		},
	}

	result, err := svc.DescribeComputeEnvironments(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe active job definitions
//
// This example describes all of your active job definitions.
func ExampleBatch_DescribeJobDefinitions_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.DescribeJobDefinitionsInput{
		Status: aws.String("ACTIVE"),
	}

	result, err := svc.DescribeJobDefinitions(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe a job queue
//
// This example describes the HighPriority job queue.
func ExampleBatch_DescribeJobQueues_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.DescribeJobQueuesInput{
		JobQueues: []*string{
			aws.String("HighPriority"),
		},
	}

	result, err := svc.DescribeJobQueues(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe a specific job
//
// This example describes a job with the specified job ID.
func ExampleBatch_DescribeJobs_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.DescribeJobsInput{
		Jobs: []*string{
			aws.String("24fa2d7a-64c4-49d2-8b47-f8da4fbde8e9"),
		},
	}

	result, err := svc.DescribeJobs(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list running jobs
//
// This example lists the running jobs in the HighPriority job queue.
func ExampleBatch_ListJobs_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.ListJobsInput{
		JobQueue: aws.String("HighPriority"),
	}

	result, err := svc.ListJobs(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list submitted jobs
//
// This example lists jobs in the HighPriority job queue that are in the SUBMITTED job
// status.
func ExampleBatch_ListJobs_shared01() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.ListJobsInput{
		JobQueue:  aws.String("HighPriority"),
		JobStatus: aws.String("SUBMITTED"),
	}

	result, err := svc.ListJobs(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To register a job definition
//
// This example registers a job definition for a simple container job.
func ExampleBatch_RegisterJobDefinition_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.RegisterJobDefinitionInput{
		ContainerProperties: &batch.ContainerProperties{
			Command: []*string{
				aws.String("sleep"),
				aws.String("10"),
			},
			Image:  aws.String("busybox"),
			Memory: aws.Int64(128),
			Vcpus:  aws.Int64(1),
		},
		JobDefinitionName: aws.String("sleep10"),
		Type:              aws.String("container"),
	}

	result, err := svc.RegisterJobDefinition(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To submit a job to a queue
//
// This example submits a simple container job called example to the HighPriority job
// queue.
func ExampleBatch_SubmitJob_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.SubmitJobInput{
		JobDefinition: aws.String("sleep60"),
		JobName:       aws.String("example"),
		JobQueue:      aws.String("HighPriority"),
	}

	result, err := svc.SubmitJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To terminate a job
//
// This example terminates a job with the specified job ID.
func ExampleBatch_TerminateJob_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.TerminateJobInput{
		JobId:  aws.String("61e743ed-35e4-48da-b2de-5c8333821c84"),
		Reason: aws.String("Terminating job."),
	}

	result, err := svc.TerminateJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update a compute environment
//
// This example disables the P2OnDemand compute environment so it can be deleted.
func ExampleBatch_UpdateComputeEnvironment_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.UpdateComputeEnvironmentInput{
		ComputeEnvironment: aws.String("P2OnDemand"),
		State:              aws.String("DISABLED"),
	}

	result, err := svc.UpdateComputeEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update a job queue
//
// This example disables a job queue so that it can be deleted.
func ExampleBatch_UpdateJobQueue_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := batch.New(cfg)
	input := &batch.UpdateJobQueueInput{
		JobQueue: aws.String("GPGPU"),
		State:    aws.String("DISABLED"),
	}

	result, err := svc.UpdateJobQueue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}