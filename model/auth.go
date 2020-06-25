package model


type AuthLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthTokenRefresh struct {
	TokenRefresh string `json:"token_refresh"`
}