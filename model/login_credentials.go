package model

type LoginCredentials struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}
