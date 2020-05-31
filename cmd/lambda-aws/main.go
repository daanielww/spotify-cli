package main

import (
	"fmt"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/daanielww/spotify-cli/api"
	"github.com/daanielww/spotify-cli/cmd/config"
	"log"
	"os"
	"time"
)

type LambdaHandler struct {
	*api.Handler
}

func (lh *LambdaHandler) HandleRequest(ctx context.Context) {

	bucket := "spotify-cli-test"
	_, err := lh.GetAndStorePlaylistsAndAlbums("playlist-album/"+time.Now().Format("2006.01.02 15:04:05"), bucket)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting playlists, albums: ", err)
	}
	_, err = lh.GetAndStoreTracks("tracks/"+time.Now().Format("2006.01.02 15:04:05"), bucket)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting tracks: ", err)
	}
}

func main() {
	c, err := config.ProductionConfig()
	if err != nil {
		log.Fatalln("error with config: ", err.Error())
	}

	lh := &LambdaHandler{
		Handler: &api.Handler{
			c.Sc,
			c.S3C,
		},
	}

	lambda.Start(lh.HandleRequest)
}
