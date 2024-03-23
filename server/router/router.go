package router

import (
	"github.com/gorilla/mux"
	"playlist-swap/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	// Route for fetching all playlists of the passed type service
	router.HandleFunc("/api/{type}/playlists", middleware.GetPlaylists).Methods("GET", "OPTIONS")
	return router
}
