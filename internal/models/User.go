package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email" validate:"required,email"`
	Phone        string    `json:"phone" validate:"required"`
	Password     string    `json:"password" validate:"required,min=6"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Role         string    `json:"role" validate:"oneof=admin user"`
	Balance      float64   `json:"balance"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserID       int       `json:"user_id"`
}
