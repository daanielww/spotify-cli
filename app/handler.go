package app

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
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

func (h *Handler) HandleRequest (w http.ResponseWriter, r *http.Request){
	res, err := h.Sc.SpotifyCombined()
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	err = h.storeDataInS3(res)
	if err != nil {
		fmt.Errorf("soft error, error storing data in S3: %s", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) storeDataInS3 (res *spotifyclient.RequestResult) (error){
	data, err := json.Marshal(res)
	if err != nil {
		return err
	}

	zippedData, err := zipHelper(data)
	if err != nil {
		return err
	}

	bucket, key := "spotify-cli", "data/" + time.Now().String()

	_, err = h.S3C.PutObject(&s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   bytes.NewReader(zippedData),
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