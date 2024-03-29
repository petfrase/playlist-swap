package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/petfrase/playlist-swap/service"
	"net/http"
	"strconv"
	"strings"
)

func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
	playlists, err := service.GetPlaylists(serviceType, offsetInt, limitInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	json.NewEncoder(w).Encode(playlists)
}
