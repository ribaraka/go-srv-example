package models

import "time"

type TableUser struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Verified  bool   `json:"verified"`
}

type Credentials struct {
	UserId       int       `json:"user_id"`
	PasswordHash string    `json:"password_hash"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type EmailVerificationToken struct {
	UserId            int       `json:"user_id"`
	VerificationToken string    `json:"verification_token"`
	GeneratedAt       time.Time `json:"updated_at"`
}
