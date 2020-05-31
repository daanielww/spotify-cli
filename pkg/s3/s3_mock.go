package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"io/ioutil"
)

// Mock AWS S3 client used for testing and for running locally/development

type MockS3 struct {
	s3iface.S3API
	Objects map[string][]byte
}

func NewMockS3() *MockS3 {
	return &MockS3{
		Objects:      map[string][]byte{},
	}
}

func (m *MockS3) PutObject(in *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	key := *in.Key

	byt, err := ioutil.ReadAll(in.Body)

	if err != nil {
		return nil, err
	}

	m.Objects[key] = byt
	return &s3.PutObjectOutput{}, nil
}

func (m *MockS3) GetObject(in *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	key := *in.Key

	_, ok := m.Objects[key]
	if !ok {
		return nil, fmt.Errorf("object doesn't exist")
	}

	return &s3.GetObjectOutput{}, nil
}



