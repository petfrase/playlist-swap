package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"playlist-swap/model"
	"playlist-swap/service"
	"strings"
)

func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the type of service to fetch playlists for
	params := mux.Vars(r)
	serviceType := strings.ToLower(params["type"])

	// Fetch playlists for the service type
	playlists, err := service.GetPlaylists(serviceType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	json.NewEncoder(w).Encode(playlists)
}
