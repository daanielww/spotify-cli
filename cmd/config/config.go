package config

import (
	"flag"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/daanielww/spotify-cli/pkg/s3"
	spotifyclient "github.com/daanielww/spotify-cli/pkg/spotify-client"
	"os"
)

type Configuration struct {
	Sc spotifyclient.SpotifyClient
	S3C s3iface.S3API
	ClientID string
	ClientSecret string
}

func ProductionConfig() (*Configuration, error) {
	c, err := getFlags()
	if err != nil {
		return nil, err
	}

	sc, err := spotifyclient.GetSpotfiyClient(os.Getenv("SPOTIFY_ID"), os.Getenv("SPOTIFY_SECRET"))
	if err != nil {
		return nil, err
	}
	c.Sc = sc

	S3C, err := s3.GetS3(false)
	if err != nil {
		return nil, err
	}
	c.S3C = S3C

	return c, nil
}

func DevelopmentConfig() (*Configuration, error) {
	c, err := getFlags()
	if err != nil {
		return nil, err
	}

	sc, err := spotifyclient.GetSpotfiyClient(c.ClientID, c.ClientSecret)
	if err != nil {
		return nil, err
	}
	c.Sc = sc

	S3C, err := s3.GetS3(true)
	if err != nil {
		return nil, err
	}
	c.S3C = S3C

	return c, nil
}

func getFlags() (*Configuration, error) {
	c, err := parseFlags()
	if err != nil {
		return nil, err
	}

	flag.Parse()

	return c, nil
}

func parseFlags() (*Configuration, error) {
	c := Configuration{}

	flag.StringVar(&c.ClientID, "id", "", "Spotify ClientID for development")
	flag.StringVar(&c.ClientSecret, "secret", "", "Spotify ClientSecret for development")

	return &c, nil
}
