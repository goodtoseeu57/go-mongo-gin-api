package models

type EmailPasswordLoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
