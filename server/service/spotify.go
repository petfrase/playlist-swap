package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/petfrase/playlist-swap/models"
	"github.com/petfrase/playlist-swap/store"
)

// GetSpotifyAuthURL returns the Spotify authorization URL
func GetSpotifyAuthURL(sessionId string) string {
	// get the url from environment variables
	url := os.Getenv("SPOTIFY_AUTH_URL")

	// generate a random state
	state := generateRandomString(16)

	// Set the session code
	store.SetUserSpotifyCode(sessionId, state)

	// create the query string
	query := fmt.Sprintf(
		"?client_id=%s&response_type=code&redirect_uri=%s&scope=%s&state=%s",
		os.Getenv("SPOTIFY_CLIENT_ID"),
		os.Getenv("SPOTIFY_REDIRECT_URI"),
		os.Getenv("SPOTIFY_SCOPE"),
		state)

	// return the full URL
	return url + query
}

// exchangeSpotifyCode exchanges the code for an access token
// Send a post request to the Spotify API with the code
// Basic auth, base64 encoded client_id:client_secret
// x-www-form-urlencoded
// - code
// - redirect_uri
// - grant_type
func ExchangeSpotifyAuthTokenForAccessToken(code string) (models.SpotifyTokenResponse, error) {
	// Create the request body
	body := fmt.Sprintf(
		"code=%s&redirect_uri=%s&grant_type=authorization_code",
		code,
		os.Getenv("SPOTIFY_REDIRECT_URI"))

	// Create the request
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)
	if err != nil {
		return models.SpotifyTokenResponse{}, err
	}

	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	base64Auth := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))

	// Set the authorization header
	req.Header.Set("Authorization", "Basic "+base64Auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Set the request body
	req.Body = io.NopCloser(strings.NewReader(body))

	// Create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.SpotifyTokenResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.SpotifyTokenResponse{}, err
	}

	// Unmarshal the response body into SpotifyTokenResponse
	var data models.SpotifyTokenResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return models.SpotifyTokenResponse{}, err
	}

	// Return the access token
	return data, nil
}

// GetPlaylists fetches all playlists of the passed type service
// Parameters: Offset int, Limit int
func FetchSpotifyPlaylists(accessToken string, offset int, limit int) ([]models.Playlist, error) {
	// Fetch all playlists from Spotify
	url := "https://api.spotify.com/v1/me/playlists?offset=" + fmt.Sprint(offset) + "&limit=" + fmt.Sprint(limit)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set the authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	fmt.Println("token:", accessToken)

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
		var imageUrl string
		if len(playlist.Images) > 0 {
			imageUrl = playlist.Images[0].URL
		} else {
			// Set a default imageUrl or leave it as an empty string
			imageUrl = "path/to/default/image.jpg" // Example default image URL
		}

		genericPlaylists = append(genericPlaylists, models.Playlist{
			ID:          playlist.ID,
			Name:        playlist.Name,
			Description: playlist.Description,
			ImageUrl:    imageUrl,
			Owner:       playlist.Owner.DisplayName,
		})
	}

	return genericPlaylists, nil
}
