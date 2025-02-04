package models

type AuthClaim struct {
	UserName string `json:"user"`
	Password string `json:"password"`
}
