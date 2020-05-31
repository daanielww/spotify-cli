package api

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	spotifyclient "github.com/daanielww/spotify-cli/pkg/spotify-client"
)

type Handler struct {
	Sc spotifyclient.SpotifyClient
	S3C s3iface.S3API
}

// API logic

// Retrieves tracks and stores in S3
func (h *Handler) GetAndStoreTracks(s3Key, s3Bucket string) (*spotifyclient.TrackRequestResult, error) {
	res, err := h.Sc.GetTracks()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	err = h.storeDataInS3(s3Key, s3Bucket, data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Retrieves playlists and albums and store them in S3
func (h *Handler) GetAndStorePlaylistsAndAlbums(s3Key, s3Bucket string) (*spotifyclient.AlbumsPlaylistRequestResult, error) {
	res, err := h.Sc.SpotifyCombinedPlaylistAlbum()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	err = h.storeDataInS3(s3Key, s3Bucket, data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Function to store data in AWS S3
func (h *Handler) storeDataInS3 (s3Key, s3Bucket string, data []byte) (error){

	_, err := h.S3C.PutObject(&s3.PutObjectInput{
		Bucket: &s3Bucket,
		Key:    &s3Key,
		Body:   bytes.NewReader(data),
	})

	if err != nil {
		return err
	}

	return nil
}

func zipHelper(in []byte) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(in); err != nil {
		return nil, err
	}
	if err := gz.Flush(); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}