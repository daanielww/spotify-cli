package spotifyclient

import "github.com/zmb3/spotify"

type MockSpotify struct{}

func (sc *MockSpotify) SpotifyCombinedPlaylistAlbum() (*RequestResult, error) {
	playlists := []spotify.SimplePlaylist{
		{
			Name:     "Playlist_test",
			ID:       "ID-123",
			IsPublic: false,
		},
	}
	albums := []spotify.SimpleAlbum{
		{
			Name: "Album_Test",
			ID: "ID-567",
			Artists: []spotify.SimpleArtist{
				{
					Name: "Daniel",
					ID: "ID-DANIEL",
				},
			},
		},
	}

	res := &RequestResult{
		Playlists: playlists,
		Albums: albums,
	}

	return res, nil
}

func (sc *MockSpotify) GetTracks() ([]spotify.FullTrack, error) {
	return nil, nil
}