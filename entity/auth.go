package entity

import "time"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Username      string     `json:"username"`
	AccessToken   string     `json:"access_token"`
	AccessExpired *time.Time `json:"access_expired"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
