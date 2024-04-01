package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/petfrase/playlist-swap/middleware"
	"github.com/petfrase/playlist-swap/store"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	// Auth endpoints
	// Route for setting the session cookie
	router.HandleFunc("/api/auth/session", setSession).Methods("GET", "OPTIONS")
	// Route for logging in to {type} service
	router.Handle("/api/{type}/login", checkSession(http.HandlerFunc(middleware.ServiceLogin))).Methods("GET", "OPTIONS")
	// Route for handling the Auth token callback
	router.Handle("/api/{type}/callback", checkSession(http.HandlerFunc(middleware.ServiceCallback))).Methods("GET", "OPTIONS")

	// Playlist endpoints
	// Route for fetching all playlists of the passed type service
	router.Handle("/api/{type}/playlists", checkSession(http.HandlerFunc(middleware.GetPlaylists))).Methods("GET", "OPTIONS")

	return router
}

// SetSession sets a session cookie for the user
func setSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Check if the user already has a session cookie
	_, err := r.Cookie("session_id")
	if err != nil {
		fmt.Println("No session cookie found")
		// No session cookie, let's create one
		sessionID := uuid.New().String() // Generate a new session ID

		// Create a new session with the session ID
		store.CreateSession(sessionID)

		// Set the cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: false,
			Secure:   false, // Set to true in HTTPS environments
			SameSite: http.SameSiteLaxMode,
		})

		// Write the response
		json.NewEncoder(w).Encode(map[string]string{"message": "new session set"})
		return
	}

	// Write the response
	json.NewEncoder(w).Encode(map[string]string{"message": "session set"})
}

// CheckSession checks if the user has a session cookie
func checkSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Check if the user has a session cookie
		_, err := r.Cookie("session_id")
		if err != nil {
			friendlyError := "no session found: please log in" + err.Error()
			http.Error(w, friendlyError, http.StatusUnauthorized)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
