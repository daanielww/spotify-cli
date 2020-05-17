package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"os"
)

func getAWSRegion() *string {
	aws_region := os.Getenv("AWS_REGION")
	if aws_region != "" {
		return &aws_region
	}

	aws_default_region := os.Getenv("AWS_DEFAULT_REGION")

	if aws_default_region != "" {
		return &aws_default_region
	}

	region := "us-east-1"
	return &region
}

func GetS3(stub bool) (s3iface.S3API, error) {
	if stub {
		return NewMockS3(), nil
	}

	sess, err := session.NewSession(&aws.Config{
		Region: getAWSRegion(),
	})

	if err != nil {
		return nil, err
	}

	S3C := s3.New(sess)
	return S3C, nil
}
