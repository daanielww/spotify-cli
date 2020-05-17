package main

import (
	"github.com/daanielww/spotify-cli/App"
	"github.com/daanielww/spotify-cli/cmd/config"
	"log"
	"net/http"
)


func main() {
	c, err := config.DevelopmentConfig()
	if err != nil {
		log.Fatalln("error with config: ", err.Error())
	}

	h := app.Handler{c.Sc, c.S3C}

	http.HandleFunc("/", h.HandleRequest)
	http.ListenAndServe(":8080", nil)
}

