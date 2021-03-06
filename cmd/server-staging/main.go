package main

import (
	"fmt"
	"github.com/daanielww/spotify-cli/api"
	"github.com/daanielww/spotify-cli/cmd/config"
	"log"
	"net/http"
)


func main() {
	c, err := config.ProductionConfig()
	if err != nil {
		log.Fatalln("error with config: ", err.Error())
	}

	h := api.Handler{c.Sc, c.S3C}

	fmt.Println("starting staging server")
	http.HandleFunc("/", h.HandleRequestPlaylistAlbum)
	http.HandleFunc("/tracks", h.HandleRequestTracks)
	http.ListenAndServe(":8080", nil)
}
