package api

import (
	s3Client "github.com/daanielww/spotify-cli/pkg/s3"
	"github.com/aws/aws-sdk-go/service/s3"
	spotifyclient "github.com/daanielww/spotify-cli/pkg/spotify-client"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_getAndStorePlaylistAlbums(t *testing.T) {
	mockHandler := &Handler{
		Sc:  &spotifyclient.MockSpotify{},
		S3C: s3Client.NewMockS3(),
	}

	key := "data/" + time.Now().Format("2006.01.02 15:04:05")
	res, err := mockHandler.GetAndStorePlaylistsAndAlbums(key)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	getObjectInput := &s3.GetObjectInput{
		Key: &key,
	}
	_, err = mockHandler.S3C.GetObject(getObjectInput)
	assert.NoError(t, err)
}

func Test_getAndStoreTracks(t *testing.T) {
	mockHandler := &Handler{
		Sc:  &spotifyclient.MockSpotify{},
		S3C: s3Client.NewMockS3(),
	}

	key := "tracks/" + time.Now().Format("2006.01.02 15:04:05")
	res, err := mockHandler.GetAndStoreTracks(key)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	getObjectInput := &s3.GetObjectInput{
		Key: &key,
	}
	_, err = mockHandler.S3C.GetObject(getObjectInput)
	assert.NoError(t, err)
}
