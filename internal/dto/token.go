package dto

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
