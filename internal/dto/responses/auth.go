package responses

type LoginResponse struct {
	Token     string        `json:"token"`
	User      *UserResponse `json:"user"`
	ExpiresAt string        `json:"expires_at"`
}

type TokenClaims struct {
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	RoleID      uint   `json:"role_id"`
	RoleName    string `json:"role_name"`
	Permissions []string `json:"permissions"`
}
