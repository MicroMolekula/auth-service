package dto

type User struct {
	ExternalId string `json:"external_id"`
	Role       string `json:"role"`
	Email      string `json:"login"`
}
