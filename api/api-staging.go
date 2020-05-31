package api

import (
	"encoding/json"
	"net/http"
	"time"
)

// Handlers to handle requests for the staging and local development servers

// Handler for default endpoint - grabbing playlists and albums
func (h *Handler) HandleRequestPlaylistAlbum(w http.ResponseWriter, r *http.Request){
	s3Key := "playlist-album/" + time.Now().Format("2006.01.02 15:04:05")
	bucket := "spotify-cli-staging"
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
	bucket := "spotify-cli-staging"
	res, err := h.GetAndStoreTracks(s3Key, bucket)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}