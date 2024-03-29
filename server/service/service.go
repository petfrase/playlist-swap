package service

import (
	"fmt"
	"github.com/petfrase/playlist-swap/models"
	"strings"
)

// GetPlaylists fetches all playlists of the passed type service
func GetPlaylists(serviceType string, offset int, limit int) ([]models.Playlist, error) {

	// Fetch the playlists by service type
	switch strings.ToLower(serviceType) {
	case "spotify":
		return FetchSpotifyPlaylists(offset, limit)
	default:
		return nil, fmt.Errorf("service type %s not supported", serviceType)
	}
}
