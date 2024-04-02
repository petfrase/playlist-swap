package models

// User Session struct
type UserSession struct {
	SpotifyAccessToken  string
	SpotifyRefreshToken string
	SpotifyExpiresIn    int // Time in seconds
	SpotifyCode         string
}

// Playlist struct
type Playlist struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Owner       string `json:"owner"`
}
