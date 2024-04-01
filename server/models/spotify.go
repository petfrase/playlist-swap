package models

// This file contains the models for the Spotify API

// SpotifyPlaylistResponse struct
// This struct is used to unmarshal the response from the Spotify API
type SpotifyPlaylistResponse struct {
	Href     string            `json:"href"`
	Limit    int               `json:"limit"`
	Next     string            `json:"next"`
	Offset   int               `json:"offset"`
	Previous string            `json:"previous"`
	Total    int               `json:"total"`
	Items    []SpotifyPlaylist `json:"items"`
}

// SpotifyPlaylist struct
type SpotifyPlaylist struct {
	Collaborative bool              `json:"collaborative"`
	Description   string            `json:"description"`
	ExternalUrls  map[string]string `json:"external_urls"`
	Href          string            `json:"href"`
	ID            string            `json:"id"`
	Images        []struct {
		Height int    `json:"height"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"images"`
	Name       string `json:"name"`
	Owner      SpotifyOwner
	Public     bool   `json:"public"`
	SnapshotID string `json:"snapshot_id"`
	Tracks     struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type SpotifyOwner struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Followers    struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"followers"`
	Href        string `json:"href"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
	DisplayName string `json:"display_name"`
}

type SpotifyTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}
