package service

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/petfrase/playlist-swap/models"
	"github.com/petfrase/playlist-swap/store"
)

// GetPlaylists fetches all playlists of the passed type service
func GetPlaylists(sessionId string, serviceType string, offset int, limit int) ([]models.Playlist, error) {

	// Fetch the playlists by service type
	switch strings.ToLower(serviceType) {
	case "spotify":
		// get the spotify token from the session store
		sessionData, exists := store.GetUserSession(sessionId)
		if !exists {
			return nil, fmt.Errorf("session not found")
		}

		return FetchSpotifyPlaylists(sessionData.SpotifyAccessToken, offset, limit)
	default:
		return nil, fmt.Errorf("service type %s not supported", serviceType)
	}
}

// Internal Helper functions

// generates a random string of length i
// with characters from the set [a-zA-Z0-9]
func generateRandomString(i int) string {
	// define the character set
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// create a byte slice of length i
	b := make([]byte, i)

	// loop i times
	for i := range b {
		// generate a random index
		r := rand.Intn(len(charSet))

		// set the byte at index i to the character at index r
		b[i] = charSet[r]
	}

	// return the byte slice as a string
	return string(b)
}
