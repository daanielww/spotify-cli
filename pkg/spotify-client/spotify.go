package spotifyclient

import (
	"context"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type SpotifyClient interface {
	SpotifyCombined() (*RequestResult, error)
}

type SpotifyStruct struct {
	*spotify.Client
}

type RequestResult struct {
	Playlists []spotify.SimplePlaylist `json:"playlists"`
	Albums    []spotify.SimpleAlbum    `json:"albums"`
}

func GetSpotfiyClient(clientID, clientSecret string) (*SpotifyStruct, error) {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		return nil, err
	}

	sc := spotify.Authenticator{}.NewClient(token)
	return &SpotifyStruct{&sc}, nil
}

func (sc *SpotifyStruct) SpotifyCombined() (*RequestResult, error) {
	playlists, err := sc.getFeaturedPlaylists()
	if err != nil {
		return nil, err
	}

	albums, err := sc.getNewAlbums()
	if err != nil {
		return nil, err
	}

	res := &RequestResult{
		Playlists: playlists,
		Albums:    albums,
	}
	return res, nil
}

func (sc *SpotifyStruct) getFeaturedPlaylists() ([]spotify.SimplePlaylist, error) {
	_, playlists, err := sc.FeaturedPlaylists()
	if err != nil {
		return nil, err
	}

	return playlists.Playlists[:2], err
}

func (sc *SpotifyStruct) getNewAlbums() ([]spotify.SimpleAlbum, error) {
	release, err := sc.NewReleases()
	if err != nil {
		return nil, err
	}

	return release.Albums[:2], err
}
