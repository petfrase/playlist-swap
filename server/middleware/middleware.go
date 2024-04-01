package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/petfrase/playlist-swap/service"
	"github.com/petfrase/playlist-swap/store"
	"net/http"
	"strconv"
	"strings"
)

func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	// Get the session ID from the cookie
	sessionCookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "session not found", http.StatusUnauthorized)
		return
	}

	// Get the type of service to fetch playlists for
	params := mux.Vars(r)
	serviceType := strings.ToLower(params["type"])

	// Get the offset and limit query parameters
	queries := r.URL.Query()
	offset := queries.Get("offset")
	limit := queries.Get("limit")

	// try to convert the offset and limit to integers
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid offset %s", offset), http.StatusBadRequest)
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid limit %s", limit), http.StatusBadRequest)
		return
	}

	// Fetch playlists for the service type
	playlists, err := service.GetPlaylists(sessionCookie.Value, serviceType, offsetInt, limitInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	json.NewEncoder(w).Encode(playlists)
}

// ServiceLogin logs in to the service of the passed type
func ServiceLogin(w http.ResponseWriter, r *http.Request) {
	// Get the session ID from the cookie
	sessionCookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "session not found", http.StatusUnauthorized)
		return
	}

	// Get the type of service to fetch playlists for
	params := mux.Vars(r)
	serviceType := strings.ToLower(params["type"])

	switch strings.ToLower(serviceType) {
	// if the service is Spotify, redirect to the Spotify login page
	case "spotify":
		http.Redirect(w, r, service.GetSpotifyAuthURL(sessionCookie.Value), http.StatusFound)
	default:
		http.Error(w, fmt.Sprintf("service type %s not supported", serviceType), http.StatusBadRequest)
	}
}

// ServiceCallback handles the callback from the service
func ServiceCallback(w http.ResponseWriter, r *http.Request) {
	// get the session ID from the cookie
	sessionCookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "session not found", http.StatusUnauthorized)
		return
	}

	// get the session data from the store
	sessionId := sessionCookie.Value
	sessionData, exists := store.GetUserSession(sessionId)
	if !exists {
		http.Error(w, "session not found", http.StatusUnauthorized)
		return
	}

	// Get the type of service to fetch playlists for
	params := mux.Vars(r)
	serviceType := strings.ToLower(params["type"])

	switch strings.ToLower(serviceType) {
	// if the service is Spotify, handle the Spotify callback
	case "spotify":
		state := r.URL.Query().Get("state")
		// Check if the session has a spotify state
		if sessionData.SpotifyCode == "" {
			http.Error(w, "session does not have a spotify code", http.StatusBadRequest)
			return
		}
		// Check if the state matches the session state
		if state != sessionData.SpotifyCode {
			http.Error(w, "invalid state", http.StatusBadRequest)
			return
		}

		code := r.URL.Query().Get("code")
		spotifyData, err := service.ExchangeSpotifyAuthTokenForAccessToken(code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		store.SetUserSpotifyTokenResponse(sessionId, spotifyData)
	default:
		http.Error(w, fmt.Sprintf("service type %s not supported", serviceType), http.StatusBadRequest)
	}
}
