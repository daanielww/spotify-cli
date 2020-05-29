package config

import (
	"flag"
)

type Configuration struct {
	Url string
	IsDevelopment bool
	GetTracks bool
}

func GetConfiguration() (*Configuration, error) {
	c, err := getFlags()
	if err != nil {
		return nil, err
	}

	if c.IsDevelopment {
		c.Url = "http://localhost:8080"
	} else {
		c.Url = "http://ec2-54-208-45-242.compute-1.amazonaws.com:8080"
	}

	if c.GetTracks {
		c.Url = c.Url+"/tracks"
	}

	return c, err
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

	flag.BoolVar(&c.IsDevelopment, "development", false, "Set to development mode")
	flag.BoolVar(&c.GetTracks, "tracks", false, "Set flag if you want to get tracks")

	return &c, nil
}



