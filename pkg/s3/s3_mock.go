package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"io/ioutil"
)

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
	fmt.Println("Putting Mock Object ", key)

	byt, err := ioutil.ReadAll(in.Body)

	if err != nil {
		return nil, err
	}

	m.Objects[key] = byt
	return &s3.PutObjectOutput{}, nil
}



