package service

import (
	"encoding/json"
	"fmt"
	"github.com/petfrase/playlist-swap/models"
	"io"
	"net/http"
)

// GetPlaylists fetches all playlists of the passed type service
// Parameters: Offset int, Limit int
func FetchSpotifyPlaylists(offset int, limit int) ([]models.Playlist, error) {
	// Fetch all playlists from Spotify
	url := "https://api.spotify.com/v1/me/playlists?offset=" + fmt.Sprint(offset) + "&limit=" + fmt.Sprint(limit)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set the authorization header
	tempToken := "idk"
	req.Header.Set("Authorization", "Bearer "+tempToken)

	// Create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into SpotifyPlaylistResponse
	var data models.SpotifyPlaylistResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	// Extract the playlists from the map
	playlists := data.Items

	// Conveert the Spotify playlists to generic playlists
	var genericPlaylists []models.Playlist
	for _, playlist := range playlists {
		genericPlaylists = append(genericPlaylists, models.Playlist{
			ID:          playlist.ID,
			Name:        playlist.Name,
			Description: playlist.Description,
			ImageUrl:    playlist.Images[0].URL,
		})
	}

	return genericPlaylists, nil
}
