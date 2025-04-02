package dto

import "github.com/google/uuid"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phoneNumber"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type RegisterRequest struct {
	Name        string `json:"name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	ConfirmPass string `json:"confirmPassword" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	RoleID      uint
}

type RegisterResponse struct {
	User UserResponse `json:"user"`
}

type UpdatedRequest struct {
	Name        string `json:"name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Password    *string `json:"password:omitempty"`
	ConfirmPass *string `json:"confirmPassword:omitempty"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	RoleID      uint
}
