package api

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	spotifyclient "github.com/daanielww/spotify-cli/pkg/spotify-client"
	"net/http"
	"time"
)

type Handler struct {
	Sc spotifyclient.SpotifyClient
	S3C s3iface.S3API
}

// Handler for default endpoint - grabbing playlists and albums
func (h *Handler) HandleRequestPlaylistAlbum(w http.ResponseWriter, r *http.Request){
	s3Key := "playlist-album/" + time.Now().Format("2006.01.02 15:04:05")
	bucket := "spotify-cli"
	res, err := h.GetAndStorePlaylistsAndAlbums(s3Key, bucket)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Hanlder for endpoint that grabs tracks
func (h *Handler) HandleRequestTracks(w http.ResponseWriter, r *http.Request){
	s3Key := "tracks/" + time.Now().Format("2006.01.02 15:04:05")
	bucket := "spotify-cli"
	res, err := h.GetAndStoreTracks(s3Key, bucket)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

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