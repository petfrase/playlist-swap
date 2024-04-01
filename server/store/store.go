package store

import (
	"sync"

	"github.com/petfrase/playlist-swap/models"
)

// Key value store for user sessions
var SessionStore = struct {
	sync.RWMutex
	Sessions map[string]models.UserSession
}{Sessions: make(map[string]models.UserSession)}

// Creates an empty session with the given session ID
func CreateSession(sessionId string) {
	SessionStore.Lock()
	SessionStore.Sessions[sessionId] = models.UserSession{}
	SessionStore.Unlock()
}

// SetUserSession stores the session data for a given session ID
func SetUserSpotifyTokenResponse(sessionId string, data models.SpotifyTokenResponse) {
	SessionStore.Lock()
	SessionStore.Sessions[sessionId] = models.UserSession{
		SpotifyAccessToken:  data.AccessToken,
		SpotifyRefreshToken: data.RefreshToken,
		SpotifyExpiresIn:    data.ExpiresIn,
	}
	SessionStore.Unlock()
}

// SetUserSpotifyCode stores the Spotify code for a given session ID
// Make sure to not overwrite the existing session data
func SetUserSpotifyCode(sessionId string, code string) {
	SessionStore.Lock()
	defer SessionStore.Unlock() // Using defer to ensure Unlock is called

	// Check if the session exists
	if sessionData, exists := SessionStore.Sessions[sessionId]; exists {
		// Modify the struct field
		sessionData.SpotifyCode = code

		// Put the modified struct back into the map
		SessionStore.Sessions[sessionId] = sessionData
	}
}

// GetSessionData retrieves the session data for a given session ID
// It returns the session data and a boolean indicating whether the session was found
func GetUserSession(sessionID string) (models.UserSession, bool) {
	SessionStore.RLock()
	data, exists := SessionStore.Sessions[sessionID]
	SessionStore.RUnlock()
	return data, exists
}
