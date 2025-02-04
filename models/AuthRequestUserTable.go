package models

import "gorm.io/gorm"

type AuthRequestUser struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func NewAuthRequestUser(name, email, password string) AuthRequestUser {
	return AuthRequestUser{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
