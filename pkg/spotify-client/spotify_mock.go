package spotifyclient

import "github.com/zmb3/spotify"

// Mock Spotify client used for testing and for running locally/development

type MockSpotify struct{}

func (sc *MockSpotify) SpotifyCombinedPlaylistAlbum() (*AlbumsPlaylistRequestResult, error) {
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

	res := &AlbumsPlaylistRequestResult{
		Playlists: playlists,
		Albums: albums,
	}

	return res, nil
}

func (sc *MockSpotify) GetTracks() (*TrackRequestResult, error) {
	tracks := []spotify.FullTrack{
		{
			SimpleTrack: spotify.SimpleTrack{
				Name: "Track_test",
				ID:   "ID-5554",
			},
		},
	}
	res := &TrackRequestResult{Tracks: tracks}
	return res, nil
}